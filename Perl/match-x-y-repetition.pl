#!/usr/bin/perl 
# w{3,5} : It will match the character w ,  3, 4 or 5 times. 
# [xyz]{5,} : It will match the character x, y or z 5 or more times. 
# \d{1, 4} : It will match any digits 1, 2, 3, or 4 times.

# https://www.hackerrank.com/challenges/matching-x-y-repetitions
$Regex_Pattern = '^\d{1,2}[a-zA-Z]{3,}\.{0,3}$';

#$Test_String = <STDIN> ;
# expect true
$Test_String="3threeormorealphabets.";
if($Test_String =~ /$Regex_Pattern/){
    print "true";
} else {
    print "false";
}
