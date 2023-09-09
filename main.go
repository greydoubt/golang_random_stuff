
package main

import "fmt"

// Animal interface to be implemented by all types of animals
type Animal interface {
    Speak() string
}

// Dog struct representing a dog
type Dog struct {
    Name string
}

// Speak method for a Dog
func (d Dog) Speak() string {
    return fmt.Sprintf("%s says Woof!", d.Name)
}

// Cat struct representing a cat
type Cat struct {
    Name string
}

// Speak method for a Cat
func (c Cat) Speak() string {
    return fmt.Sprintf("%s says Meow!", c.Name)
}

// function to get a speaking animal's sound
func speakingAnimal(a Animal) {
    fmt.Println(a.Speak())
}

func main() {
    dog := Dog{Name: "Buddy"}
    cat := Cat{Name: "Whiskers"}

    animals := []Animal{dog, cat}

    // Iterate over the slice of animals, printing each one's Speak() string.
    for _, animal := range animals {
        speakingAnimal(animal)
    }
}
