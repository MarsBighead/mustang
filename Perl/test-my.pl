#!/usr/bin/perl

#keywords my
my $Scalar =1;                #Private variable, action scope is the package/file
sub Subroutine{  
    my $Scalar =2;            #Private variable, action scope is in the subroutine 
    $Scalar +=1; 
    print $Scalar,"\n";
} 
&Subroutine;                     #输出3  
&Subroutine;                     #输出3   
print $Scalar."\n";              #输出1

my $step=9;
my $j=2;
my $sql="";
calculate_test($sql, $step*$j);
sub calculate_test{
    my ($template , $num)=@_;
    print "\$num is $num\n";
}
#https://www.cs.ait.ac.th/~on/O/oreilly/perl/advprog/ch03_02.htm
# Unable to apply with variable which has been declared by my
#Typeglobs, we mentioned earlier, can be localized (with local only) and assigned to one another. Assigning a typeglob has the effect of aliasing one identifier name to another. 
$array = 1;
@array = (1..10);
$spud   = "Wow!";
@spud   = ("idaho", "russet");
*potato = *spud;   # Alias potato to spud using typeglob assignment
print "$potato\n"; # prints "Wow!"
print @potato, "\n"; # prints "idaho russet"

*protato = *array;
print "@protato\n";
