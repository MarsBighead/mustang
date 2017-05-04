#!/usr/bin/perl

# Enter your code here. Read input from STDIN. Print output to STDOUT

my $n=<STDIN>;
my $f = factorial($n);
print "$f\n";

sub factorial{
   my $n = shift;
   if ($n==1) {
      return $n;
   }else{
      return $n*factorial($n-1);
   }
}
