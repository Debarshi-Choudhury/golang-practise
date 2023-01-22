package main

import (
	exmpl "random_examples/random_examples_package"
)

/*
	This is an example of importing a package from the same module
*/

func main() {
	exmpl.SayHello()
	exmpl.ReadingFromChannelAfterBeingClosed()
	exmpl.SayBye()
}
