#!/usr/bin/env perl
use strict;
use warnings;
use utf8;
use URI;
use DDP;

my $uri = URI->new("https://example.com/foo/bar?c=3");
my $h = +{ a => "1 1", b => "2&2", d => 3 };
#$uri->query_form(%{ $uri->query_form || +{} }, %$h); # error
my %a = $uri->query_form;
$uri->query_form(%a, %$h);
print $uri->as_string . "\n";

$uri = URI->new("https://example.com/foo/bar?c=3");
my $uri2 = URI->new;
$uri2->query_form(+{a => "1 1", b => "2&2", d => 3});
$uri->fragment($uri2->query);
print $uri->as_string . "\n";
