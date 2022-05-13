package main

import "fmt"

//ШП стратегия - поведенческий шаблон проектирования, предназначенный для определения семейства алгоритмов,
//инкапсуляции каждого из них и обеспечения их взаимозаменяемости.
//Это позволяет выбирать алгоритм путём определения соответствующего класса.
//Шаблон Strategy позволяет менять выбранный алгоритм независимо от объектов-клиентов, которые его используют.

type strategy interface {//интерфейс стратегии
	getSalary() (int, error)
}

type context struct {//контекст, хранящий в себе стратегию и выполняющий ее
	strategy
}

func (c *context) setStrategy(strategy strategy) { 
	c.strategy = strategy
}

func (c *context) doStrategy() {
	if salary, err:=c.strategy.getSalary(); err!= nil {
		fmt.Println(err)
	} else {
		fmt.Println(salary)
	}
}

type director struct {
	name string
	salary int
}

func (d *director) getSalary() (int, error) {
	salary:=d.salary*10
	if salary==0 {
		return 0, fmt.Errorf("Не указана зарплата директора по имени %s", d.name)
	}
	return salary, nil
}

type developer struct {
	name string
	salary int
}

func (d *developer) getSalary() (int, error) {
	salary:=d.salary*2
	if salary==0 {
		return 0, fmt.Errorf("Не указана зарплата для разработчика по имени %s", d.name)
	}
	return salary, nil
}

func main() {
	developer:=&developer{
		name: "Kirill", 
		salary: 100000,
	}
	director:=&director{
		name: "John",
		salary: 200000,
	}
	person:=&context{}
	person.setStrategy(developer)
	person.doStrategy()
	person.setStrategy(director)
	person.doStrategy()
	director.salary=0
	person.doStrategy()
}