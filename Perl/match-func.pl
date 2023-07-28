#!/usr/bin/perl
use strict;
#use warnings;

my $re_pattern = "^func(?:\s\(\w+\s\*\w+\))?\s(\w+)\(";
#my $re_pattern = '^func\s\(\w+\s\*\w+\)\s(\w+)\)';
#$Test_String = <STDIN> ;
my $test_string= {
    "func"  => "func loadData(db *sql.DB, filename string) error {",
    "method" =>"func (df *DataFilename) Load(dns string) error {",
};
foreach my $key(keys %{$test_string}) {
    my $Test_String=$test_string->{$key};
    print "test: $Test_String\n";
    if($Test_String =~ /$re_pattern/){
        print "$1,true\n";
    } else {
        print "false\n";
    }
}
