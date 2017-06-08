**mustang**
------------
Partly exercise code of  mine, inlude but not limit Golang, Perl, PHP, Python, Nodejs, etc.

- 1 Generics    
use C++ after C++99 version   

```
- g++ -std=c++0x generics.cpp -o generics
```

- 2 Docker    
Build application with docker
```
docker run --rm \
   -v "$(pwd)":/home/hbu/mustang/docker/hello \ 
   -w "$(pwd)" golang:1.8.1 \
   go build -v
```
