package itertools

// Counter is an iterator that counts in a specific step, it can be used to generate a sequence of numbers, or to iterate over a range of numbers.
type Counter struct {
	start   int
	step    int
	value   int
	begin   bool
	reverse bool
}

// NewCounter creates a new Counter with the given start, step and a specified order.
func NewCounter(start, step int, reverse bool) *Counter {
	return &Counter{
		start:   start,
		step:    step,
		value:   0,
		begin:   false,
		reverse: reverse,
	}
}

// Count return the next expected value through counting.
func (iter *Counter) Count() int {
	if !iter.begin {
		iter.begin = true
		iter.value = iter.start
	} else if iter.reverse {
		iter.value -= iter.step
	} else {
		iter.value += iter.step
	}
	return iter.value
}

// Reset the iterator to its initial state.
func (iter *Counter) Reset() {
	iter.value = 0
	iter.begin = false
}

// Jump will do Count n times and return the new value.
func (iter *Counter) Jump(n int) int {
	if !iter.begin {
		iter.begin = true
		iter.value = iter.start
		n--
	}
	if iter.reverse {
		iter.value -= n * iter.step
	} else {
		iter.value += n * iter.step
	}
	return iter.value
}
