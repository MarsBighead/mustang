#!/usr/bin/perl
use strict;
use warnings;

for my $v(0..10){
    if ($v%2==0) {
         print "if# Even number: $v\n";
    }
    unless ($v%2==1) {
         print "unless# Even number: $v\n";
    }

}
