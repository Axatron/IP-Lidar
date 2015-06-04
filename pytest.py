#!/usr/bin/python
import time
import fcntl
import os
import struct

samples = 100

I2C_SLAVE = 0x0703
i2cdevfd = os.open('/dev/i2c-1', os.O_RDWR)
fcntl.ioctl(i2cdevfd, I2C_SLAVE, 0x62)

def readDistance():
    os.write(i2cdevfd, '\x00\x04')

    count = 0
    start = time.time()
    while True:
        try:
            count +=1
            os.write(i2cdevfd, '\x8f')
            break
        except OSError:
            continue

    val = os.read(i2cdevfd, 2)
#    print repr(val)
    return (struct.unpack('>H', val)[0], time.time() - start)

start = time.time()
idx = 0
values = [0] * samples
while True:
    values[idx], eltime = readDistance()
    idx += 1
    if idx == samples:
        idx = 0
        print sum(values) / float(samples), time.time() - start
        start = time.time()
