package main

import "fmt"

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

func main() {
	samples := NewSamples(50)
	for i := 0; i < 50; i++ {
		fmt.Println(samples)
		samples.AddValue(20)
	}

	for i := 0; i < 50; i++ {
                fmt.Println(samples)
                samples.AddValue(300)
        }
}
