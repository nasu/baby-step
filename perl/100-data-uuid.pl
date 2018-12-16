#!/bin/env perl
use strict;
use warnings;
use utf8;
use Data::UUID;
use DDP;

sub main() {
    my $ug = Data::UUID->new;
    my $uuid = $ug->create;
    my $a, $b;
    p $b = +{ a => $a = $ug->to_string($uuid), b => length($a) };
    p $b = +{ a => $a = $ug->to_hexstring($uuid), b => length($a) };
    p $b = +{ a => $a = $ug->to_b64string($uuid), b => length($a) };
    p $b = +{ a => $a = $ug->create_hex, b => length($a) };
    p $b = +{ a => $a = $ug->create_b64, b => length($a) };
}
main();
