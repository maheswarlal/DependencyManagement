package main

import (
	"fmt"

	"github.com/maheswarlal/dog"
	"github.com/maheswarlal/puppy"
)

func main() {
	puppy.Taglal4()
	puppy.Taglal5()
	puppy.Taglal8()
	puppy.Taglal8()
	puppy.Fromlal1()
	puppy.Fromlal2()
	S1 := puppy.Bark()
	S2 := puppy.Barks()
	fmt.Println(S1)
	fmt.Println(S2)
	S3 := puppy.BigBark()
	fmt.Println(S3)

	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())
	fmt.Println(puppy.Bark())
	fmt.Println(dog.WhenGrownUp("Tom"))

}
