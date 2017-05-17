#!/usr/bin/perl
use strict;
use warnings;

# Enter your code here. Read input from STDIN. Print output to STDOUT
my $nm=<STDIN>;
chomp $nm;
my ($n,$m)=split / /,$nm;
#my @val=split (//,0 x ($n+1));
my @val=split (//,0 x $n);
for (my $r=0;$r<$m;$r++){
   my $tmp=<STDIN>;
   chomp ($tmp);
   my ($a, $b, $k)=split / /,$tmp;
   $val[$a-1]+=$k;
   if ($b<$n){
       $val[$b]-=$k;
   }
   print join(" ",@val)."\n";
}
my $max= $val[0];
my $sum= $val[0];
foreach (1..$n-1){
  $sum+=$val[$_];
  if ($max<$sum){
      $max=$sum;
  }
}
print "$max";
