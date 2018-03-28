#!/usr/bin/perl

BEGIN {
    push (@INC,'.');
}
use strict;
use warnings;

use Reporter;

use Data::Dumper;
my $rhash={
    "_data" => [
                [1,2,3,4],
                [1,2.2,3,4]
                ],
    "_header" => ['ID','Name','MemorySizeMB','Reservation'],
};
my  $object = new Reporter(); 
$object->set($rhash);
$object->export("converted.csv",".csv");
