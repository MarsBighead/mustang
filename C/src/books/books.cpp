// ifstream constructor.
#include <iostream>     // std::cout
#include <fstream>
#include "book/Sales_item.h"
using namespace std;



int main()  
{ 
 
  /* ofstream: Stream class to write on files
    ifstream: Stream class to read from files
    fstream: Stream class to both read and write from/to files.*/
  cout<<__APPLE__<<endl;
  Sales_item book, book1;
  cin>>book>> book1;
  cout<<book+book1<<endl;

  return 0;
} 
