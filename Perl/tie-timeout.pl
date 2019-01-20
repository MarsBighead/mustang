#!/usr/bin/perl
use Tie::Scalar::Timeout;
tie my $k, 'Tie::Scalar::Timeout', EXPIRES => '+2s';
$k = 123;
print "\$k = $k\n";
sleep(3);
# $k is now undef
print "\$k = $k\n";
