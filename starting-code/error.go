package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("hahahaha")
	}
	return arg + 3, nil
}

type argError struct {
	arg int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42{
		return -1, &argError{arg, "cannot work with it!"}
	}
	return arg + 3, nil
}

func main(){
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil{
			fmt.Println(e)
		}else{
			fmt.Println(r)
		}
	}
}