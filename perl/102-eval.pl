#!/bin/env perl
use strict;
use warnings;
use utf8;
use POSIX qw/strftime/;
use Time::Piece;
my $t;
eval {
    my $format = "%a %b %d %H:%M:%S %Z %Y";
    my $s = strftime($format, gmtime());
    $t = Time::Piece->strptime($s, $format); # error!
};
if ($@) {
    warn $@;
    $t = localtime;
}
print $t->epoch;
print "end\n";
