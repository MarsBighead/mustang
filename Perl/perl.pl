#!/usr/bin/perl

use strict;
use warnings;

use Data::Dumper;
my $string="abcd,acd,tgc,sss,ug";
my @arr=split /\,/,$string,-1;
print "Arr\t@arr\n";

my @ref=\(1,2,3);
print "@ref\n";
print "OS $^O\n";
#my @ref2=@$ref;

my $hash={
   "ph"=>{
        "name"=>"penghua",
        "age"=> 18
        },
   "array"=>[1,2,3,4]
};

print Dumper($hash);
print "A".."Z";

print <<END
here document
This is test $string "<<"
END

