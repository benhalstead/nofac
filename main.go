package main

import "github.com/graniticio/granitic/v2"
import "nofac/bindings"

func main() {
	granitic.StartGranitic(bindings.Components())
}
