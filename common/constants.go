// Package common contains common constants, functions and errors used throughout the flex codebase.
package common

import "reflect"

// ArrayType incluedes array and slice.
var ArrayType = []reflect.Kind{reflect.Array, reflect.Slice}

// SequenceType includes array, slice and string.
var SequenceType = append(ArrayType, reflect.String)

// IterableContainers includes array, slice, string and map.
var IterableContainers = append(SequenceType, reflect.Map)

// hashTableLoadFactor is the preseted load factor of the hash table.
const hashTableLoadFactor = 0.8

// reHashThreshold is the threshold of the load factor to trigger rehashing.
const reHashThreshold = 6.5
