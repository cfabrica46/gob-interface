package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

//Dog es exportado ...
type Dog struct {
}

//S ...
func (d *Dog) S() {
	fmt.Println("GUAU")
}

//Cat es exportado ...
type Cat struct {
}

//S ...
func (c *Cat) S() {
	fmt.Println("MIAU")
}

//Animal es exportado ...
type Animal interface {
	S()
}

func main() {

	gob.Register(&Dog{})
	gob.Register(&Cat{})

	var b bytes.Buffer

	enco := gob.NewEncoder(&b)
	deco := gob.NewDecoder(&b)

	var into Animal

	into = new(Dog)

	err := enco.Encode(&into)

	if err != nil {
		log.Fatal(err)
	}

	into = new(Cat)

	err = enco.Encode(&into)

	if err != nil {
		log.Fatal(err)
	}

	//Ahora vamos a decodificar

	var get Animal

	err = deco.Decode(&get)
	if err != nil {
		log.Fatal(err)
	}

	get.S()

	err = deco.Decode(&get)
	if err != nil {
		log.Fatal(err)
	}

	get.S()

}
