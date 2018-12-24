#!/usr/bin/perl

use strict;
use warnings;

use Data::Dumper;
my $arr=[(1..10)];
print Dumper($arr);
print  ref($arr)."\n";
my $string="abcd,acd,tgc,sss,ug";
my @arr=split /\,/,$string,-1;
print  ref(\@arr)."| ".length($arr[0])." plain\n";
my $string="abcd,acd,tgc,sss,ug";
print length(@arr)." length\n";
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

print length(%$hash)." length of hash\n";
print Dumper($hash);
print "A".."Z";

print <<END
here document
This is test $string "<<"
END

