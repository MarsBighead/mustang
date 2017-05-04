#!/usr/bin/perl 
# w{3,5} : It will match the character w ,  3, 4 or 5 times. 
# [xyz]{5,} : It will match the character x, y or z 5 or more times. 
# \d{1, 4} : It will match any digits 1, 2, 3, or 4 times.

#Backreferences To Failed Groups
#https://www.hackerrank.com/challenges/backreferences-to-failed-groups/copy-from/43616149
$Regex_Pattern = '^[0-9]{2}(\-?)([0-9]{2}\1){2}[0-9]{2}';

#$Test_String = <STDIN> ;
# expect true
$Test_String="12-34-56-87";
if($Test_String =~ /$Regex_Pattern/){
    print "$1|$2\n";
    print "true";
} else {
    print "false";
}
