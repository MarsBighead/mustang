#!/usr/bin/perl 
use strict;
use warnings;


findFileType("/home/hbu/Desktop/",".php");
sub findFileType{
	my $dir=shift;
	my $fileType=shift;
	#print "$dir\n";
	opendir (DIR,$dir);
	my @arrayFile=grep {!(/^\.$/ or /^\.\.$/)} readdir(DIR);
	#print "@arrayFile\n";
	my @array=();
	foreach(@arrayFile){		
		my $dirChildName="$dir"."$_";
		if(-d $dirChildName){
			if(-s $dirChildName){   
				push (@array,findFileType("$dirChildName/","$fileType"));
			}else{ next;}
		}
		if(-f $dirChildName){
			$fileType=~s/^\.//;
			if($dirChildName=~/\.($fileType)$/){
				print "Targetype\t$fileType|$1\t$dirChildName\n";
				push (@array,$dirChildName);
			}
		}        
        #print "\n";
	}
	@arrayFile=();
	return @array;
}
