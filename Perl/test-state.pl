#!/usr/bin/perl
use strict;
use warnings;
use feature 'state';

# Please add use feature 'state'; in the head lines to enforce function state

#keywords my
my $Scalar =1;                #Private variable, action scope is the package/file
sub Subroutine{  
    state $Scalar =2;            #Private variable, action scope is in the subroutine 
    $Scalar +=1; 
    print $Scalar,"\n";
} 
&Subroutine;                     #输出3  
&Subroutine;                     #输出3   
print $Scalar."\n";              #输出1
