package common

import "reflect"

var List = []reflect.Kind{reflect.Array, reflect.Slice}

var Sequence = append(List, reflect.String)

var IterableContainers = append(Sequence, reflect.Map)

const hashTableFillFactor = 0.8

const reHashThreshold = 6.5
