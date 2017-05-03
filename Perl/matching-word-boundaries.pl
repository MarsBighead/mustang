#!/usr/bin/perl 
# Matching Word Boundaries
# https://www.hackerrank.com/challenges/matching-word-boundaries

$Regex_Pattern = '\w*((\b|)[aeiouAEIOU][A-Za-z]*\b)+\w*';
#$Test_String = <STDIN> ;
$Test_String="found3 isdvnslknc98098sdcsdbc";
if($Test_String =~ /$Regex_Pattern/){
    print "true";
} else {
    print "false";
}
