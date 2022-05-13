package main

import "fmt"

//ШП состоние — поведенческий шаблон проектирования. Используется в тех случаях, когда во время выполнения
//программы объект должен менять своё поведение в зависимости от своего состояния.

//Этот подход можно применить к отдельным объектам, например, если у нас есть объект Стажёр и он может находится в
//4 состояниях - l0, l1, l2, и отдел. И в зависимости от состояний результат его работы приводит к разным последствиям.
//Наш первоначальный объект должен будет хранить ссылку на один из объектов состояний.

//Представим нашего стажера, который может находится в 4-х состояниях, как перечислено выше.

//Также у стажера есть два действия - делает задачу и сдает задачу

type trainee struct {
	l0         state
	l1         state
	l2         state
	department state

	current state

	performance bool
	made        bool
}

func newTrainee(performance, made bool) *trainee {
	trainee := &trainee{
		performance: performance,
		made:        made,
	}
	l0 := &l0{
		trainee: trainee,
	}
	l1 := &l1{
		trainee: trainee,
	}
	l2 := &l2{
		trainee: trainee,
	}
	department := &department{
		trainee: trainee,
	}

	trainee.setState(l0) //первое состояние - л0
	trainee.l0 = l0
	trainee.l1 = l1
	trainee.l2 = l2
	trainee.department = department

	return trainee
}

func (t *trainee) setState(s state) {
	t.current = s
}

func (t *trainee) performingTask() error {
	return t.current.perfomingTask()
}

func (t *trainee) protectionTask() error {
	return t.current.protectionTask()
}

type state interface { //интерфейс состояния
	perfomingTask() error
	protectionTask() error
}

type l0 struct { //состояние л0
	trainee *trainee
}

func (l *l0) perfomingTask() error {
	if !l.trainee.performance {
		l.trainee.performance = true
		return fmt.Errorf("Вы еще не делаете эту задачу lo(, но теперь начинаете")
	} else {
		fmt.Println("Продолжайте в том же духе и приступайте к сдачи l0 как можно скорее")
		l.trainee.performance = false
		return nil
	}
}

func (l *l0) protectionTask() error {
	if !l.trainee.made {
		l.trainee.made = true
		return fmt.Errorf("Вы еще не сделали задачу и не приступили к сдаче l0, но уже практически доделали")
	} else {
		fmt.Println("Молодец, ты сдал задачу l0!")
		l.trainee.setState(l.trainee.l1) //переходим в следующее состояние л1 если сдали л0
		l.trainee.made = false
		return nil
	}
}

type l1 struct { //состояние л1
	trainee *trainee
}

func (l *l1) perfomingTask() error {
	if !l.trainee.performance {
		l.trainee.performance = true
		return fmt.Errorf("Вы еще не делаете эту задачу( l1, но теперь начинаете")
	} else {
		fmt.Println("Продолжайте в том же духе и приступайте к сдачи l1 как можно скорее")
		l.trainee.performance = false
		return nil
	}
}

func (l *l1) protectionTask() error {
	if !l.trainee.made {
		l.trainee.made = true
		return fmt.Errorf("Вы еще не сделали задачу и не приступили к сдаче l1, но уже практически доделали")
	} else {
		fmt.Println("Молодец, ты сдал задачу l1!")
		l.trainee.setState(l.trainee.l1) //переходим в следующее состояние л1 если сдали л0
		l.trainee.made = false
		return nil
	}
}

type l2 struct { //состояние л2
	trainee *trainee
}

func (l *l2) perfomingTask() error {
	if !l.trainee.performance {
		l.trainee.performance = true
		return fmt.Errorf("Вы еще не делаете эту задачу( l2, но теперь начинаете")
	} else {
		fmt.Println("Продолжайте в том же духе и приступайте к сдачи l2 как можно скорее")
		l.trainee.performance = false
		return nil
	}
}

func (l *l2) protectionTask() error {
	if !l.trainee.made {
		l.trainee.made = true
		return fmt.Errorf("Вы еще не сделали задачу и не приступили к сдаче l2, но уже практически доделали")
	} else {
		fmt.Println("Молодец, ты сдал задачу l2!")
		l.trainee.setState(l.trainee.l1) //переходим в следующее состояние л1 если сдали л0
		l.trainee.made = false
		return nil
	}
}

type department struct { //состояние отдел
	trainee *trainee
}

func (d *department) perfomingTask() error {
	if !d.trainee.performance {
		d.trainee.performance = true
		return fmt.Errorf("Вы еще не делаете эту задачу в отделе(")
	} else {
		fmt.Println("Продолжайте в том же духе и приступайте к сдачи задачи в отделе как можно скорее")
		d.trainee.performance = false
		return nil
	}
}

func (d *department) protectionTask() error {
	if !d.trainee.made {
		d.trainee.made = true
		return fmt.Errorf("Вы еще не сделали задачу и не приступили к сдаче задачи в отделе")
	} else {
		fmt.Println("Молодец, ты хорошо показываешь себя в отделе ")
		d.trainee.made = false
		return nil
	}
}

func main() {
	trainee := newTrainee(true, false)
	if err := trainee.performingTask(); err != nil {
		fmt.Println(err)
	}
	if err := trainee.protectionTask(); err != nil {
		fmt.Println(err)
	}
	if err := trainee.protectionTask(); err != nil {
		fmt.Println(err)
	}
	if err := trainee.performingTask(); err != nil {
		fmt.Println(err)
	}
	if err := trainee.performingTask(); err != nil {
		fmt.Println(err)
	}
}
