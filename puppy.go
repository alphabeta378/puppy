package puppy

import "github.com/alphabeta378/dog"

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.When_grown(Bark())
}

func BigBarks() string {
	return dog.When_grown(Barks())
}