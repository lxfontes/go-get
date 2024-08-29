package main

import (
	"fmt"

	"go.wasmcloud.dev/component"
	"go.wasmcloud.dev/provider"
)

func main() {
	fmt.Println(provider.One)
	fmt.Println(component.Two)
}
