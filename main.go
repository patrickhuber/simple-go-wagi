package main

import "os"

func main() {
	println("content-type: text/plain;utf-8")
	println("")
	for _, e := range os.Environ() {
		println(e)
	}
}
