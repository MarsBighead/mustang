#!/usr/bin/perl


$n = <STDIN>;
chomp $n;
$arr_temp = <STDIN>;
@arr = split / /,$arr_temp;
chomp @arr;
my $n = scalar(@arr)-1;
for (my $i=$n;$i>=0;$i--){
   if ($i==0){
      print "$arr[$i]\n";
   } else{
      print "$arr[$i] "
   }
}

