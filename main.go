package main

import (
    "fmt"
)

var I2C_SLAVE	int = 0x0703
var I2CBusPath	string = "/dev/i2c-1"
var SensorAddr	int = 0x62

func main() {
	i2c, err := NewI2cBus(I2CBusPath)
	if err != nil {
        panic(err)
    }
    fmt.Println(i2c)
}
