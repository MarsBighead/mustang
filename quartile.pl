#!/usr/bin/perl
use Data::Dumper;

my $len=<STDIN>;

chomp $len;

my $pos=getQuartlePos($len);
print Dumper($pos);

sub getQuartlePos{
    my $len=shift;
    my $hash={
         "q1"=>[],
         "q3"=>[],
    };
    # print Dumper($hash);
    if ($len%4==1 or $len%4==0){
        my $qe1=int($len/4);
        my $qs1=$qe1-1;
        push (@{$hash->{"q1"}}, $qs1,$qe1);
        my $qs3=$len-1-$qe1;
        my $qe3=$qs3+1;
        push (@{$hash->{"q3"}}, $qs3,$qe3);
    } elsif($len%4==2 or $len%4==3){
        my $q1=int($len/4);
        push (@{$hash->{"q1"}}, $q1);
        my $q3=$len-1-$q1;
        push (@{$hash->{"q3"}}, $q3);
       
    }
    return $hash;
}

