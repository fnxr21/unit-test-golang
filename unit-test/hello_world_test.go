package unittest

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)







// ========================== unit test basic
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
	if result != "Hello IVAN" {
		t.Fail()
		fmt.Println("test still continue,right ")
	}
}

// this fail and not  continue because use t.FailNow()
// # go test -v  -run TestHelloWorldFailNow

func TestHelloWorldFailNow(t *testing.T) {
	result := HelloWorld("IVAN")
	if result != "Hello IVAN" {
		t.FailNow()
		fmt.Println("test not continue,right ")
	}
}

//for note they send args or infor what error

// this fail and still  continue because use t.Error()
// # go test -v  -run TestHelloWorldError

func TestHelloWorldError(t *testing.T) {
	result := HelloWorld("IVAN")
	if result != "Hello IVAN" {
		t.Error("Result must be 'Hello Ivan'")
		fmt.Println("test continue,right ")
	}
}

// this fail and still  continue because use t.Error()
// # go test -v  -run TestHelloWorldFatal

func TestHelloWorldFatal(t *testing.T) {
	result := HelloWorld("IVAN")
	if result != "Hello IVAN" {
		t.Fatal("Result must be 'Hello Ivan'")
		fmt.Println("test not continue,right ")
	}
}

// ========================== unit test testify
// # go test -v  -run=TestHelloRequire

func TestHelloRequire(t *testing.T) {
	result := HelloWorld("IVAN")
	assert.Equal(t, "Hello IVAN", result, "Result must be 'Hello IVAN'")
	//this will stop if failed and not continue
	fmt.Println("TestHelloAssertion with require done")
}

// # go test -v  -run TestHelloAssertion
func TestHelloAssertion(t *testing.T) {
	result := HelloWorld("IVAN")
	assert.Equal(t, "Hello IVAN", result, "Result must be 'Hello IVAN'")
	fmt.Println("TestHelloAssertion with asserth done")
}

// ========================== skip
// skip ini hanya dalam keadaan tertentu
func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("cant not run on mac os")
	}

	//sample unittest
	result := HelloWorld("IVAN")
	assert.Equal(t, "Hello IVAN", result, "Result must be 'Hello IVAN'")
}

// ============================ test main
func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")
	//runing all testing
	m.Run()
	// after
	fmt.Println("AFTER UNIT TEST")
}

// ================= sub test

// go test -v -run=TestSubTest

// run sub test function spesific
// go test -v -run=TestSubTest/Fandi
func TestSubTest(t *testing.T) {
	t.Run("Fandi", func(t *testing.T) {
		result := HelloWorld("Fandi")
		assert.Equal(t, "Hello Fandi", result, "Result must be 'Hello Fandi'")
	})
	t.Run("Nur", func(t *testing.T) {
		result := HelloWorld("Nur")
		assert.Equal(t, "Hello Nur", result, "Result must be 'Hello Nur'")
	})
}

// =================== table test

func TestTableHelloWorld(t *testing.T) {

	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Fandi",
			request:  "Fandi",
			expected: "Hello Fandi",
		},
		{
			name:     "Nur",
			request:  "Nur",
			expected: "Hello Nur",
		},
	}
	// looping

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}

}
