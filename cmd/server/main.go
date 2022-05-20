package main

import "fmt"

//Run - si responsible for the instantiation and
//startup for our go application
func Run() error {
	fmt.Println("starting up our application")
	return nil
}

func main() {
	fmt.Println("go rest api")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
