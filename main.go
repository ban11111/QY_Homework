package main

import (
	"QY_Homework/router"
	"flag"
	"QY_Homework/tools"
)

var Env string

func main() {
	env := flag.String("env", "test", "test, dev")
	flag.Parse()
	tools.ENV = *env
	router.Start_Server()
}