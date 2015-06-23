package main

import (
	"fmt"
	"math/rand"
)

type Samples struct {
	idx	int
	samps	[]uint16
}

func NewSamples(length int) (*Samples) {
	samp := new(Samples)
	samp.samps = make([]uint16, length)
	return samp
}

func (self *Samples) AddValue(val uint16) {
	if (self.idx == len(self.samps)) {
		self.idx = 0
	}

	self.samps[self.idx] = val
	self.idx++
}

func rand16() (uint16) {
	return uint16(rand.Int31() & 65535)
}

func main() {
	samples := NewSamples(15)
	for i := 0; i < 45; i++ {
		fmt.Println(samples.samps)
		samples.AddValue(rand16())
	}
}
