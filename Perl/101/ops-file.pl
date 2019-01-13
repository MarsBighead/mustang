#!/usr/bin/perl
use strict;
use warnings;

my $filename="mult_sequence.fasta";
open FH, "<$filename"  or die "Open $filename failled, $!";
my $NO=0;
while(<FH>){
	$NO++;
    my $line=$_;
	print "$NO\t$line";
}
close(FH);
