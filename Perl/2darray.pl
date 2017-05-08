#!/usr/bin/perl
#Today, we're building on our knowledge of Arrays by adding another dimension. Check out the Tutorial tab for learning materials and an instructional video!


$arr_i = 0;
@arr = ();
while($arr_i < 6){
   my $arr_temp = <STDIN>;
   my @arr_t = split / /,$arr_temp;
   chomp @arr_t;
   push @arr,\@arr_t;
   $arr_i++;
}
my $sum=0;

for (my $i=1;$i<=4;$i++){
   for (my $j=1;$j<=4;$j++){
      my $sum1=$arr[$i-1][$j-1]+$arr[$i-1][$j]+$arr[$i-1][$j+1];
      my $sum2=$arr[$i][$j];
      my $sum3=$arr[$i+1][$j-1]+$arr[$i+1][$j]+$arr[$i+1][$j+1];
      my $tmpSum=$sum1+$sum2+$sum3;
      if ($i==1 && $j==1){
         $sum =$tmpSum;  
      }else {
         if ($tmpSum > $sum ){
            $sum =$tmpSum; 
         }
      }
   }
}

print "$sum";

