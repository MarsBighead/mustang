#!/usr/bin/perl 

package Reporter;
use Data::Dumper;

sub new
{
    my $class = shift;
    my $self = {
        _data => shift,
        _type  => shift,
        _op       => shift,
    };
    # Print all the values just for clarification.
    print Dumper($self->{_data}); 
    print "Last Name is $self->{_type}\n";
    print "SSN is $self->{_op}\n";
    bless $self, $class;
    return $self;
}
sub setOp {
    my ( $self, $rhash ) = @_;
    print Dumper($rhash);
    $self->{_data} = $rhash if defined($rhash);
    return $self->{_op};
}

sub getOp {
    my( $self ) = @_;
    return $self->{_op};
}
1;
