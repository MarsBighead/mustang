#!/usr/bin/perl
use strict;
use warnings;
use Data::Dumper;
use Mars::Mustang;

my $status_code=mg_toger();
print "$status_code\n";

my $a = [1,2,3,4,5];
print "type \$a ".ref($a)."\n";
my $h = {
    "paul"=>"duan",
    "num" => $a,
};

print Dumper($h);
print "type \$h ".ref($h)."\n";
print "type \$h->{num} ".ref($h->{"num"})."\n";
