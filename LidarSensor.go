package main

var SensorAddr int = 0x62

type LidarSensor struct {
    i2cAddr     int
    bus         *I2cBus
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
