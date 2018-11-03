package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

func (h *Human) Sayhi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (h *Human) Sing(changdianshenmeba string) {
	fmt.Println("La la, la la la, la la la la la...", changdianshenmeba)
}

func (h *Human) Guzzle(nidaodiheleduoshao string) {
	fmt.Println("Guzzle Guzzle Guzzle...", nidaodiheleduoshao)
}

func (e *Employee) Sayhi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //此句可以分成多行
}

func (s *Student) BorrowMoney(nitamazaiwuruxueshengma float32) {
	s.loan += nitamazaiwuruxueshengma
}

//Employee实现SpendSalary方法
func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount // More vodka please!!! Get me through the day!
}

//定义 interface
type Men interface {
	Sayhi()
	Sing(changdianshenmeba string)
	Guzzle(nidaodiheleduoshao string)
}

type YoungChap interface {
	SayHi()
	Sing(changdianshenmeba string)
	BorrowMoney(nitamazaiwuruxueshengma float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(changdianshenmeba string)
	SpendSalary(amount float32)
}

func main() {

	mike := Student{
		Human{
			"Mark", 25, "222-222-YYYY",
		},
		"yizhong",
		111,
	}

	sam := Employee{
		Human{
			"Sam", 45, "111-888-XXXX",
		},
		"Golang Inc",
		123,
	}

	paul := Student{
		Human{
			"Sam", 45, "111-888-XXXX",
		},
		"Golang Inc",
		123,
	}

	var i ElderlyGent
	i.SayHi()

	x := make([]Men, 3)
	x[0], x[1], x[2] = paul, sam, mike
	for _, value := range x {
		value.Sayhi()
	}

}
