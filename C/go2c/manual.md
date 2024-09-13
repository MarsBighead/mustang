# C中调用Go函数

## Go代码生成C共享库

注意：`//export PrintMessage`中，`//`与`export`间不可以有空格，否则无法生成c代码类型库，即`.h`文件。

```shell
go build -o msg_output.o -buildmode=c-shared msg_output.go
```

## C代码调用共享库编译

```shell
hbu@Pauls-MacBook-Air go2c % gcc -o usego  usego.c ./msg_output.o
hbu@Pauls-MacBook-Air go2c % ./usego                             
About to call a Go function!
A go funcation!
Product: 276
It worked!
```
