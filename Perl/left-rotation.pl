#!/usr/bin/perl
use strict;
use warnings;
my $nd=<STDIN>;
chomp $nd;
my ($n,$d)=split / /,$nd;
my $s=<STDIN>;
chomp $s;
my @a=split / /,$s;
for (my $i=0;$i<$d;$i++){
    my $tmp = shift @a;
    push (@a,$tmp);
}
print join(" ",@a) ;
