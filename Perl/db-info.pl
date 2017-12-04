#!/usr/bin/perl
use strict;
use warnings;
use DBI;

my @drivers = DBI->available_drivers() ;
die "No dirvers defined! \n" unless @drivers; 
foreach my $driver (@drivers) { 
    print "Driver: $driver \n"; 
    my @data_sources = DBI->data_sources ($driver); 
    foreach my $data_source (@data_sources) { 
        print "\t Data Source: $data_source \n"; 
    }
    print "\n"; 
}
