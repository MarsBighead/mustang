#!/usr/bin/perl

# In this challenge, we practice calculating the interquartile range. 
# We recommend you complete the Quartiles challenge before attempting this problem.
# https://www.hackerrank.com/challenges/s10-interquartile-range

my $n=<STDIN>;
chomp $n;
my $vals=<STDIN>;
chomp $vals;
my @arr=split / /,$vals;
my $count_vals=<STDIN>;
chomp $count_vals;
my @arrc=split / /,$count_vals;

my @arr_set=();
for (my $i=0;$i<$n;$i++){
  for (my $j=0;$j<$arrc[$i];$j++){
     push (@arr_set, $arr[$i]);
  }
   
}
my @sort_set = sort{$a<=>$b} @arr_set;
my $alen=scalar(@sort_set);
my $pos =getMedianPos($alen);
print "len $alen\n";
my ($q1,$q3)=();
if (scalar(@$pos)==1){
   if ($pos->[0]%2==0){
      my $pq1=$pos->[0]/2;
      $q1=($sort_set[$pq1-1]+$sort_set[$pq1])/2;
      my $pq3=$alen-$pq1-1;
      $q3=($sort_set[$pq3+1]+$sort_set[$pq3])/2;
      
   }else{
      my $pq1=int($pos->[0]/2);
      $q1=$sort_set[$pq1-1];
      my $pq3=$alen-$pq1;
      $q3=$sort_set[$pq3];
      print "pos $pq1 $pq3 |val  $q1 $q3\n"
   }
} else{
   if ($pos->[1]%2==0){
      my $pq1=$pos->[1]/2;
      $q1=($sort_set[$pq1-1]+$sort_set[$pq1])/2;
      my $pq3=$alen-$pq1-1;
      $q3=($sort_set[$pq3+1]+$sort_set[$pq3])/2;
   }else{
      my $pq1=int($pos->[1]/2)+1;
      $q1=$sort_set[$pq1-1];
      my $pq3=$alen-$pq1;
      $q3=$sort_set[$pq3];
   }

}

printf("%.1f", $q3-$q1);

sub getMedianPos{
   my $len=shift;
   my $pos=[];
   if ($len%2==0){
      my $p1=$len/2-1;
      my $p2=$len/2;
      push (@$pos,$p1,$p2);
   }else{
      my $p= int($len/2);
      push (@$pos,$p);
   }
   return $pos;
}

