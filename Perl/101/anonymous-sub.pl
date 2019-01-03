#!/usr/bin/perl

my $subref = sub { 
    my $val=shift;
    if ($val ne ""){
        print "First parameter with value: $val. ";
    }
    print "anonymous subroutine\n"; 
};

&$subref;
&$subref("a parameter");
$subref->();
$subref->("a parameter");


