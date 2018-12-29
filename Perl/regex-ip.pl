#!/usr/bin/perl

use strict;
use warnings;

my $ip="0.0.0.0";
if (scalar(@ARGV)>0){
   $ip=shift @ARGV;
}
print "ip $ip\n";
#my $regex="((?:(?:25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))\.){3}(?:25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d))))";
if ($ip=~/^(?:(?:0\.){3}0)|((?:25[0-5]|2[0-4]\d|((1\d{2})|[1-9]))((?:\.)(?:25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))){3})$/){
   print "$ip passed!\n";
} else{
   print "$ip not passed!\n";
}
 my $x="0";
if ($x=~/^([1-9])$/){
  print "0-9 $x\n";
}

my $api="http://reviewboard.eng.vmware.com/api/review-requests/1456719/";

if ($api=~/^http(?:s)?:\/\/(?:(?:\w+\.)+)\w+\/api\/review-requests\/([1-9]\d+)\//){
   print "API https or http: $api\n$1\n";
} else{
   print "Unmatch\n";
}

