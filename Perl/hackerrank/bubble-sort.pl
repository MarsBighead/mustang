#!/usr/bin/perl

use strict;
use warnings;

# Complete the countSwaps function below.
sub countSwaps {
    my ($a)=@_;
    my $cnt=0;
    my $num=$#{$a};
    for my $i(0..$num) {
        for my $j(0..$num-1) {
            # Swap adjacent elements if they are in decreasing order
            if ($a->[$j] > $a->[$j + 1]) {
                ($a->[$j], $a->[$j + 1])=($a->[$j + 1], $a->[$j]);
                $cnt++;
            }
        }
    }
    print "Array is sorted in $cnt swaps.\n";
    print "First Element: ".$a->[0]."\n";
    print "Last Element: ".$a->[$num]."\n";
}

my $n = <>;
$n =~ s/\s+$//;

my $a = <>;
$a =~ s/\s+$//;
my @a = split /\s+/, $a;

countSwaps \@a;
