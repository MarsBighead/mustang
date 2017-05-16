#!/usr/bin/perl
use strict;
use warnings;

# hexadecimal transfer with decimal
my $src ='abcd';
my $n = unpack('H*',$src);
print "Unpacked $src value is $n.\n";

my $hex ="61626365";
my $bin = pack('H*',$hex);
print "$bin\n";

# Extract data from string
 my $data = '1234567890aABCDE,FG';
my ($one, $two, $three) = unpack('A10xA5xa2',$data);
print "$one, $two,$three\n";
#这里的一个字符，就用一个A表示。
#有多少个字符， 就在A后加一个数字，比如'ABC', 是三个字符， 就表示:A3 
#比如上面的1234567890有十个字符， 就表示A10
#而那个小写a， 我们不要， 就用x表示， 这个x表示跳过， 在pack中代表着null
#当然， 如果你要提取所有字节， 可以在后面加*, 比如A*
