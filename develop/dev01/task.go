package dev01

//Создать программу печатающую точное время с использованием NTP -библиотеки. Инициализировать как go module.
//Использовать библиотеку github.com/beevik/ntp. Написать программу печатающую текущее время / точное время с 
//использованием этой библиотеки.

import (
	"fmt"
	"os"
	t "time"
	"github.com/beevik/ntp"
)

func currentTime() error {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")//возвращает время используя информацию с удаленного сервера
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v ", err)
		os.Exit(1)
		return err
	}
	fmt.Println("Current Local Time:")
	fmt.Println(time)
	return nil
}

func currentTimeWithTime() (error, t.Time) {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v ", err)
		os.Exit(1)
		return err, time
	}
	fmt.Println("Current Local Time:")
	fmt.Println(time)
	return nil, time
}