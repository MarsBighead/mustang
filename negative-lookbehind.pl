#!/usr/bin/perl

$Regex_Pattern = '([\w\W])(?!\1)|([\w\W]$)';
#Forbidden capture
$Regex_Pattern ='(?:([\w\W])(?!\1))|(?:[\w\W]$)';
#$Test_String = <STDIN> ;
$Test_String = "###$$$$"; 
$count = () = $Test_String =~ /$Regex_Pattern/g;
print "Number of matches : ",$count ;
