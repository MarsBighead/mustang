#!/usr/bin/perl -w
# Filename : client.pl
 
use strict;
use Socket;
 
# initialize host address and port
my $host = shift || 'localhost';
my $port = shift || 7890;
my $server = "localhost";  # host address
 
# create socket and connect to it
socket(SOCKET,PF_INET,SOCK_STREAM,(getprotobyname('tcp'))[2])
   or die "Unable create socket $!\n";
connect( SOCKET, pack_sockaddr_in($port, inet_aton($server)))
   or die "Unable to connect port: $port! \n";
 
my $line;
while ($line = <SOCKET>) {
        print "$line\n";
}
close SOCKET or die "close: $!";
