package main

import (
	"math"
)

type Samples struct {
	idx	int
	length	int
	samps	[]uint16
}

func NewSamples(length int) (*Samples) {
	samp := new(Samples)
	samp.samps = make([]uint16, length)
	samp.length = length
	return samp
}

func (self *Samples) Total() (uint) {
	var total uint = 0
	for _, i := range self.samps {
		total += uint(i)
	}
	return total
}

func (self *Samples) Avg() (uint16) {
	bar := float64(self.Total()) / float64(self.length)
	return uint16(math.Floor(bar + 0.5))
}

func (self *Samples) Mode() (uint16) {
	values := make(map[uint16]int)
	var max int
	var best uint16

	for _, i := range self.samps {
		_, ok := values[i]
		if !ok {
			values[i] = 1
		} else {
			values[i] += 1
		}
	}

	for k, v := range values {
		if v > max {
			best = k
			max = v
		}
	}
	return best
}

func (self *Samples) AddValue(val uint16) {
	if (self.idx == len(self.samps)) {
		self.idx = 0
	}
	self.samps[self.idx] = val
	self.idx++
}

/*
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
*/
