#!/usr/bin/perl

$i=4;
$d=4.0;
$s='Hackerrank';

# Declare second integer, double, and String variables.
$add_i=<STDIN>;
$add_d=<STDIN>;
$add_s=<STDIN>;

printf ("%d\n",($i+$add_i));
printf("%.1f\n", ($d+$add_d));
printf("%s", $s." $add_s\n");

# # Read and save an integer, double, and String to your variables.
#
# # Print the sum of both integer variables on a new line.
#
# # Print the sum of the double variables on a new line.
#
# # Concatenate and print the String variables on a new line
# # The 's' variable above should be printed first.
