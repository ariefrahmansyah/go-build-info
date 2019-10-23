package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		// binary has not been built with module support
		panic("at the disco")
	}

	fmt.Printf("%+v", info)
}
