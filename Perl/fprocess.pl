#!/usr/bin/perl

my $pid = fork();
if (not defined $pid) {
    print "resources not avilable.\n";
} elsif ($pid == 0) {
    print "IM THE CHILD \n";
    sleep 5;
    print "IM THE CHILD2 \n";
    exit(0);
} else {
    print "IM THE PARENT \n";
    waitpid($pid,0);
}
print "HIYA "; 
