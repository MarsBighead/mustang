#!/usr/bin/perl
use warnings;
use strict;
use Storable;
use Data::Dumper;

my %table =("001","Wu","002","Lu");

store \%table, 'file';
my $hashref = retrieve('file');
print Dumper($hashref);
