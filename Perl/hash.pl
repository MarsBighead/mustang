#!/usr/bin/perl

use strict;
use warnings;

my $n = <STDIN>;
chomp $n;
my $hash={};
for (my $i=0;$i<$n;$i++){

   my $s = <STDIN>;
   chomp ($s);
   my ($key,$value) = split / /,$s;
   $hash->{$key}=$value;
}
for (my $i=0;$i<$n;$i++){
   my $key =<STDIN>;
   if ($key eq ""){
       next;
   }
   chomp($key);
   if (exists($hash->{$key})){
      print "$key=".$hash->{$key}."\n";
   }else{
      print "Not found\n";
   }
}
