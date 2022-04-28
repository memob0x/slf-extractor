package main

var global struct {
	CreateApp func()
}

func main() {
	global.CreateApp()
}
