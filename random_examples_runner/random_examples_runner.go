package main

//remember, you do not import modules,you import packages inside modules
import (
	exmpl "random_examples/random_examples_package"
)

func main() {
	exmpl.SayHello()
	exmpl.ReadingFromChannelAfterBeingClosed()
	exmpl.SayBye()
}
