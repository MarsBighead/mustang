#!/usr/bin/perl
use strict;
use warnings;
use POSIX ":sys_wait_h";

our $zombies = 0; 
our $kid_proc_num = 0;
$SIG{CHLD} = sub { $zombies++ }; 
my $maxchild=10;
foreach my $item (1..5000) {
    for my $i(1..$maxchild) { 
        my $pid = fork(); 
        if( !defined($pid) ) { exit 1; } 
            unless($pid) { 
            print "Child poresss, $item, pid=$$.\n";
            print rand($item)."\n";
            sleep(2);
            exit 0; 
        } 
        $kid_proc_num++; 
    }
    while (1) { 
        if( $zombies > 0 ) { 
            $zombies = 0; 
            my $collect; 
            while(($collect = waitpid(-1, WNOHANG)) > 0) { 
                $kid_proc_num--; 
            }
        }
        if($kid_proc_num==0) { last; } 
        else { next; } 
   }
}