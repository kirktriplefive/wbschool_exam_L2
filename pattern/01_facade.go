package main

import (
	"fmt"
	"math/rand"
)

//Шаблон проектирования Фасад - структурный шаблон проектирования, позволяющий скрыть сложность
//системы путём сведения всех возможных внешних вызовов к одному объекту, делегирующему их соответствующим объектам системы.

//Применение этого шаблона проектирования заключается в том, чтобы инкапсулировать сложную систему в
//единственный интерфейсный объект. Это необходимо для связи между подсистемами.

//Наш Фасад
type customerFacade struct {
	person *person
	item *item
}

func newCostomerFacade(personId, itemId int ) *customerFacade {
	customerFacade := &customerFacade{
		person: newPerson(personId),
		item: newItem(itemId), 
	}
	return customerFacade
}

func (c *customerFacade) getItem(personId, itemId int) error{
	if err:=c.person.getPerson(personId); err != nil {
		return err
	}
	name, err:=c.item.getItemById(itemId)
	if err != nil {
		return err
	}
	fmt.Println("Ваш товар это - ", name)
	return nil
}

//Челоек - покупатель
type person struct{
	id int
}

func newPerson(personId int) *person {
	return &person{
		id: personId,
	}
}

//Функция для определения существования такого покупателя, если его нет - мы выводим ошибку
func (p *person) getPerson(personId int) error {
	if p.id != personId {
		return fmt.Errorf("Такого покупателя не существует")
	}
	return nil
}

//Товар 
type item struct {
	id int
	name string
}

func newItem(itemId int) *item {
	return &item{
		id: itemId,
		name: RandStringRunes(10), 
	}
}

//Функция для поиска названия товара
func (i *item) getItemById(itemId int) (string, error) {
	if i.id != itemId {
		return "", fmt.Errorf("Такого товара не существует")
	}
	return i.name, nil
}





func main() {
	customerFacade:=newCostomerFacade(1,1)
	if err:=customerFacade.getItem(1, 1); err!= nil {
		fmt.Println(err)
	}

	if err:=customerFacade.getItem(2, 1); err!= nil {
		fmt.Println(err)
	}

	if err:=customerFacade.getItem(1, 2); err!= nil {
		fmt.Println(err)
	}
	customerFacade=newCostomerFacade(3,10)
	if err:=customerFacade.getItem(3, 10); err!= nil {
		fmt.Println(err)
	}

	


}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//Функция для создания рандомной строки
func RandStringRunes(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}