#include <stdio.h>
#include "msg_output.h"


int main(int argc, char **argv){
    GoInt x=12;
    GoInt y=23;
    printf("About to call a Go function!\n");
    printMessage();
    GoInt p=Multiply(x,y);
    printf("Product: %d\n",(int)p);
    printf("It worked!\n");
    return 0;
}