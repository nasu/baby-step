#!/usr/bin/env perl
use strict;
use warnings;
use utf8;

local $\="\n";
my @a = (3,4,2,5,1,8,6,10,9,7);
print join ' ', @a;
print join ' ', sort @a;
print join ' ', sort { $a <=> $b } @a;
print join ' ', (sort { $a <=> $b } @a)[0..0];
print join ' ', (sort { $a <=> $b } @a)[0..9];

