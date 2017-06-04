#!/usr/bin/perl

#keywords my
my $Scalar =1;                #Private variable, action scope is the package/file
sub Subroutine{  
    my $Scalar =2;            #Private variable, action scope is in the subroutine 
    $Scalar +=1; 
    print $Scalar,"\n";
} 
&Subroutine;                     #输出3  
&Subroutine;                     #输出3   
print $Scalar."\n";              #输出1
