#!/usr/bin/perl 

package Reporter;
use Data::Dumper;
use Text::CSV;

sub new
{
    my $class = shift;
    my $self = {};
    # Print all the values just for clarification.
    bless $self, $class;
    return $self;
}
sub export {
    my ($self, $output, $type) = @_;
    my $csv = Text::CSV->new({binary => 1});
    $csv->eol ("\r\n");
    print Dumper($self->{_data});
    open my $fh, ">:encoding(utf8)", $output or die "$output: $!";
    $csv->print ($fh,  $self->{_header});
    $csv->print ($fh, $_) for (@{$self->{_data}});
    close $fh;
    return $output;
}
sub set {
    my ($self, $rhash) = @_;
    $self->{_data} = $rhash->{_data} if defined($rhash);
    $self->{_header} = $rhash->{_header} if defined($rhash);
}

sub load {
    my($self,$tablename,$pghash) = @_;
}
sub export {
    
}
1;
