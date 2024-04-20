package common

import "reflect"

var List = []reflect.Kind{reflect.Array, reflect.Slice}

var Sequence = append(List, reflect.String)

var IterableContainers = append(Sequence, reflect.Map)
