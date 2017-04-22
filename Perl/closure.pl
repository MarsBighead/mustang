#!/usr/bin/perl

use strict;
use warnings;

my $string="abcd,acd,tgc,sss,ug";

print "O $string\n";

{

my $string="abcd,acd,tgc";
print "B $string\n";

}
print "O $string\n";
