package main

import "fmt"

type user struct {
	name string
	age  int
}

func main() {
	// Способы инициализации экземпляра структуры
	u1 := user{name: "Pasha", age: 32}
	println(u1.String())
	u2 := &user{name: "Anton", age: 33}
	println(u2.String())
	u3 := new(user)
	u3.name = "Pasha"
	u3.age = 32
	println(u3.String())
}

func (u *user) String() string {
	return fmt.Sprintf("%s %d", u.name, u.age)
}
