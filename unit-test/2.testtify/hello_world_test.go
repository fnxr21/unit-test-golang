package unittest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ========================== unit test testify
// # go test -v  -run=TestHelloRequire

func TestHelloRequire(t *testing.T) {
	result := HelloWorld("IVAN")
	require.Equal(t, "Hello IVAN", result, "Result must be 'Hello IVAN'")
	//this will stop if failed and not continue
	fmt.Println("TestHelloAssertion with require done")
}

// # go test -v  -run TestHelloAssertion
func TestHelloAssertion(t *testing.T) {
	result := HelloWorld("IVAN")
	assert.Equal(t, "Hello IVAN", result, "Result must be 'Hello IVAN'")
	fmt.Println("TestHelloAssertion with asserth done")
}
