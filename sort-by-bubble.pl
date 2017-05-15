#!/usr/bin/perl

$n = <STDIN>;
chomp $n;
$a_temp = <STDIN>;
@a = split / /,$a_temp;
chomp @a;
# Write Your Code Here
my $c=0;
for (my $i=0;$i<scalar(@a);$i++){
   for (my $j=$i+1;$j<scalar(@a);$j++){
       if ( $a[$i]>$a[$j]){
          my $tmp=$a[$i];
          $a[$i]=$a[$j];
          $a[$j]=$tmp;
          $c++;
       }
   }
}

print "Array is sorted in $c swaps.\n";
print "First Element: ".$a[0]."\n";
print "Last Element: ".$a[scalar(@a)-1]."\n";
