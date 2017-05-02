#!/usr/bin/perl

$Regex_Pattern = '^([\w\d\.\`\~\!\@\#\$\%\^\&\*\~\(\)\-\_\']{3}\.){3}[\w\d\.\`\~\!\@\#\$\%\^\&\*\~\(\)\-\_\']{3}$';

$Test_String = <STDIN> ;
 
$Test_String =~ /$Regex_Pattern/g ? print 'true' : print 'false';
