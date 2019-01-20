#!/usr/bin/perl
use strict;
use warnings;
use Data::Dumper;
use Tie::CharArray;
my $foobar = 'a string';
tie my @foo, 'Tie::CharArray', $foobar;
print Dumper(\@foo);
$foo[0] = 'A';    # $foobar = 'A string'
push @foo, '!';   # $foobar = 'A string!'
print Dumper(\@foo);
