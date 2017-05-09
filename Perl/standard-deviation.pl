#!/usr/bin/perl
use strict;
use warnings;

# In this challenge, we practice calculating standard deviation. 
# Check out the Tutorial tab for learning materials and an instructional video!

my $n=<STDIN>;
chomp $n;
my $vals=<STDIN>;
chomp $vals;

my @arr=split / /,$vals;
my $sum=0;
for (my $i=0;$i<$n;$i++){
   $sum+=$arr[$i];
}
my $avg=$sum/$n;
$sum=0;
for (my $i=0;$i<$n;$i++){
   $sum+=($arr[$i]-$avg)**2;
}

my $sig=sqrt($sum/$n);
printf("%.1f",$sig);
