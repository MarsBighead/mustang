#include <iostream>

using namespace std;
int i=42;
int main(){
    cout <<"'a' is ASCII "<<'a'<<"\n";
    cout <<"\"a\" is string with \'\\0\' "<<"a"<<"\n";
    cout <<"L'a' is Unicode a, 2 bytes "<<L'a'<<"\n";
    cout <<"L\"a\" is Unicode string, 4 bytes "<<L"a"<<"\n";

    cout <<"10:10u:10L:10uL:012:0xC "<<10<<":"; 
    cout <<10u<<":";
    cout <<10L<<":";
    cout <<10uL<<":";
    cout <<012<<":";
    cout <<0xC<<"\n";

    cout <<"3.14:3.14f:3.14L "<<3.14<<":"; 
    cout <<3.14f<<":";
    cout <<3.14L<<"\n";
    //int month = 09; //invalid
    //int day = 07;
    //cout <<"month:"<<month<<"day"<<"\n";
    std::string global_str;
    cout<<global_str<<"| is\n";
    //int catch-22  //invalid
    double Double =3.14;
    //int 1_or_2=1;  //invalid
    int i = 100;
    int j = i;
    cout <<"j is "<<j<<"\n";
    int &refj = j;
    cout <<"&refj is "<<refj<<"\n";
    j=::i;
    cout <<"j is "<<j<<"\n";
    cout <<"&refj is "<<refj<<"\n";
    int x, &ri=x;
    x=5;ri=10;
    cout <<"x is "<<x<<" ri is "<<ri<<endl;
    int *pi=&x;
    if( pi == nullptr ){
        cout <<"nullptr pointor *pi is "<<pi<<endl;
    }else{
        cout <<"value of pointor *pi is "<<*pi<<endl;
    }
}
