#!/usr/bin/perl
use warnings;
use strict;
use DateTime::Format::Pg;
use Data::Dumper;
use DateTime::Format::Strptime;
use DateTime::Duration;

my $dt = DateTime::Format::Pg->parse_datetime( '2017-07-16 23:12:01' );

my $dt1=DateTime::Format::Pg->parse_datetime( '2017-07-16 23:12:01' );
#print Dumper($dt);
#$dt->add( days => 1 )->subtract( seconds => 20 );
$dt->add( hours => 12 )->add( seconds => 20 );
#print Dumper($dt->formatter());
#       2003-01-16 23:12:0


=pod
my $strp = DateTime::Format::Strptime->new();
=cut
my $strp = DateTime::Format::Strptime->new(
        pattern   => '%T',
        locale    => 'en_AU',
        time_zone => 'Australia/Melbourne',
        );


print $strp->format_datetime($dt)."\n";

my $ft= DateTime::Format::Pg->format_datetime($dt);
print "$ft\n";
my $dur=$dt-$dt1;
print Dumper($dur);
print $dt->datetime."\n";
print $dt1->datetime."\n";
my %hash=$dur->deltas();
print Dumper(\%hash);
