#!/bin/env perl
use strict;
use warnings;
use utf8;
use Time::Piece;
use DDP;
use POSIX qw/strftime/;
my $format = "%a %b %d %H:%m:%S UTC %Y";
my $s = strftime($format, gmtime());
p $s; #e.g. "Sun Dec 30 02:12:55 JST 2018"
#my $t = Time::Piece->strptime($s, $format);
# Cannot convert %Z, therefore %Z need to replace a static string - e.g. UTC -.
my $t = Time::Piece->strptime($s, $format);
p $t->strftime($format);
