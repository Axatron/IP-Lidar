package main

import (
//    "syscall"
    "fmt"
)

var I2C_SLAVE	int = 0x0703
var I2CBusPath	string = "/dev/i2c-1"
var SensorAddr	int = 0x62

type I2cBus struct {
	devfd		int
	devpath		string
}

func (self *I2cBus) Open() error {
	var err error
	self.devfd, err = syscall.Open(self.devpath, syscall.O_RDWR, 0777)
	return err
}

func (self I2cBus) SetAddr(addr int) error {
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(self.devfd), uintptr(I2C_SLAVE), uintptr(addr))
	return syscall.Errno(errno)
}

func (self I2cBus) Write(buf []byte) (int, error) {
	return syscall.Write(self.devfd, buf)
}

func (self I2cBus) Read(buf []byte) (int, error) {
	return syscall.Read(self.devfd, buf)
}

type LidarSensor struct {
	i2cAddr		int
	bus		*I2cBus
	Distance	uint16
}

func (self *LidarSensor) Init() {
	self.bus.Open()
	self.bus.SetAddr(self.i2cAddr)
}


