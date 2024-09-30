package unittesthelloworld

import (
	"fmt"
	"testing"
)

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

// this fail but still continue because use t.Fail()
// # go test -v  -run TestHelloWorldFail

func TestHelloWorldFail(t *testing.T) {
	result := HelloWorld("IVAN")
	if result != "Hello IVAN fail" {
		t.Fail()
	}
	fmt.Println("test still continue,right ")
}

// this fail and not  continue because use t.FailNow()
// # go test -v  -run TestHelloWorldFailNow

func TestHelloWorldFailNow(t *testing.T) {
	result := HelloWorld("IVAN")
	if result != "Hello IVAN fail" {
		t.FailNow()
	}
	fmt.Println("test not continue,right ")
}

//for note they send args or infor what error

// this fail and still  continue because use t.Error()
// # go test -v  -run TestHelloWorldError

func TestHelloWorldError(t *testing.T) {
	result := HelloWorld("IVAN")
	if result != "Hello IVAN fail" {
		t.Error("Result must be 'Hello Ivan'")
	}
	fmt.Println("test continue,right ")
}

// this fail and still  continue because use t.Error()
// # go test -v  -run TestHelloWorldFatal

func TestHelloWorldFatal(t *testing.T) {
	result := HelloWorld("IVAN")
	if result != "Hello IVAN fail" {
		t.Fatal("Result must be 'Hello Ivan'")
	}
	fmt.Println("test not continue,right ")
}
