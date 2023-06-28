package mathsmod

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	result, err := Add(10, 20)
	if err != nil {
		t.Error("error")
	} else {
		fmt.Println(result)
	}

	// this test will fail
	result, err = Add(10, 0)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
