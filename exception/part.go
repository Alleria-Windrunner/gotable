package exception

import (
	"fmt"
)

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
	message := fmt.Sprintf("longLen is %d, shortLen is %d", long, short)
	err := &ColLenError{createBaseError(message), long, short}
	return err
}
