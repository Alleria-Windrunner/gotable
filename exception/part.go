package exception

import "fmt"

type PartNumberError struct {
	*baseError
}

func PartNumber(partLen int) *PartNumberError {
	err := &PartNumberError{createBaseError("Part length now is " + fmt.Sprint(partLen))}
	return err
}

type ColLenError struct {
	*baseError
	longLen  int
	shortLen int
}

func ColLen(long int, short int) *ColLenError {
	err := &ColLenError{createBaseError("longLen is %d, shortLen is %d"), long, short}
	return err
}
