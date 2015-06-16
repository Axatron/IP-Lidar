package main

import (
	"syscall"
)

const (
	I2C_RETRIES = iota + 0x0701
	I2C_TIMEOUT
	I2C_SLAVE
	I2C_TENBIT
	I2C_FUNCS
	I2C_SLAVE_FORCE
	I2C_RDWR
	I2C_PEC
)

type I2cBus struct {
	devfd   int
	devpath string
}

func NewI2cBus(dev string) (*I2cBus, error) {
	i2c := &I2cBus{devpath: dev}
	err := i2c.Open()
	return i2c, err
}

func (self *I2cBus) Open() error {
	var err error
	self.devfd, err = syscall.Open(self.devpath, syscall.O_RDWR, 0777)
	return err
}

func (self I2cBus) SetAddr(addr int) error {
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(self.devfd), uintptr(I2C_SLAVE), uintptr(addr))
	if errno != 0 {
		return syscall.Errno(errno)
	}
	return nil
}

func (self I2cBus) WriteRegister(reg byte, data []byte) (int, error) {
	data = append([]byte{reg}, data...)
	return self.Write(data)
}

func (self I2cBus) ReadRegister(reg byte, buf []byte) (int, error) {
	size, err := self.Write([]byte{reg})
	if err != nil {
		return size, err
	}
	return self.Read(buf)
}

func (self I2cBus) Write(buf []byte) (int, error) {
	return syscall.Write(self.devfd, buf)
}

func (self I2cBus) Read(buf []byte) (int, error) {
	return syscall.Read(self.devfd, buf)
}
