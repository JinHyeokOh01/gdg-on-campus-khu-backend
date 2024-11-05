package main

import (
	"fmt"
)

type User struct {
	Name string
	age  int
}

func change(a *User, b *User) {
	temp := a.Name
	a.Name = b.Name
	b.Name = temp
}

func sorting(list []User) {
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i].Name > list[j].Name {
				change(&list[i], &list[j])
			}
		}
	}
}

func main() {
	list := []User{
		{"Paul", 19},
		{"John", 21},
		{"Jane", 35},
		{"Abraham", 25},
	}
	sorting(list)
	for _, user := range list {
		fmt.Println(user.Name)
	}
}
