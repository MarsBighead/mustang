#!/usr/bin/perl
use Data::Dumper;

my ($start,$end)=(1, 15);

for ($start..$end){
    print "$_\n";
}
my @a=(split //, 0 x $end);
print "\n@a\n";
