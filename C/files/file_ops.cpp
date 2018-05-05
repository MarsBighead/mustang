// ifstream constructor.
#include <iostream>     // std::cout
#include <fstream>      // std::ifstream
#include <errno.h>
using namespace std;

int input (string filename,int argc,char *argv[]);
int output (string filename);
int main(int argc,char *argv[])  
{
  if (argc<=1){
    cerr <<"Error! No parameter input, exiting..."<<endl;
    return 1;
  } 
  string filename=argv[1];
  cout <<"Step1: start input data into file: "<<filename<<endl;
  input(filename, argc, argv);
  cout <<"Step2: start output data into file: "<<filename<<endl;
  output(filename);
  return 0;
} 

int output (string filename){
  ifstream ifs (filename, ifstream::out);
  char c = ifs.get();

  while (ifs.good()) {
    cout << c;
    c = ifs.get();
  }

  ifs.close();
  return 0;
}

int input (string filename,int argc,char *argv[]){
  ofstream f1; 
  f1.open(filename);
  for (int i=1;i<argc;i++){
    cout <<"argv is: "<<argv[i]<<endl;
    f1<<"argv is: "<<argv[i]<<endl;
  }
  f1.close();
  return 0;
}