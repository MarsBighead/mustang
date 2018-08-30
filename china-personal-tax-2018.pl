#!/usr/bin/perl
use strict;
use warnings;
use Data::Dumper;

if(scalar(@ARGV)<1){
    print "Please input your income\n";
    exit(1);
}
my $income=$ARGV[0];

my $hash={
    36000=>0.03,
    144000=>0.10,
    300000=>0.20,
    420000=>0.25,
    660000=>0.30,
    960000=>0.35,
    -1=>0.45,
};
my $mhash={};
my $rhash={};
my ($base, $tax)=(0,0);
foreach my $key(sort{$a<=>$b} keys %{$hash}){
   if ($key=~m/^\d+$/){
       my $rate=$hash->{$key};
       $key=$key/12;
       $tax=$tax+($key-$base)*$rate;
       $mhash->{$key}=$tax;
       $base=$key;
       $rhash->{$key}=$rate;
   }else{
       $rhash->{$key}=$hash->{$key};
   }
}

my $qcd=5000;
$income-=$qcd;
my ($m, $n)=(0,0);
if ($income>80000){
    $n=($income-80000)*$rhash->{-1}+$rhash->{80000}; 
    $m=$income-$n;
}elsif($income<=80000 && $income>55000){
    $n=($income-80000)*$rhash->{80000}+$mhash->{80000}; 
    $m=$income-$n;
}elsif($income<=55000 && $income>35000){
    $n=($income-55000)*$rhash->{55000}+$mhash->{55000}; 
    $m=$income-$n;
}elsif($income<=35000 && $income>12000){
    $n=($income-35000)*$rhash->{35000}+$mhash->{35000}; 
    $m=$income-$n;
}elsif($income<=12000 && $income>3000){
    $n=($income-12000)*$rhash->{12000}+$mhash->{12000}; 
    $m=$income-$n;
}elsif($income<=3000){
    $n=($income-3000)*$rhash->{3000}+$mhash->{3000}; 
    $m=$income-$n;
}
$m+=$qcd;
print "Income: $income\nTax: $n\nAfter-tax income: $m\n";
