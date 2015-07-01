package main

import (
	"fmt"
)

var I2CBusPath string = "/dev/i2c-1"

func main() {
	i2c, err := NewI2cBus(I2CBusPath)
	if err != nil {
		panic(err)
	}

	sensor, err := NewLidarSensor(i2c)
	if err != nil {
		panic(err)
	}

	samples := NewSamples(11)
	for {
		distance, err := sensor.ReadDistance()
		if err != nil {
			panic(err)
		}
		samples.AddValue(distance)
		fmt.Println(samples)
		fmt.Println(samples.Mode())
	}
}
