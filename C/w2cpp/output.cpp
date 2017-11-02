#include <iostream>

int sum();

int main(){
	std::cout<< "/*";
	std::cout<< "*/";
	//std::cout<< /* "*/" */;
	std::cout<< /* "*/" /* "/*" */;
	return sum();
}

int sum(){
    int sum=0,val=1;
    while(val<=10){
        sum+=val;
	val++;
    }
    std::cout <<"Sum of 1 to 10 inclusive is"
	      <<sum<<std::endl;
    return 0;
}
