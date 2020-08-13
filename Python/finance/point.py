#!/usr/bin/python3

import numpy as np
import matplotlib.pyplot as plt

y=np.random.standard_normal((1000,2))
print(y)
c=np.random.randint(0,10,len(y))
plt.figure(figsize=(7,5))
#plt.plot(y[:,0], y[:,1],'ro', marker='o')
plt.scatter(y[:,0], y[:,1], c=c, marker='o')
plt.colorbar()
plt.grid(True)
plt.xlabel('1st')
plt.xlabel('2nd')
pic='Scatter Plot'
plt.title(pic)
plt.savefig(pic+'.png', format='png')
