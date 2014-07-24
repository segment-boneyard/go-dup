package dup

import "github.com/bmizerany/assert"
import "testing"
import "bufio"
import "fmt"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func TestDup(t *testing.T) {
	r, _, err := Dup(1, "stdout")
	check(err)

	go func() {
		fmt.Printf("testing\n")
	}()

	buf := bufio.NewReader(r)
	line, err := buf.ReadString('\n')
	check(err)

	assert.Equal(t, "testing\n", line)
}
