package main

import "fmt"

//ШП цепочка вызовов(ответсвенности) - В объектно-ориентированном дизайне шаблон цепочки ответственности представляет собой шаблон поведенческого проектирования,
//состоящий из источника командных объектов и ряда объектов обработки.
//Каждый объект обработки содержит логику, которая определяет типы командных объектов, которые он может обрабатывать; остальные передаются следующему объекту обработки в цепочке.
//Также существует механизм добавления новых объектов обработки в конец этой цепочки.

//Нам необходимо превратить все поведения в объекты, чтобы проверка выполнения была в одом классе с одним методом выполнения
//Каждый из объектов обработчика будет иметь ссылку на следующий обработчик в цепи, чтобы была возможность передать обработку дальше(но не обязательно)

type processor interface {
	execute(*user)
	nextProc(processor)
}

type login struct {
	nextProcess processor
}

func (l *login) execute(user *user) {
	if user.is_signup || user.is_login {
		if user.is_signup && !user.is_login { 
			fmt.Println("Вы уже зарегистрированы")
			user.is_login = true
		} else if user.is_login && !user.is_signup{
			fmt.Println("Ошибка")
		} else {
			fmt.Println("Вы уже вошли в систему")
		}
	} else {
		l.nextProcess.execute(user)
	}
}

func (l *login) nextProc(processor processor) { //Установка следующего обработчика в цепочке
	l.nextProcess=processor
}

type signup struct{
	nextProcess processor
}

func (s *signup) execute(user *user) {
	user.is_signup=true
	fmt.Println("Теперь вы зарегистрированы и сейчас будет произведен вход в систему")
	user.is_login =true
}

func (s *signup) nextProc(processor processor) {//Установка следующего обработчика в цепочке, в нашей задаче необязательно, но удобно для расширения количества обработчиков
	s.nextProcess=processor
}

type user struct {
	is_login bool
	is_signup bool
}

func main() {
	login:=&login{}
	signup:=&signup{}

	login.nextProc(signup)

	user:=&user{is_login: true, is_signup: false}

	login.execute(user)
	login.execute(user)

}