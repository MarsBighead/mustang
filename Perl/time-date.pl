#!/usr/bin/perl
use strict;
use warnings;
use Data::Dumper;

my $current_time = localtime();
print Dumper($current_time);
my ($sec,$min,$hour,$mday,$mon,$year,$wday,$yday,$isdst) = localtime();
my $month=$mon+1;
$month=sprintf("%02d", $month) if ($month<10);
$mday=sprintf("%02d", $mday) if ($mday<10);
$current_time = sprintf("%s%s%s", $year+1900, $month, $mday);
print $current_time." date\n";
$current_time=default_date();
print $current_time."\n";


if (0){
    print "true\n";
}else{
    print "false\n";
}

sub default_date {
    my ($sec,$min,$hour,$mday,$mon,$year,$wday,$yday,$isdst) = localtime();
    my $month=$mon+1;
    if ($mday>1){
        ($month, $mday) = ($month+1, 1);
    }
    return sprintf("%04d%02d%02d", $year+1900, $month, $mday);
}
