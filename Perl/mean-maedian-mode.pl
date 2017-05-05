#!/usr/bin/perl
#In this challenge, we practice calculating the mean, median, and mode. 

my $n=<STDIN>;
chomp($n);
my $line=<STDIN>;
my @arr =split / /,$line;
my $sum=0;
for (my $i=0;$i<$n;$i++){
   for (my $j=$i+1;$j<$n;$j++){
       if ($arr[$j]<$arr[$i]){
          my $tmp=$arr[$i];
          $arr[$i]=$arr[$j];
          $arr[$j]=$tmp;
       }
   }
   $sum+=$arr[$i];
}
my $mean = $sum/$n;
my $median=0;
if ($n%2==0){
   $median=($arr[$n/2-1]+$arr[$n/2])/2;
}else{
   $median=$arr[int($n/2)];
}
printf("%.1f\n%.1f\n",$mean,$median);
my $mode=$arr[0];
my %hash=();
for (my $m=0;$m<$n;$m++){
    $hash{$arr[$m]}++;
}
my $count=$hash{$arr[0]};
for my $key( keys  %hash){
   if ($hash{$key}>$count || ($key<$mode && $hash{$key}==$count)){
        $mode=$key;
        $count =$hash{$key};
   }
}
print $mode;

