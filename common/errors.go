package common

import "errors"

// ErrNotFunc is an error indicating that the input parameter is not a function as expected.
var ErrNotFunc = errors.New("the input parameter is not a function")

// ErrUnexpectedParamCount is an error indicating that the number of input parameters is unexpected.
var ErrUnexpectedParamCount = errors.New("unexpected number of input parameters")

// ErrUnexpectedReturnCount is an error indicating that the number of return values is unexpected.
var ErrUnexpectedReturnCount = errors.New("unexpected number of return values")

// ErrNotIterable is an error indicating that the input parameter is not iterable as expected.
var ErrNotIterable = errors.New("the input parameter is not iterable")

// ErrNotSeq is an error indicating that the input parameter is not a slice, array or string as expected.
var ErrNotSeq = errors.New("the input parameter is not a slice, array or string")

// ErrNotJudgeFunc is an error indicating that the input parameter is not a function returning a boolean value as expected.
var ErrNotJudgeFunc = errors.New("the input parameter is not a function returning a boolean value")

// ErrInvalidCapacity is an error indicating that the capacity is not a positive integer as expected.
var ErrInvalidCapacity = errors.New("the capacity must be a positive integer")

// ErrZeroStep is an error indicating that the step is zero while it is expected to be a non-zero integer.
var ErrZeroStep = errors.New("the step cannot be zero")

// ErrNotList is an error indicating that the input parameter is not a slice or array as expected.
var ErrNotList = errors.New("the input parameter must be a slice or array")

// ErrInvalidRange is an error indicating that the range is invalid with the step.
var ErrInvalidRange = errors.New("the range is invalid with the step: [start < end and step < 0] or [start > end and step > 0]")

// ErrOutOfRange is an error indicating that the index is out of range.
var ErrOutOfRange = errors.New("the index is out of range")

// ErrListLengthMismatch is an error indicating that the length of the input lists are not fully the same as expected.
var ErrListLengthMismatch = errors.New("the length of the input lists are not equal")

// ErrTooManyArguments is an error indicating that there are more arguments than expected.
var ErrTooManyArguments = errors.New("too many arguments")

// ErrEmptyList is an error indicating that the input list is empty while it is expected to have at least one element.
var ErrEmptyList = errors.New("the input list is empty")

// ErrEmptySet is an error indicating that the input set is empty while it is expected to have at least one element.
var ErrEmptySet = errors.New("the input set is empty")

// ErrEmptyDict is an error indicating that the input dict is empty while it is expected to have at least one key-value pair.
var ErrEmptyDict = errors.New("the input dict is empty")

// ErrKeyNotFound is an error indicating that the key is not found in the dict as expected.
var ErrKeyNotFound = errors.New("the key is not found in the dict")
