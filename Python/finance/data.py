#!/usr/bin/python3

import numpy as np
import pandas_datareader as pdr
import matplotlib.pyplot as plt
import pandas as pd

goog=pdr.data.DataReader('GOOG', data_source='yahoo', start='3/14/2020',end='5/14/2020')
#print(goog)
goog['Log_Ret']=np.log(goog['Close']/goog['Close'].shift(1))
print(goog.index)
#print(goog['Log_Ret'].rolling(22).std())
goog['Volatility']=(goog['Log_Ret'].rolling(5).std())*np.sqrt(5)
print(goog)

#goog[['Close', 'Volatility']].plot(subplot=True, color='blue', figsize=(8,6))

goog[['Close', 'Volatility']].plot(subplots=True, color='green', figsize=(8,6))

#lt.show()
plt.savefig('google-1.png', format='png')