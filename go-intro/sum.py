#!/usr/bin/env python -u
# encoding: utf-8

def tortoise(a):
    sum = 0
    for i in range(0, len(a)):
      sum += a[i]
    return sum

print tortoise([10,45,99, 1000000])
