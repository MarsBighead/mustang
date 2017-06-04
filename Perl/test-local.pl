#!/usr/bin/perl

#keywords our  
our $Scalar =1;                  #Global variable,  action scope is the package/file
sub Subroutine{  
    local $Scalar =2;            #local variable, action scope is in the subroutine 
    $Scalar +=1; 
    print $Scalar,"\n";
} 
&Subroutine;                     #输出3  
&Subroutine;                     #输出3   
print $Scalar."\n";              #输出1
