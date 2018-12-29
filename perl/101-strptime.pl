#!/bin/env perl
use strict;
use warnings;
use utf8;
use Time::Piece;
use DDP;
use POSIX qw/strftime/;
my $format = "%a %b %d %H:%M:%S %Z %Y";
my $s;
{
    local $ENV{TZ} = "UTC";
    $s = strftime($format, gmtime());
    p $s; #e.g. "Sun Dec 30 02:12:55 JST 2018"
}
#NOTE: Cannot convert %Z, therefore %Z need to replace a static string - e.g. UTC -.
my $format2 = $format =~ s/%Z/UTC/r;
my $t = Time::Piece->strptime($s, $format2);
p $t->strftime($format);

my $l = strftime($format,  localtime($t->epoch()));
p $l;
