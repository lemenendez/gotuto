package main


import (
	"fmt"	
	"bufio"
	"os"
	"strings"
)

type Animal interface {
	Speak()
	Eat()
	Move()
}

type Cow struct {
	food, move, noise string
}

func (a Cow) Speak() {
	fmt.Printf("%v", a.noise)
} 

func (a Cow) Eat() {
	fmt.Printf("%v", a.food)
}

func (a Cow) Move() {
	fmt.Printf("%v", a.move)
}

type Bird struct {
	food, move, noise string
}

func (a Bird) Speak() {
	fmt.Printf("%v", a.noise)
} 

func (a Bird) Eat() {
	fmt.Printf("%v", a.food)
}

func (a Bird) Move() {
	fmt.Printf("%v", a.move)
}

type Snake struct {
	food, move, noise string
}

func (a Snake) Speak() {
	fmt.Printf("%v", a.noise)
} 

func (a Snake) Eat() {
	fmt.Printf("%v", a.food)
}

func (a Snake) Move() {
	fmt.Printf("%v", a.move)
}


func readCommand(cmd string) (string, string, string) {
	words := strings.Split(cmd, " ")
	if len(words)<3 {
		return "","",""
	}
	return words[0], words[1], words[2]
}

func newAnimal(t string) Animal {
	var animal Animal
	switch t {
	case "cow":
		animal = Cow{"grass", "walk", "moo"}
	case "bird":
		animal = Bird{"worms", "fly", "peep"} 
	case "snake":
		animal = Snake{"mice", "slither", "hsss"}
	}
	return animal
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	animals:=make(map[string] Animal)

	for {
		fmt.Print(">")
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			cmd1, cmd2, cmd3 := readCommand(text)
			switch cmd1 {
			case "newanimal": 				
				animals[cmd2] = newAnimal(cmd3)
			case "query":				
				if a, ok := animals[cmd2]; ok {
					switch cmd3 {
					case "move":
						a.Move()
					case "speak":
						a.Speak()
					case "eat":
						a.Eat()
					}					
				}
			}		
		}
}

}


