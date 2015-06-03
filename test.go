package main

import (
	"fmt"
	"syscall"
	)

var I2C_SLAVE	int = 0x0703
var I2CBusPath	string = "/dev/i2c-1"
var SensorAddr	int = 0x62

func main() {
	devfd, err := syscall.Open(I2CBusPath, syscall.O_RDWR, 0777)
	if (err !=  nil) {panic(err)}

	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(devfd), uintptr(I2C_SLAVE), uintptr(SensorAddr))
	if (errno != 0) {panic(syscall.Errno(errno))}

	_, err = syscall.Write(devfd, []byte("\x00\x04"))
	if (err !=  nil) {panic(err)}

	for {
		_, err = syscall.Write(devfd, []byte("\x8f"))
		if (err !=  nil) {
			errno, ok := err.(syscall.Errno)
			if (!ok) {
				panic(err)
			}
			if (errno == 5) {
				continue
			} else {
				panic(err)
			}
		} else {
			break
		}
	}

	response := make([]byte, 2)
	_, err = syscall.Read(devfd, response)
        if (err !=  nil) {panic(err)}

	fmt.Println(response)

}
