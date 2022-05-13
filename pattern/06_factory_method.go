package main

import "fmt"

//ШП фабричный метод - порождающий шаблон проектирования, предоставляющий подклассам (дочерним классам) интерфейс
//для создания экземпляров некоторого класса. В момент создания наследники могут определить, какой класс создавать.
//Иными словами, данный шаблон делегирует создание объектов наследникам родительского класса.
//Это позволяет использовать в коде программы не конкретные классы, а манипулировать абстрактными объектами на более
//высоком уровне.

//таким образом мы можем переопередлять фабричный метод для создания разных типов объектов
//Эти объекты должны иметь общий интерфейс!
//В Го отсутсвуют  классы и наследственность к которым мы привыкли в обычном ООП

type worker interface {//интерфейс - общий для всех работников
	getSalary() int
	getActivity() string
	getName() string
	setName(string) 
	setSalary(int)
	setActivity(string)
}

type person struct {//тип - человек(работник)
	name string
	salary int
	activity string
}

func (d *person) setName(name string) {
	d.name = name
}

func (d *person) setSalary(s int) {
	d.salary = s
}

func (d *person) setActivity(language string) {
	d.activity = language
}

func (d *person) getName() string {
	return d.name

}
func (d *person) getSalary() int {
	return d.salary
}

func (d *person) getActivity() string {
	return d.activity
}

type developer struct {
	person
}

func newDeveloper() worker {
	return &developer{
		person: person{
			name: "John",
			salary: 100000, 
			activity: "Go",
		},
	}
}

type director struct {
	person
}

func newDirector() worker {
	return &developer{
		person: person{
			name: "Kirill",
			salary: 150000, 
			activity: "Sale",
		},
	}
}

func fabric(position string) (worker, error) {//фабрика, определяющая создание объекта
	switch position {
		case "director":
			return newDirector(), nil
		case "developer":
			return newDeveloper(), nil
	}
	return nil, fmt.Errorf("Нет такого работника")
}

func main() {
	if developer, err:=fabric("developer"); err!= nil {
		fmt.Println(err)
	} else {
		fmt.Println("developer: ", developer.getName(), "зп: ", developer.getSalary(), "яп:", developer.getActivity())
	}
	if director, err:=fabric("director"); err!= nil {
		fmt.Println(err)
	} else {
		fmt.Println("director: ", director.getName(), "зп: ", director.getSalary(), "отдел:", director.getActivity())
	}
	if director, err:=fabric("kirill"); err!= nil {
		fmt.Println(err)
	} else {
		fmt.Println(director)
	}
}







