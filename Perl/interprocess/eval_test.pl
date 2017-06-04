#!/usr/bin/perl
use strict;
use warnings;

#########################################################################
# If inovking the eval code in the main code with strict,warnings mode, #
# the eval code should follow the same rules                            #
#########################################################################

my  @program=();
push ( @program,'my $i = 1;');
push ( @program,'my $i = 3; my $j = 2; my $k = $i + $j');
push ( @program, 'my $i = 3; return 24; my $k = $i'); 
 
foreach my $exp (@program)
{
    my $rtn= eval($exp);
    print "$rtn\n";
}
