package main

import (
	"math"
	"sort"
)

type SamplesSlice []uint16
func (p SamplesSlice) Len() int           { return len(p) }
func (p SamplesSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p SamplesSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

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

func (self *Samples) Sorted() ([]uint16) {
	temp := make([]uint16, self.length)
	copy(temp, self.samps)
	sort.Stable(SamplesSlice(temp))
	return temp
}


func (self *Samples) Sum() (uint) {
	var total uint = 0
	for _, i := range self.samps {
		total += uint(i)
	}
	return total
}

func (self *Samples) Mean() (uint16) {
	bar := float64(self.Sum()) / float64(self.length)
	return uint16(math.Floor(bar + 0.5))
}

func (self *Samples) Mode() (uint16) {
	values := make(map[uint16]int)
	var max int
	var best uint16

	for _, i := range self.samps {
		values[i]++
	}

	for k, v := range values {
		if v > max {
			best, max = k, v
		}
	}
	return best
}

func (self *Samples) Median() (uint16) {
	sorted := self.Sorted()
	return sorted[self.length / 2]
}

func (self *Samples) Midrange() (uint16) {
	sorted := self.Sorted()
	return (sorted[0] + sorted[self.length-1]) / 2
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
