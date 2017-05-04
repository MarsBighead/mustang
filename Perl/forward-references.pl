

$Regex_Pattern = '^(?:tac){2,}(?:tictac)*(?:tac)*(?|(tic)|(tac))$';

$Test_String = <STDIN> ;
if($Test_String =~ /$Regex_Pattern/){
    print "true";
} else {
    print "false";
}
