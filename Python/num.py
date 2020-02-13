#!/bin/python3
import time
import numpy as np

if __name__ == '__main__':
    arr=np.arange(10)
    print(arr)
    t0 = time.time()
    arr=np.zeros([5000,5000])
    t1 = time.time()

    arr2=[[0 for _ in range(5000)] for _ in range(5000)]
    t2 = time.time()
    print(str(t1-t0), str(t2-t1))
    arr[0][0]=1
    print(arr[0][0])