#!/usr/bin/perl

$n = <STDIN>;
chomp $n;

$s=sprintf("%b",$n);
print "$s\n";
@a=split //,"$s";
my ($i,$j)=(0,0);
for (my $x=0;$x<=scalar(@a);$x++){
print "$a[$i]\n";
   if ($a[$x]=="1"){
      $i++;
   }else{
      if ($i>$j){
         $j=$i;
      }
      $i=0;
   }
}
print "$i\n";
if ( $j==0 && $i>0){ 
   $j=$i;
}
print "$j";
