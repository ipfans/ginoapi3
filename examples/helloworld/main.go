package main

import (
	"github.com/ipfans/ginoapi3"
)

func main() {
	e := ginoapi3.Default()
	e.Info(ginoapi3.NewInfo("Hello World", "1.0.0"))
	e.Run()
}
