package main

import "fmt"

// Структура Human
type Human struct {
	Name string // Поля класса (имя и возраст)
	Age  int
}

// Метод Speak для Human
func (h Human) Speak() {
	fmt.Println("My name is", h.Name, ", and I am", h.Age, "years old.")
}

// Определяем структуру Action, которая встраивает Human, это позволяет Action наследовать методы и поля Human
type Action struct {
	Human // Встраивание структуры Human
	Speed int
}

// Метод Run для Action
func (a Action) Run() {
	fmt.Println(a.Name, "is running at speed", a.Speed)
}

func task1() {
	// Инициализируем структуру Action
	action := Action{
		Human: Human{
			Name: "John",
			Age:  30,
		},
		Speed: 10,
	}
	// Теперь мы можем вызвать методы, определенные для Human, напрямую на экземпляре Action
	action.Speak()       // Выведет: My name is John, and I am 30 years old.
	action.Run()         // Выведет: John is running at speed 10
	action.Human.Speak() // А еще можно вот так, также выводит "My name is John, and I am 30 years old."
}
