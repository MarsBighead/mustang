# mustang C/C++

------------
C/C++ learning code part, and it's aiming at learning some C/C++ features to make me know more about programming method.

1. Connect PostgreSQL with libpq

C::database, compile sample:

```Shell
gcc -I/usr/local/opt/postgresql/include  -L//usr/local/opt/postgresql/lib -lpq ping.cpp
```

2.Build count sales book data

```shell
[hbu@hbu C]$ g++ -std=c++0x  books/sum.cpp
[hbu@hbu C]$ ./a.out
0-201-78345-x 3 20.00
0-201-78345-x 2 25.00
0-201-78345-x 5 110 22
```