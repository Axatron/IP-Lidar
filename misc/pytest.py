#!/usr/bin/python
import time
import fcntl
import os
import struct

samples = 101

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

def avg(vals):
    return sum(vals) / len(vals)


def mode(vals):
    v = {}
    for val in vals:
        if val not in v:
            v[val] = 1
        else:
            v[val] += 1

    results = sorted(v, key=lambda x: v[x], reverse = True)
    if results[0] == results[1]:
        print "DUP"
    return results[0]

def median(vals):
    v = vals[:]
    v.sort()
    return v[samples/2]

def midrange(vals):
    return (min(vals) + max(vals)) / 2
#def process(vals):
#    foo = vals[:]
#    av = avg(foo)
#    for each in foo:
#        if (av+1) < each:
#           foo.remove(each)
#        if (av-1) > each:
#           foo.remove(each)
#
#    return avg(foo)

#start = time.time()
idx = 0
count = 0
values = [0] * samples
while True:
    values[idx], eltime = readDistance()
    idx += 1
    if idx == samples:
        idx = 0
        count += 1
        if count == 1000: break
#        print round(process(values)), '\t', round(avg(values))
        print avg(values), '\t', mode(values), '\t', median(values), '\t', midrange(values), '\t', min(values), '\t', max(values)

