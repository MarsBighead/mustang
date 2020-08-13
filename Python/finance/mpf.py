#!/usr/bin/python3

import mplfinance as mpf
import pandas_datareader as pdr
import matplotlib.pyplot as plt
import pandas as pd

stock='IBM'
#pdr.data.DataReader()
data= pdr.get_data_yahoo(stock, '2020/1/1', '2020/8/1') 
#pd.to_datetime(data['date'])

#print(data.index)
#print(data.columns)

#fig,ax=plt.subplots(figsize=(8, 5))
 
#fig.subplots_adjust(bottom=0.1) 
#s = mpf.make_mpf_style(base_mpf_style='classic',mavcolors=['lime','b','r'])
#mpf.plot(data,type='candle',  volume=True, figscale=1, mav=(7,11,15),style=s)
mpf.plot(data, type='candle',title=stock+' Inc',volume=True,mav=(20,40),
    scale_padding=0.5,
    figscale=1.5,style='charles')

'''
mpf.plot(data, type='candle', volume=True, title=stock+' Inc',
     figscale=1, 
     mav=(7,11,15),
     style=s)
'''
#mpf.plot(data, type='candle', mav=(3,6,9),volume=True,show_nontrading=True)

#mpf.make_marketcolors()

#plt.grid(True)
#ax.xaxis_date()
#ax.autoscale_view()
#plt.setp(plt.gca().get_xtickleables(), rotation=30)

#plt.show()
 
'''
import requests
import time
import pandas as pd
import matplotlib.pyplot as plt

def datetime_timestamp(dt):
     time.strptime(dt, '%Y-%m-%d %H:%M:%S')
     s = time.mktime(time.strptime(dt, '%Y-%m-%d %H:%M:%S'))
     return str(int(s))

s = requests.Session()

#Replace B=xxxx
cookies = dict(B='c650m5hchrhii&b=3&s=tk')

#Replace crumb=yyyy
crumb = 'NMhMTCv7QpM'
stock="VMW"

begin = datetime_timestamp("2020-01-01 09:00:00")
    
end = datetime_timestamp("2020-04-30 09:00:00")

r = s.get("https://query1.finance.yahoo.com/v7/finance/download/"+stock+"?period1="+begin+"&period2="+end+"&interval=1d&events=history&crumb="+crumb,cookies=cookies,verify=False)
   
filename=stock+'.csv'
f = open(filename, 'w')
f.write(r.text)
f.close()    
    
    
es = pd.read_csv(filename, index_col=0,parse_dates=True, sep=",", dayfirst=True)
    
data = pd.DataFrame({stock : es["Adj Close"][:]}) 
    
print(data.info())
    
data.plot(subplots=True, grid=True, style="b", figsize=(8, 6))
    
plt.show()
'''