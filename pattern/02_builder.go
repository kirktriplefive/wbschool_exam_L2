package main

import "fmt"

//Шаблон проектирования Строитель - порождающий шаблон проектирования предоставляет способ создания составного объекта.
//Он позволяет создавать сложные объекты пошагово. Мы выносим конструирование объекта за пределы его собственного класса,
//порчив это отдельным классам(строителям). Мы разбиваем процесс на отдельные шаги, поочередно вызывая методы строителя.

type builder interface {//интерфейс строителя 
	setCustomerName(name string)
	setCustomerDiscount(discount int)
	getCustomer() customer

}

func getBuilder(builderType string) builder {//выбор строителя для разных типов покупателей
	switch builderType {
	case "vip":
		return &vipBuilder{}
	case "normal": 
		return &normalBuilder{}
	default: 
		return nil
	}
}

//Строитель для ВИП-клиентов, имеющих скидку
type vipBuilder struct {
	name string
	discount int
}

func newVipBuilder() *vipBuilder{//конструктор для Вип-строителя
	return &vipBuilder{}

}

func (v *vipBuilder) setCustomerName(name string) {
	v.name = name
}

func (v *vipBuilder) setCustomerDiscount(discount int) {
	v.discount = discount
}

func (v *vipBuilder) getCustomer() customer {
	return customer{
		name: v.name,
		discount: v.discount,
	}
}

//Строитель для обычных клиентов, у которых скидка = 0
type normalBuilder struct {
	name string
	discount int
}

func newNormalBuilder() *vipBuilder{//конструктор для обычного строителя
	return &vipBuilder{}

}

func (n *normalBuilder) setCustomerName(name string) {
	n.name = name
}

func (v *normalBuilder) setCustomerDiscount(discount int) {
	v.discount = 0
}

func (n *normalBuilder) getCustomer() customer {
	return customer{
		name: n.name,
		discount: 0,
	}
}

type director struct {//иногда при использовании данного ШП используется Директор, который задает последовательность выполнения шагов
	builder builder

}

func newDirector(b builder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b builder) {
	d.builder = b
}

func (d *director) newCustomer(name string, discount int) customer{
	d.builder.setCustomerName(name)
	d.builder.setCustomerDiscount(discount)
	return d.builder.getCustomer()

}


type customer struct{
	name string
	discount int

}

func main() {
	vipBuilder:=getBuilder("vip")
	normalBuilder:=getBuilder("normal")

	director:=newDirector(vipBuilder)
	vipCustomer:=director.newCustomer("Kirill", 99)

	fmt.Println("VIP:", vipCustomer.name, "discount:", vipCustomer.discount)

	director = newDirector(normalBuilder)
	normalCustomer:=director.newCustomer("Vlad", 0)

	fmt.Println(normalCustomer.name, "discount:", normalCustomer.discount )


}