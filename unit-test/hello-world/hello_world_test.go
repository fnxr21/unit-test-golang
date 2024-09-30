package unittesthelloworld

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("IVAN")
	if result != "Hello IVAN" {
		panic("Result is not Hello Ivan")
	}
}
func TestHelloIvan(t *testing.T) {
	result := HelloWorld("IVAN")
	if result != "Hello IVAN" {
		panic("Result is not Hello Ivan")
	}
}
