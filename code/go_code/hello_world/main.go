package main

import (
	"fmt"
)

func say(m string) {   // function, input only, still needs type
        fmt.Println(m) // now you can say() stuff in stead of fmt.Println() stuff
}

func main() {
	fmt.Println("hello world1")
	say("hello world2")
        fmt.Printf("%[1]v\n", "hello world3")
	print("hello stderr world\n")
}
