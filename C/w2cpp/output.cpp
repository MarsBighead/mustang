#include <iostream>

int sum();

int kickoff();

int main(){
    int num=10, *i;
    //i=&num;
    std::cout <<"uninitilze int i is "<<*i<<" address of num is "<<&num<<'\n';
    std::cout <<"\\115 is "<<'\115'<<'\n';
    std::cout <<"\\x4d is "<<'\x4d'<<'\n';
    std::cout <<"\\7 is "<<'\7'<<'\n';
    std::cout <<"\\7 is "<<'\7'<<'\n';
    std::cout <<"L'a' is "<<L'a'<<'\n';
    std::cout <<"u8\"hi!\" is "<<u8"hi!"<<'\n';
    std::cout <<"1E-3F is "<<1E+3F<<'\n';
    std::cout <<"\a1E-3F is "<<1E+3F<<'\n';
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

int kickoff(){
	std::cout<< "/*";
	std::cout<< "*/";
	//std::cout<< /* "*/" */;
	std::cout<< /* "*/" /* "/*" */;
	return sum();
}
