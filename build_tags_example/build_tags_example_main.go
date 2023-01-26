package main

import (
	"fmt"
)

var features = []string{
	"Free feature #1",
	"Free feature #2",
}

func main() {
	for _, feature := range features {
		fmt.Println(">>", feature)
	}
}

/*
If you build like this:
$ go build
=> you will get only the free features

If you build like this:
$ go build -tags "pro"
=> you will get the free and pro features

If you build like this:
$ go build -tags "enterprise"
=> you will get the free, pro and enterprise features

*/
