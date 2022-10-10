package main

import "fmt"

type Sender interface {
	Send(msg string) error
}

type Email struct {
	Address string
}

func (e *Email) Send(msg string) error {
	fmt.Printf("MSG:%v and Address:%v \n", msg, e.Address)
	return nil
}

type Phone struct {
	Number  int
	Balance int
}

type Caller interface {
	Call(Number int) error
}

type IPhone interface {
	Caller
	Sender
	MyFunc()
	MyFunc2()
}

func (p *Phone) Send(msg string) error {
	fmt.Printf("MSG:%v and Phone Number:%v and Balance:%v \n", msg, p.Number, p.Balance)
	return nil
}

func Notify(s Sender) {
	err := s.Send("Notify Message")
	if err != nil {
		fmt.Println(err)
		return
	}

	phone, ok := s.(*Phone)
	if ok {
		fmt.Println(phone.Balance)
	}

	fmt.Println("Success")
}

func Notify2(s Sender) {
	err := s.Send("Notify Message")
	if err != nil {
		fmt.Println(err)
		return
	}

	switch s.(type) {
	case *Email:
		fmt.Println("This is no balance attribute")
	case *Phone:
		phone, ok := s.(*Phone)
		if ok {
			fmt.Println(phone.Balance)
		}
	}

	fmt.Println("Success")
}

func Notify3(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("This is integer type. Please Exit!")
	}

	s, ok := i.(Sender)
	if !ok {
		fmt.Printf("Sender Attribute is Required Not %T. Please exit\n", i)
		return
	}

	err := s.Send("Notify Message")
	if err != nil {
		fmt.Println("Message Failed to send")
		return
	}

	fmt.Println("Success")
}

func main() {
	email := &Email{Address: "golang@demo.com"}
	Notify(email)

	phoneNumber := &Phone{
		Number:  012312034123,
		Balance: 1000,
	}

	Notify(phoneNumber)

	fmt.Println("-------------------------")

	Notify2(email)
	Notify2(phoneNumber)

	fmt.Println("-------------------------")

	Notify3(email)
	Notify3(phoneNumber)
	Notify3(12)
	Notify3("Hello")

}
