#!/usr/bin/perl

open F,"tc05.txt";
while(<F>){
    my $line = $_;
    chomp($line);
    # my $Regex_Pattern = '(?:href\=\"([\w:;\.\/&=\?]+)\"[\w\W]*\>([\w\W]*)\<\/a\>)+';
    my $Regex_Pattern ='<a\shref=\"([^"]*)\"\s?[^>]*>(?:<b>)?([\w|\s|/|.|,]*)(?:<\/b>)?';
    while ($line =~m/$Regex_Pattern/g){
        my ($url, $content) = ($1, $2);
        $content =~ s/^\s+//g;
        print "$url,$content\n";
    }
}
exit;
my $n=<STDIN>;
chomp $n;
for (my $i=0;$i<$n;$i++){
   my $input = <STDIN>;
   chomp($input);
   my $Regex_Pattern = 'href\=\"([\w:;\.\/&=\?]+)\"[\w\W]*\>([\w\W]*)\<\/a\>';
   if ($input =~ /$Regex_Pattern/){
      print "$1,$2\n";
   }
}

