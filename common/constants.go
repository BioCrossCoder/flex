package common

import "reflect"

var Sequence = []reflect.Kind{reflect.Array, reflect.Slice, reflect.String}

var IterableContainers = append(Sequence, reflect.Map)
