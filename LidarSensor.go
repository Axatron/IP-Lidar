package main

type LidarSensor struct {
        i2cAddr         int
        bus             *I2cBus
        Distance        uint16
}

func (self *LidarSensor) Init() {
        self.bus.Open()
        self.bus.SetAddr(self.i2cAddr)
}
