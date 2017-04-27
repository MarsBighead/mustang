#!/usr/bin/perl

$mc  = <STDIN>;
$tip = <STDIN>;
$tax = <STDIN>;

$totalCost = $mc +$mc*($tip/100)+$mc*($tax/100);

printf("The total meal cost is %d dollars.\n",int($totalCost+0.5));
