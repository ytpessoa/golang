package main

import "runtime/debug"

func f10() {
	debug.PrintStack()
}
func f9() {
	f10()
}
func f8() {
	f9()
}
func main() {
	f8()
}
