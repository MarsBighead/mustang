#!/usr/bin/perl

use strict;
use warnings;

my $s="ATATCGTTTTATCGTT";
# $s="AATTTCCGACTTGCATGACGAGTCAGCGTTCCATCTGATCGAGTCTCCGAAGAACAAATACCCCTACTCAGTTGTGAGCCCCTTTACCGTGAGGACAGGGTCCTTGATGTCGTCTCCTAATTTGCGTTGCGGCTCAACATGTTGTACATAGTGGGGCCAGCCCCAGGGATTTTGTAATTTCTACACTCCATATACGGGACAAGGGTGAGCATTTCCGGGCTTGGATAGGGGCTGCAAGAAAATATCTGGACGTAAGAACTTAATGCCATTCCTACATCCTCGATACCTCGTCTGTCAGAGCAATGAGCTGGTTAGAGGACAGTATTGGTCGGTCATCCTCAGATTGGGGACACATCCGTCTCTATGTGCGTTCCGTTGCCTTGTGCTGACCTTGTCGAACGTACCCCATCTTCGAGCCGCACGCTCGACCAGCTAGGTCCCAGCAGTGGCCTGATAGAAAAATTACCTACGGGCCTCCCAATCGTCCTCCCAGGGTGTCGAACTCTCAAAATTCCCGCATGGTCGTGCTTCCGTACGAATTATGCAAACTCCAGAACCCGGATCTATTCCACGCTCAACGAGTCCTTCACGCTTGGTAGAATTTCATGCTCGTCTTTTGTATCCGTGTAAGTAGGAGGCCGCTGTACGGGTATCCCAGCCTTCGCGCTCTGCTGCAGGGACGTTAACACTCCGAACTTTCCATATACGGGACAAGGGTGAGCATTTCCGGGCTTGGATAGGGGCTGCAAGAAAATATCTGGACGTAAGAAGCTCTGAGGGATCCTCACGGAGTTAGATTTATTTTCCATATACGGGACAAGGGTGAGCATTTCCGGGCTTGGATAGGGGCTGCAAGAAAATATCTGGACGTAAGAAGAGTGATGTTTGGAATGCCAACTTCCATGCACGCCAATTGAGCAATCAGGAGAATCGAGTGCTGTTGACCTAGACCTTGTCAGAAGTATGAATTAACCGCGCGTGTAGGTTTGTCGCTCGACCTGCAAGGGTGCACAATCTGGACTGTCGTCGGCGAACGCTTTCATACGCCTACAAACCGCGTTGCTGGTCGAATCGATCTCACCACCGGCCTTGCAGGATTCTAATTATTCTCTCTCGGTGAGACTGCCGGCGGTCCATGGGTCTGTGTTTCGCTTCAAGCAGTGATATACTGGCGTTTTGTGACACATGGCCACGCACGCCTCTCGTTACTCCCAAT";
my @a=split "",$s;
#my $matrix=[];
my ($matrix ,$lat,$lon,$value)=([],0,0,0);
print "Length of the string ".length($s)."\n";
for (my $i=0;$i<@a;$i++){
    my $c1 = $a[$i];
    for (my $j=0;$j<=$i;$j++){
        my $c2 = $a[$j];
        if ($i>0 && $j>0){
           my $p = $matrix->[$i-1]->[$j-1];
           $matrix->[$i]->[$j] = center_alignment($a[$i],$a[$j],$p);
        }else {
           $matrix->[$i]->[$j]=edge_alignment($a[$i],$a[$j]);
        }

        if ($j<$i){
            if ($matrix->[$i]->[$j]>$value){
                $value = $matrix->[$i]->[$j];
                $lat = $i;
                $lon = $j;
            }
        }
    }
}


print "  ".join (" ",@a)."\n";
for (my $i=0;$i<scalar(@$matrix);$i++){
    print "$a[$i] ";
    my $row=$matrix->[$i];
    for my $val(@$row){
        print "$val ";
    }
    print "\n";
}
print "\n\n#######\n#######\n";
printf("End Position(%d,%d).\nScore value(length) %s.\nOffset of substr %d\n", $lat, $lon, $value, $value-1);

my $offset = $value -1;
my $s1=substr($s,$lat-$offset,$value);
my $s2=substr($s,$lon-$offset,$value);
print "s1#$s1\n";
print "s2#$s2\n";



sub edge_alignment{
    my ($c1, $c2)=@_;
    if ($c1 eq $c2){
        return 1;
    }else{
        return 0;
    }
}

sub center_alignment{
    my ($c1, $c2, $p)=@_;
    if ($c1 eq $c2){
        $p++;
    }else{
        $p=0;
    }
    return $p;
}
