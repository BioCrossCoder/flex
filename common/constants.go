package common

import "reflect"

var ArrayType = []reflect.Kind{reflect.Array, reflect.Slice}

var SequenceType = append(ArrayType, reflect.String)

var IterableContainers = append(SequenceType, reflect.Map)

const hashTableFillFactor = 0.8

const reHashThreshold = 6.5
