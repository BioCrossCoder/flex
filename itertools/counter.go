package itertools

type Counter struct {
	start   int
	step    int
	value   int
	begin   bool
	reverse bool
}

func NewCounter(start, step int, reverse bool) *Counter {
	return &Counter{
		start:   start,
		step:    step,
		value:   0,
		begin:   false,
		reverse: reverse,
	}
}

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

func (iter *Counter) Reset() {
	iter.value = 0
	iter.begin = false
}

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
