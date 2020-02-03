#!/usr/bin/perl

use strict;
use warnings;

# Complete the repeatedString function below.
sub repeatedString {
    my ($s, $n)=@_;
    my @arr=split //, $s;
    my $hash={};
    my $single_count=0;
    foreach my $e(@arr){
        if ($e eq "a"){
            $single_count++;
        }
    }
    return 0 if ($single_count==0);
    my $repeated_num=int($n/length($s));
    my $cnt=$repeated_num*$single_count;
    my $remain_num=$n%length($s);
    foreach my $e(@arr[0..$remain_num-1]){
        if ($e eq "a"){
            $cnt++;
        }
    }
    return $cnt;

}


my $s="ceebbcb";
my $n=817723;

my $result = repeatedString $s, $n;
print "$result\n";
