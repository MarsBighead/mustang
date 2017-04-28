#!/usr/bin/perl
use warnings;
use strict;

# Task 
# Given an integer, , perform the following conditional actions:
# 
# If  is odd, print Weird
# If  is even and in the inclusive range of  to , print Not Weird
# If  is even and in the inclusive range of  to , print Weird
# If  is even and greater than , print Not Weird
# Complete the stub code provided in your editor to print whether or not  is weird.

my $N = <STDIN>;
chomp $N;

if ($N%2==1){
    print "Weird\n";
} elsif( $N%2==0 && $N>=2 && $N<=5 ){
    print "Not Weird\n";
}elsif( $N%2==0 && $N>=6 && $N<=20 ){
    print "Weird\n";
}elsif( $N%2==0 && $N>20 ){
    print "Not Weird\n";
}
