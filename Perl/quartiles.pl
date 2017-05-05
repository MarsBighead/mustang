#!/usr/bin/perl


my @arr = split / /,"3 7 8 5 12 14 21 15 18 14";
#my @s = sort{$a <=> $b} @arr;
for (my $i=0;$i<scalar(@arr);$i++){
   for (my $j=$i+1;$j<scalar(@arr);$j++){
      if ($arr[$j]<$arr[$i]){
         my $tmp = $arr[$i];
         $arr[$i] = $arr[$j];
         $arr[$j] = $tmp;
      }

   }
   print "@arr\n";
}
my @s=@arr;
print "@s\n";
$Q2=median_Number(@s);
if ( scalar(@s)%2==0 ){
    $h=scalar(@s)/2;
    $Q1=median_Number(@s[0..$h-1]);
    $Q3=median_Number(@s[$h..scalar(@s)-1]);
}else {
    $h=int(scalar(@s)/2);
    $end_pos1=$h-1;
    print "|$h \n";
    $Q1=median_Number(@s[0..$end_pos1]);
    print "|$h \n";
    $start_pos2 =$h+1;
    $Q3=median_Number(@s[$start_pos2..scalar(@s)-1]);
    print @s[$start_pos2..scalar(@s)-1];
}
print "$Q1\n";
print "$Q2\n";
print "$Q3\n";


sub median_Number{
    my @s= @_;
    if (scalar(@s)%2==0){
       my $h  = scalar(@s)/2;
       return ($s[$h]+$s[$h-1])/2;
    } else {
       my $h  = int(scalar(@s)/2);
       return  $s[$h];
    }
}
