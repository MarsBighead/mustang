#!/usr/bin/perl -w
# Filename : server.pl
 
use strict;
use warnings;
use Socket;        # Define PF_INET and SOCK_STREAM
 
# monitor port
my $port = shift || 7890;
my $proto = getprotobyname('tcp');
my $server = "localhost";  # set local host address 
 
# create socket, multiplexed port， and create multiple connection
socket(SOCKET, PF_INET, SOCK_STREAM, $proto)
   or die "Unable to open socket $!\n";
setsockopt(SOCKET, SOL_SOCKET, SO_REUSEADDR, 1)
   or die "Unable to set SO_REUSEADDR $!\n";
 
# Bind the port and monitor
bind( SOCKET, pack_sockaddr_in($port, inet_aton($server)))
   or die "unable to bind the port: $port! \n";
 
listen(SOCKET, 5) or die "listen: $!";
print "Launch vist port：$port\n";
 
# recieve the request
my $client_addr;
while ($client_addr = accept(NEW_SOCKET, SOCKET)) {
   # send them a message, close connection
   my $name = gethostbyaddr($client_addr, AF_INET );
   print NEW_SOCKET "Msg come from server end!";
   print "Connection recieved from $name\n";
   close NEW_SOCKET;
}

