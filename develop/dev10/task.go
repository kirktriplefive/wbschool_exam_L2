package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	flagTimeout := flag.String("timeout", "10s", "timeout for connection (duration)")
	flag.Parse()
	if len(flag.Args()) < 2 {
		flag.Usage()
		os.Exit(2)
	}

	timeout, err := time.ParseDuration(*flagTimeout)
	
	if err != nil {
		flag.Usage()
		os.Exit(2)
	}
	address:=flag.Arg(0)+":"+flag.Arg(1)//записываем адресс в строку
	fmt.Println(address)

	client:=newClient(address, timeout)

	if err := client.dial(); err != nil {
		log.Fatalln("Ошибка подключения:", err)
	}

	abort := client.readFromWriteToConn()
	client.waitOSKill()

	<-abort

	if err := client.cancelReadWriteClose(); err != nil {
		log.Fatalln(err)
	}

	time.Sleep(time.Second)

}

type client struct {
	serverAddr  string
	timeout     time.Duration
	conn        net.Conn
	ctx         context.Context
	cancel      context.CancelFunc
	errChan   chan bool
	chanForSTDIN   chan string
	lastMessage string
}

func newClient(address string, timeout time.Duration) client { //Конструктор для клиента 
	c := client{
		serverAddr: address,
		timeout:    timeout,
		errChan:  make(chan bool),
		chanForSTDIN:  make(chan string),
	}
	c.ctx, c.cancel = context.WithCancel(context.Background())
	return c
}

func (c *client) dial() error {
	var err error
	dialer := &net.Dialer{Timeout: c.timeout}
	c.conn, err = dialer.Dial("tcp", c.serverAddr)
	if err == nil {
		log.Printf("Подключён к: %s", c.serverAddr)
		log.Println("'CRTL+D' для отключения")
	}
	return err
}

func (c *client) waitOSKill() {
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
		sig := <-ch
		fmt.Println(sig)
		c.errChan <- true
	}()
}

func (c *client) close() error {
	log.Print("Прекращаем подключение ")
	if err := c.conn.Close(); err != nil {
		return err
	}
	log.Print("Подключение завершено")
	return nil
}

func (c *client) cancelReadWriteClose() error {
	c.cancel()
	time.Sleep(time.Second)
	if err := c.close(); err != nil {
		return err
	}
	return nil
}

func (c *client) readFromConn() chan bool {
	go c.read()
	return c.errChan
}

func (c *client) readFromWriteToConn() chan bool {
	go c.read()
	go c.write()
	return c.errChan
}

func (c *client) read() {
	reply := make([]byte, 1024)
OUTER:
	for {
		select {
		case <-c.ctx.Done():
			log.Print("Заканчиваем чтение")
			break OUTER
		default:
			if err := c.conn.SetReadDeadline(time.Now().Add(time.Second)); err != nil {
				log.Println(err)
			}
			n, err := c.conn.Read(reply)
			if err != nil {
				if err == io.EOF {
					log.Println("Заканчиваем чтение")
					c.errChan <- true
					break OUTER
				}
				if netErr, ok := err.(net.Error); ok && !netErr.Timeout() {
					log.Println(err)
				}
			}
			if n == 0 {
				break
			}
			bs := reply[:n]
			if len(bs) != 0 {
				c.lastMessage = string(bs)
			}
			fmt.Println(c.lastMessage)
		}
	}
	log.Println("Закончили чтение")
}

func (c *client) write() {
	go func(stdin chan<- string) {
		reader := bufio.NewReader(os.Stdin)
		for {
			s, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					log.Print("Заканчиваем запись")
					c.errChan <- true
					return
				}
				log.Println(err)
			}
			stdin <- s
		}
	}(c.chanForSTDIN)

OUTER:
	for {
		select {
		case <-c.ctx.Done():
			log.Print("Заканчиваем запись")
			break OUTER
		default:

		STDIN:
			for {
				select {
				case stdin, ok := <-c.chanForSTDIN:
					if !ok {
						break STDIN
					}
					if _, err := c.conn.Write([]byte(stdin)); err != nil {
						log.Println(err)
					}
					c.lastMessage = stdin
				case <-time.After(time.Second):
					break STDIN
				}
			}
		}
	}
	log.Println("Закончили запись")
}