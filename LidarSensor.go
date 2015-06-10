package main

var SensorAddr int = 0x62

type LidarSensor struct {
	i2cAddr int
	bus     *I2cBus
	//    Distance    uint16
}

func NewLidarSensor(i2cbus *I2cBus) (*LidarSensor, error) {
	sensor := &LidarSensor{bus: i2cbus, i2cAddr: SensorAddr}
	err := sensor.bus.SetAddr(sensor.i2cAddr)
	if err != nil {
		return nil, err
	}
	return sensor, nil
}

func (self LidarSensor) ReadDistance(buf []byte) (uint16, error) {
	var result uint16
	buf = make([]byte, 2)

	self.bus.WriteRegister(0x00, []byte{0x04})
	_, err := self.bus.ReadRegister(0x8f, buf)
	if err != nil {
		panic(err)
	}

	result = (uint16(buf[0]) << 8) + uint16(buf[1])
	return result, nil
}
/*
	var max_ioerror int = 10
	for {
		len, err := self.bus.Read(buf)
		if (err !=  nil) {
			errno, ok := err.(syscall.Errno)
			if (!ok) {
				panic(err)
			}
			if (errno == 5) {
				if max_ioerror <= 0 {
					panic(err)
				}
				max_ioerror -= 1
				continue
			} else {
				panic(err)
			}
		} else {
			break
		}
	}
	return len
}
*/
