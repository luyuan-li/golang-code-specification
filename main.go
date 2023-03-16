package main

import (
	"runtime"

	"github.com/golang-code-specification/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 8)
	cmd.Execute()
}
