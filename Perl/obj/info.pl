#!/usr/bin/perl

BEGIN {
    push (@INC,'.');
}
use strict;
use warnings;

use Reporter;

use Data::Dumper;
my $rhash={
    "header" => [1,2,3,4],
    "data" => ['1a','2b','3c','4d'],
};
print Dumper($rhash);
my  $object = new Reporter( $rhash, "convert", 'load');
# Get first name which is set using constructor.
my $firstName = $object->getOp();

print "Before Setting First Name is : $firstName\n";

# Now Set first name using helper function.
$object->setOp( "Mohd." );

# Now get first name set by helper function.
$firstName = $object->getOp();
print "Before Setting First Name is : $firstName\n";

