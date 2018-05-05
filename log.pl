#!/usr/bin/perl
use Log::Log4perl qw(:easy);
use Data::Dumper;
Log::Log4perl->easy_init($ERROR);

DEBUG "This doesn't go anywhere";
ERROR "This gets logged";

# ... or standard mode for more features:

#Log::Log4perl::init('/etc/log4perl.conf');

#--or--

# Check config every 10 secs
#Log::Log4perl::init_and_watch('/etc/log4perl.conf',10);

#--then--

$logger = Log::Log4perl->get_logger('house.bedrm.desk.topdrwr');

$logger->debug('this is a debug message');
$logger->info('this is an info message');
$logger->warn('etc');
$logger->error('..');
$logger->fatal('..');


my $width=16;
my $row=[1..14];
if ($width >scalar(@$row)){
    my $w=scalar(@$row);
    map {print "-|,"; push (@$row, ""); } (0..$width-$w-1); 
}
print "\n";
print Dumper($row);
print scalar(@$row)."\n";

my $value=10;
#my $result = $value + (listop ($value + $value, $value));

