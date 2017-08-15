#!/usr/bin/perl
use strict;
use warnings;
use Data::Dumper;

use Text::CSV;

 my @rows;
 my $csv = Text::CSV->new ( { binary => 1 } )  # should set binary attribute.
                 or die "Cannot use CSV: ".Text::CSV->error_diag ();
 
open my $fh, "<:encoding(utf8)", "new.csv" or die "test.csv: $!";
while ( my $row = $csv->getline( $fh ) ) {
    #$row->[2] =~ m/pattern/ or next; # 3rd field should match
    push @rows, $row;
}
$csv->eof or $csv->error_diag();
close $fh;

print Dumper(\@rows);
$csv->eol ("\n");

open $fh, ">:encoding(utf8)", "new1.csv" or die "new.csv: $!";
$csv->print ($fh, $_) for @rows;
close $fh or die "new.csv: $!";

exit;

# parse and combine style


#$status = $csv->combine(@columns);    # combine columns into a string
#$line   = $csv->string();             # get the combined string

#$status  = $csv->parse($line);        # parse a CSV string into fields
#@columns = $csv->fields();            # get the parsed fields

#$status       = $csv->status ();      # get the most recent status
#$bad_argument = $csv->error_input (); # get the most recent bad argument
#$diag         = $csv->error_diag ();  # if an error occured, explains WHY

#$status = $csv->print ($io, $colref); # Write an array of fields
                                      # immediately to a file $io
#$colref = $csv->getline ($io);        # Read a line from file $io,
                                      # parse it and return an array
#   $csv->column_names (@names);          # Set column names for getline_hr ()
#      $ref = $csv->getline_hr ($io);        # getline (), but returns a hashref
#      $eof = $csv->eof ();                  # Indicate if last parse or
                                           
# getline () hit End Of File
#       $csv->types(\@t_array);               # Set column types
