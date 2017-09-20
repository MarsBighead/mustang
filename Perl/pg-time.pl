#!/usr/bin/perl
use warnings;
use strict;
use DateTime::Format::Pg;
use Data::Dumper;
my $dt = DateTime::Format::Pg->parse_datetime( '2017-07-16 23:12:01' );

print Dumper($dt);
#       2003-01-16 23:12:0
#       DateTime::Format::Pg->format_datetime($dt);
