package main

import (
	"syscall"
)

type I2cBus struct {
	devfd	int
	devpath	string
}

func NewI2cBus(dev string) (I2cBus, error) {
	i2c := &I2cBus{devpath: dev}
	err := i2c.Open()
	return *i2c, err
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
