package mathsmod

import (
	"errors"
	"time"
)

// create a custom error model
type DivisionByZero struct {
	msg      string
	errCode  int
	funcName string
}

// implement the Error method
func (derr *DivisionByZero) Error() string {
	return derr.msg
}

func Add(no1 int, no2 int) (int, error) {
	if no1 == 0 {
		return 0, errors.New("no1 cannot be zero")
	} else if no2 == 0 {
		return 0, errors.New("no2 cannot be zero")
	} else {
		return no1 + no2, nil
	}
}

func Sub(no1 int, no2 int) (res int, err error) {
	if no1 == 0 {
		return 0, errors.New("no1 cannot be zero")
	} else if no2 == 0 {
		return 0, errors.New("no2 cannot be zero")
	} else {
		return no1 - no2, nil
	}
}

func Div(no1 int, no2 int) (res int, err *DivisionByZero) {
	time.Sleep(100)
	if no1 == 0 {
		return 0, &DivisionByZero{msg: "cannot divide by zero", errCode: 10, funcName: "mathsmod.div()"}
	} else if no2 == 0 {
		return 0, &DivisionByZero{"no2 cannot divide by zero", 10, "test"}
	} else {
		return no1 / no2, nil
	}
}
