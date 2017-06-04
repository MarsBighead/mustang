#!/usr/bin/perl

#keywords our  
our $Scalar =1;                  #, 作用域为包  
sub Subroutine{  
    our $Scalar =2;              #全局, 作用域为包  
    $Scalar +=1; 
    print $Scalar,"\n";
} 
&Subroutine;                     #输出3  
&Subroutine;                     #输出3   
print $Scalar."\n";                   #输出3  
