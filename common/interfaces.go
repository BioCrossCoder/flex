package common

type Iterator interface {
	Next() bool
	Value() any
	Pour() any
}
