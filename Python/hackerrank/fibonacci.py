def fibonacci1(n):
    if n==0:
        return 0
    if n==1:
        return 1
    return fibonacci(n-1)+fibonacci(n-2)

    # Write your code here.
def fibonacci(n):
    f1=0
    f2=1
    for i in range(n):
        f1,f2=f2, f1+f2
    return f1


n = 2
print(fibonacci(n))
print(fibonacci1(n))