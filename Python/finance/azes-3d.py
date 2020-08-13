#!/usr/bin/python3

import numpy as np
import pandas_datareader as pdr
import matplotlib.pyplot as plt
import pandas as pd
from mpl_toolkits.mplot3d import Axes3D

#quotes=pdr.data.DataReader('^GDAXI', data_source='yahoo', start='3/14/2020',end='5/14/2020')
strike=np.linspace(50, 150, 24)
ttm=np.linspace(0.5, 2.5, 24)
strike, ttm=np.meshgrid(strike, ttm)
print(strike)
iv=(strike-100)**2/(100*strike)/ttm


fig = plt.figure(figsize=(9, 6))
ax = fig.gca(projection='3d')
surf = ax.plot_surface(strike, ttm, iv, rstride=2, cstride=2, linewidth=0.5,
   cmap=plt.cm.rainbow,
   antialiased=True)

ax.set_xlabel('strike')
ax.set_ylabel('time-to-maturity')
ax.set_zlabel('implied volatility')

fig.colorbar(surf, shrink=0.5, aspect=5)
plt.show()
'''
fig = plt.figure(figsize=(9, 6))

ax = fig.add_subplot(111, projection='3d')

ax.view_init(30,60)

ax.scatter(strike, ttm, iv, zdir='z', s=25, c='b', marker='^')
x.set_xlabel('strike')
ax.set_ylabel('time-to-maturity')
ax.set_zlabel('implied volatility')

plt.show()
'''