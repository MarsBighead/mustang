#!/usr/bin/python3

from numpy import *

S0 = 100
K = 105
T = 1.0
r = 0.05
sigma = 0.2
I = 1000000
z = random.standard_normal(I)
ST = S0*exp((r-0.5*sigma**2)*T+sigma*sqrt(T)*z)
hT = maximum(ST-K,0)
C0 = exp(-r*T)*sum(hT)/I

print("Value of the European All Option %5.3f" % C0)