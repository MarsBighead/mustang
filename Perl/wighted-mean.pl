#!/usr/bin/perl
# In the previous challenge, we calculated a mean.
# In this challenge, we practice calculating a weighted mean.
# Check out the Tutorial tab for learning materials and an instructional video!


my $n=<STDIN>;
chomp $n;
my $x=<STDIN>;
chomp $x;
my $w=<STDIN>;
chomp $w;
my @arrx=split / /,$x;
my @arrw=split / /,$w;
my ($t, $b)=(0,0);

for (my $i=0;$i<$n;$i++){
   $t+=$arrx[$i]*$arrw[$i];
   $b+=$arrw[$i];
}

printf("%.1f",$t/$b);
