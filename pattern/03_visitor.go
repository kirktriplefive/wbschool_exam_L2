package main

import "fmt"

//поведенческий шаблон проектирования, описывающий операцию, которая выполняется над объектами других классов. 
//При изменении visitor нет необходимости изменять обслуживаемые классы.

type person interface{//интерфейс человека
	getPersonType() string
	accept(visitor)
}

type developer struct {//тип - разработчик
	name string
	salary int
	language string
}

func (d *developer) accept(v visitor){
	v.getDeveloper(d)
}

type director struct {//тип - разработчик
	name string
	salary int
	department string
}

func (d *director) accept(v visitor){
	v.getDirector(d)
}

type visitor interface {//интерфейс для посетителей
	getDeveloper(*developer)
	getDirector(*director)
}

type salaryGetter struct {//определенный посетитель для подсчета итоговой зарплаты
	salary float32
}

func (s *salaryGetter) getDeveloper(d *developer) {
	salary:=float32(d.salary)*0.7
	s.salary=salary
	fmt.Println("Итоговая зарплата разработчика", salary)
} 

func (s *salaryGetter) getDirector(d *director) {
	salary:=float32(d.salary)*0.9
	s.salary=salary
	fmt.Println("Итоговая зарплата разработчика", salary)
} 

type activityGetter struct {
	activity string
}

func (a *activityGetter) getDeveloper(d *developer) {
	a.activity = d.language
	fmt.Println("Разработчик программирует на: ", d.language)
} 

func (a *activityGetter) getDirector(d *director) {
	a.activity=d.department
	fmt.Println("Директор главный в отделе под названием: ", d.department)
} 

func main() {
	developer:=&developer{name: "Kirill", salary: 100000, language: "GoLang"}
	director:=&director{name: "Ivan", salary: 200000, department: "Sale"}

	salaryGetter:=&salaryGetter{}
	activityGetter:=&activityGetter{}

	developer.accept(salaryGetter)
	director.accept(salaryGetter)

	developer.accept(activityGetter)
	director.accept(activityGetter)
	
}