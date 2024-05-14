package main

import (
	"examples/greetings"
	"fmt"
	"log"
)

func main() {
	// name, num, isNil := greetings.Hello("mock mock mock")
	// name2, num2, isNil2 := greetings.Hello("phayuphat")

	// if isNil2 != nil || isNil != nil {
	// 	log.Fatal(isNil2)
	// }
	// log.SetPrefix("greetings: ")
	// log.SetFlags(0)

	// greetingsV3.Greetings()
	// fmt.Println(value)
	// fmt.Println(quote.Go())
	// fmt.Println(name, "=======", num, "========", isNil)
	// fmt.Println(name2, "=====", num2, "============", isNil2)

	list_name := []string{
		"phayuphat",
		"phayu",
		"phat",
	}
	value, greetingsNil := greetings.Hellos(list_name)
	if greetingsNil != nil {
		log.Fatal(greetingsNil)
	}

	for _, el := range value {
		fmt.Println(el)
	}

	// fmt.Println(value)
	// desc --> catch error when have  errors then exist process

}
