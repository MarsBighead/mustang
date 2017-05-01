#!/usr/bin/perl
$n = <STDIN>;

chomp $n;
$arr_temp = <STDIN>;
@arr = split / /,$arr_temp;
chomp @arr;
$r = pickupPlusMinus(@arr);
printf("%.6f\n%.6f\n%.6f\n",$r->[2],$r->[1],$r->[0]);

sub pickupPlusMinus{
   my @a=@_;
   my $r=[0,0,0];
   foreach my $i(@a){
      if ($i<0){
         $r->[1]++;
      }elsif ($i>0){
         $r->[2]++;
      }else{
         $r->[0]++; 
      }
   }
   $r->[0]=$r->[0]/scalar(@a);
   $r->[1]=$r->[1]/scalar(@a);
   $r->[2]=$r->[2]/scalar(@a);
   return $r;
}

