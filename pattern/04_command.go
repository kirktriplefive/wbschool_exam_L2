package main

import "fmt"

//шаблон проетикрования команда - поведенческий шаблон проектирования, используемый при объектно-ориентированном
//программировании, представляющий действие. Объект команды заключает в себе само действие и его параметры.
//То есть при использовании данного ШП мы превращаем запросы в объекты, что позволяет нам передавать их как аргументы, ставить в очередь, логировать и поддерживать отмену операций

type order struct {//тип заказ. реализуем случай, когда заказ приходит от клиента и когда он приходит от разработчика(тестовый)
	command command 
}

func (o *order) send(s string) {
	o.command.execute(s)
}

type command interface{
	execute(string)
}

type checkout struct {
	executor executor
}

func (c *checkout) execute(s string) {
	c.executor.add(s)
}

type deleting struct {
	executor executor
}

func (d *deleting) execute(s string) {
	d.executor.delete(s)
}

type executor interface {
	add(string)
	delete(string)
}

type database struct {
	action string
}

func (d *database) add(s string) {
	if s=="add"{
		d.action=s
		fmt.Println("Добавляем заказ в базу данных")
	} else {
		fmt.Println("Нет такого действия для базы данных, если вы хотите добавить, нужно add")
	}

}

func (d *database) delete(s string) {
	if s=="delete"{
		d.action=s
		fmt.Println("Удаляем заказ из базы данных")
	} else {
		fmt.Println("Нет такого действия для базы данных, если вы хотите удалить, нужно delete")
	}

}

func main() {
	database := &database{}

	checkout := &checkout{
		executor: database,
	}

	deleting := &deleting{
		executor: database,
	}

	addOrder := &order{
		command: checkout,
	}

	deleteOrder := &order{
		command: deleting,
	}

	deleteOrder.send("add")//будет ошибка
	deleteOrder.send("delete")//верная работа

	addOrder.send("add")//верно
	addOrder.send("мне нужно удалить но не полностью(изменить)")

}



