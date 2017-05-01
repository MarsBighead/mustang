#!/usr/bin/perl 
use strict;
use warnings;

my $n = <STDIN>;
chomp $n;

while ($n>0){
   my $s = <STDIN>;
   chomp $s;
   my $r = getEvenAndOdd($s);
   print $r->[0]." ".$r->[1]."\n";
   $n--;
}


sub getEvenAndOdd{
   my $s = shift;
   my @a =  split "",$s;
   my $even,$odd=();
   for (my $i=0; $i<scalar(@a); $i++){
      if ($i%2==0){
         $even.=$a[$i];
      }else{
         $odd.=$a[$i];
      }
   }
   return [$even, $odd]
}
