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

- 2 ls SELinux parameter
 SELinux options:

       --lcontext
              Display security context.   Enable -l. Lines will probably be too wide for most displays.

       -Z, --context
              Display security context so it fits on most displays.  Displays only mode, user, group, security context and file name.

       --scontext
              Display only security context and file name.

       --help display this help and exit

       --version
              output version information and exit

