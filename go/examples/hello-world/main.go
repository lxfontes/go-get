package main

import (
	"fmt"

	"go.wasmcloud.dev/go/component"
	"go.wasmcloud.dev/go/provider"
)

func main() {
	fmt.Println(provider.One)
	fmt.Println(component.Two)
}
