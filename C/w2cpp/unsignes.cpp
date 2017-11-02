#include <iostream>

int compute_unsign();

int main(){
    compute_unsign();
}

int compute_unsign(){
    unsigned u1 = 42, u2 =10;
    std::cout <<"u1-u2="<<u1 - u2<<std::endl;
    std::cout <<"u2-u1="<<u2 - u1<<std::endl;
    int i=10, i2=42;
    std::cout <<"i-i2="<<i - i2<<std::endl;
    std::cout <<"i2-i="<<i2 - i<<std::endl;
    std::cout <<"i-u="<<i - u1<<std::endl;
    std::cout <<"u-i="<<u1 - i<<std::endl;
    for (int i =10;i>=0;i--)
	    std::cout <<"type of int "<< i << std::endl;
    for (unsigned u = 11; u > 0;){
	u--; 
        std::cout <<"type of unsigned "<< u << std::endl;
    }
    unsigned u =11;
    while(u >0){
	    u--;
	    std::cout <<"type of unsigned "<< u << std::endl;
    }
    return 0;
}
