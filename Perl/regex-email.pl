#!/usr/bin/perl

use strict;
use warnings;


my $regex= '^(\\s*)(([^+=\/<>\\\\(){}%\‘|!#&*?~\[\\]\\,;:\\s@“]+(\\.[^<>()\\[\\]\\\\.,;:\\s@“]+)*)|(".+"))@((\\[[0-9]{1,3}\\.[0-9]{1,3}.[0-9]{1,3}\\.[0-9]{1,3}])|(([a-zA-Z\\-0-9]+\\.)+[a-zA-Z]{2,}))(\\s*)$';
my $emails=['ab@test.com', ' ab@test.com', ' ab@test.com ', '  ab@test.com', 'ab@test.com ' ];
for my $email(@$emails){
    if ( $email=~/$regex/){
        print "Matched\n";
    } else{
        print "|$email| Not matched\n";
    }
}

my $pattern='^\s*((http(s)?://)|)(([a-zA-Z0-9\._-]+\.[a-zA-Z]{2,6})|([0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}))(:[0-9]{1,4})*(/[a-zA-Z0-9\&%_\./-~-]*)?(\s*)$';
my $urls=[
    "www.localhost.com",
    " http://localhost.com",
    "http://localhost.com",
    "  http://localhost.com ",
    "  http://localhost.com    ",
];

print "URL:\n";
for my $url(@$urls){
    if ( $url=~/$pattern/){
        print "Matched\n";
    } else{
        print "|$url| Not matched\n";
    }
}
