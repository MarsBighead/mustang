#!/usr/bin/perl

use strict;
use warnings;
use Data::Dumper;

my @xyz=();

while (<DATA>){
    next if (/^\s*#.*$/);
    push @xyz,[split];
}

my $i=0;
foreach my $pt(@xyz){
    print  "point",$i++,": x = $pt->[0], y = $pt->[1], z = $pt->[2]\n";
}

# Get X axis
my (@x)=();
for my $pt(@xyz){
    push @x,$pt->[0];
}
print Dumper(\@x);
# Get X axis with map
@x = map{$_->[0]} @xyz;
print Dumper(\@x);

__DATA__
#points coordinate data
1 2 3
4 5 6
7 8 9
