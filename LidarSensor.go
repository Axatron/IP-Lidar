package main

type LidarSensor struct {
    i2cAddr     int
    bus         *I2cBus
    Distance    uint16
}
