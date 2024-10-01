package unittest

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
