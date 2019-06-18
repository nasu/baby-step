#!/usr/bin/env perl
use strict;
use warnings;
use utf8;
use HTTP::Request;
my $req = HTTP::Request->new(GET => 'https://www.example.com/foo/bar?q=query');
print $req->uri->as_string . "\n";
