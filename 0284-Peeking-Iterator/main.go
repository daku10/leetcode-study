package main

/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type Iterator struct {
	Nums  []int
	point int
}

func (this *Iterator) hasNext() bool {
	return this.point != len(this.Nums)
}

func (this *Iterator) next() int {
	n := this.Nums[this.point]
	this.point++
	return n
}

type PeekingIterator struct {
	i          *Iterator
	buf        int
	bufHasNext bool
}

func Constructor(iter *Iterator) *PeekingIterator {
	bufHasNext := iter.hasNext()
	var buf int
	if bufHasNext {
		buf = iter.next()
	}
	return &PeekingIterator{
		i:          iter,
		buf:        buf,
		bufHasNext: bufHasNext,
	}
}

func (this *PeekingIterator) hasNext() bool {
	return this.bufHasNext
}

func (this *PeekingIterator) next() int {
	b := this.buf
	if this.i.hasNext() {
		this.buf = this.i.next()
	} else {
		this.bufHasNext = false
	}
	return b
}

func (this *PeekingIterator) peek() int {
	return this.buf
}
