using Interpolations,Plots

t = 0:.1:1
x = sin.(2π*t)
y = cos.(2π*t)
A = hcat(x,y)

itp = scale(interpolate(A, (BSpline(Cubic(Natural())), NoInterp()), OnGrid()), t, 1:2)

tfine = 0:.01:1
xs, ys = [itp(t,1) for t in tfine], [itp(t,2) for t in tfine]


scatter(x, y, label="knots")
plot!(xs, ys, label="spline")