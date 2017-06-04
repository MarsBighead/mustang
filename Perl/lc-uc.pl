#!/usr/bin/perl
use strict;

use warnings;

my $ss="Abbsbsbs";
my $uss=uc($ss);
print "uc(\$ss)=$uss\n";
my $lf=lcfirst($uss);
print "lcfirst(\$ss)=$lf\n";
my $lss=lc($ss);
print "lc(\$ss)=$lss\n";
my $uf=ucfirst($lss);
print "ucfirst(\$ss)=$uf\n";
