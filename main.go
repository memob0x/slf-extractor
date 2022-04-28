package main

import (
	"fmt"
	"log"
)

var global struct {
	id int

	CreateApp func()
}

const INT_ID_BUILD_CLI = 1
const INT_ID_BUILD_GUI = 2

func main() {
	switch global.id {

	case INT_ID_BUILD_CLI:
		fmt.Printf("slf-extractor command line interface \n")

		break

	case INT_ID_BUILD_GUI:
		fmt.Printf("slf-extractor graphic user interface \n")

		break

	default:
		log.Fatal("err")
	}

	global.CreateApp()
}
