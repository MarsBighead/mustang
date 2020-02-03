#!/usr/bin/perl

use strict;
use warnings;

#
# Complete the 'diagonalDifference' function below.
#
# The function is expected to return an INTEGER.
# The function accepts 2D_INTEGER_ARRAY arr as parameter.
#

sub diagonalDifference {
    # Write your code here
    my ($arr)=@_;
    my $num=$#{$arr};
    my ($ld,$rd)=(0,0);
    for my $n(0..$num){
        $ld+=$arr->[$n]->[$n];
        $rd+=$arr->[$n]->[$num-$n];
    }
    my $result=$rd-$ld;
    if ($ld > $rd){
        $result=$ld-$rd;
    }
    return $result;
}

open(my $fptr, '>', $ENV{'OUTPUT_PATH'});

my $n = ltrim(rtrim(my $n_temp = <STDIN>));

my @arr = ();

for (1..$n) {
    my $arr_item = rtrim(my $arr_item_temp = <STDIN>);

    my @arr_item = split /\s+/, $arr_item;

    push @arr, \@arr_item;
}

my $result = diagonalDifference \@arr;

print $fptr "$result\n";

close $fptr;

sub ltrim {
    my $str = shift;

    $str =~ s/^\s+//;

    return $str;
}

sub rtrim {
    my $str = shift;

    $str =~ s/\s+$//;

    return $str;
}
