#!/usr/bin/perl

use warnings;
use strict;
use threads;

print localtime(time)."\n";
my $j = 0;
my $thread;
while(){
    last if ($j>10);
    while(scalar(threads->list())<5){
        $j++;
        threads->new(\&ss,$j, $j);
    }
    foreach $thread(threads->list(threads::all)) {
        if ($thread->is_joinable()){
            $thread->join(); 
            print scalar(threads->list())."\tthreads num:$j\t".localtime(time)."\n"; 
        } 
    
    }
}

foreach $thread(threads->list(threads::all)) {
    $thread->join();
    print scalar(threads->list())."\t$j\t".localtime(time)."\n";
}

sub ss(){
    my ($t, $s) = @_;
    sleep($t);
    print "sub ss:$s, sleep $t secs\t";
}
