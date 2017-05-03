#!/usr/bin/perl 



# expect true
$Test_String="think?";
# $Test_String = <STDIN> ;
if($Test_String =~ /$Regex_Pattern/){
    print "true";
} else {
    print "false";
}
