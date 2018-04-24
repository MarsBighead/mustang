#!/usr/bin/perl

use strict;
use warnings;

use Data::Dumper;
use VMware::UM::Report;

use Storable qw(dclone);

my $xx;
if (!defined($xx)){
   print "\$xx is undefined!\n";
}

my $mem=um_format2mb("20 GB");
print "Memory size is  $mem\n";
my @list = (
       sub{ map {"num:$_"} (1 .. 10)}->(),
       sub{ map {"str:$_"} ('a','b','c','d','e')}->(),
); 
print "@list\n";
print ref(\@list)."\n";
print ref(\$mem)."\n";
my $a=[(1..10)];
my $b=[@$a];

$b->[0]=12;
#print Dumper($b);
#print Dumper($a);
my $relation_attribute={
    type => "varchar(__LEN__)",
	len  => 2048,
	pos  => [1,2],
};
my $x =[];
map { push (@$x, dclone($relation_attribute)); } (1..3);
print Dumper($x);
#my @array=map { dclone($relation_attribute); } (1..3);
#print Dumper(\@array);

my $rattr=dclone($relation_attribute);
$rattr->{pos} = [4,5];

my $array=[$relation_attribute, $rattr];
shift (@{$array->[0]->{pos}});
shift (@{$array->[0]->{pos}});
if (scalar(@{$array->[0]->{pos}})==0) {
   shift(@$array);
}
print Dumper($array);
my @array = ('first', 'second', 'third', 'fourth');
splice (@array, 0, 2);
print "@array\n";
# 删除数组第二个元素
@array = ('first', 'second', 'third', 'fourth');
splice (@array, 1, 1);
print "@array\n";  
@array        = ("a", "e", "i", "o", "u");
my @removedItems = splice(@array, 0 , 2, ("A", "E", "I"));

print "Removed items: @removedItems\n";
print "Saved items: @array\n";
