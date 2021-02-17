package main

import (
	f "fmt"
	"runtime"
)

/*
main function
Go executes this program using this fucntion.
There shoudl only be one main file in a main package.
Executable programs are called commands.
*/
func main() {
	f.Println("scope")
	f.Println(runtime.NumCPU())
}
