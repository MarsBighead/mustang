#!/usr/bin/perl 
use strict;
use warnings;

setTypeFinder("file1","file2",1);

sub setTypeFinder{
	my ($file1,$file2,$setTypeFinder)=@_;	
	open F1,"$file1";
	open F2,"$file2";	
	my %a=map{ $_=~s/[\s\r]+//g;$_ =>1 } <F1>;
	my %b=map{ $_=~s/[\s\r]+//g;$_ =>1 } <F2>;
	my @keys=keys %b;	
	print "key\t@keys\n";
	#intersection
	my @arrayRepeat=  grep( $a{$_}, keys %b );
	#union
	my @B1=grep(!defined $a{$_}, keys %b);
	my @B2=grep(!defined $b{$_}, keys %a);
	#complement
	my %C=map {$_=>1} keys %a,keys %b;
	#universal
	my %count=();
	my @un=grep{++$count{$_}<2} (@B1,@B2);	
	#chomp @arrayRepeat;
	print "Repeat one:\n@arrayRepeat\n";
	print "universal:\n@un\n";	
}
