package common

import "errors"

var ErrNotFunc = errors.New("the input parameter is not a function")

var ErrIllegalParamCount = errors.New("illegal number of input parameters")

var ErrIllegalReturnCount = errors.New("illegal number of return values")

var ErrNotIterable = errors.New("the input parameter is not iterable")

var ErrNotSeq = errors.New("the input parameter is not a sequence (slice/array/string)")

var ErrNotJudgeFunc = errors.New("the input func is not a function for judge")

var ErrNoReturn = errors.New("the input function has no return values")

var ErrInvalidCapacity = errors.New("the capacity must be an positive integer")

var ErrNotBool = errors.New("the value is not a boolean value")

var ErrNotMap = errors.New("the input parameter is not a map")

var ErrZeroStep = errors.New("the step cannot be zero")

var ErrNotList = errors.New("the input parameter must be a slice or array")

var ErrInvalidRange = errors.New("the range is invalid with the step: [start < end and step < 0] or [start > end and step > 0]")

var ErrOutOfRange = errors.New("the index is out of range")

var ErrListLengthMismatch = errors.New("the length of the input lists are not equal")

var ErrTooManyArguments = errors.New("too many arguments")

var ErrEmptyList = errors.New("the input list is empty")
