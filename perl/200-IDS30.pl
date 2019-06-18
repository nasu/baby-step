#!/usr/bin/env perl
use strict;
use warnings;
use utf8;

# cf. https://wiki.sei.cmu.edu/confluence/display/perl/IDS30-PL.+Exclude+user+input+from+format+strings
my $host = `hostname`;
chop($host);
my $prompt = "$ENV{USER}\@$host";
 
sub validate_password {
  my ($password) = @_;
  my $is_ok = ($password eq "goodpass");
  printf "$prompt: Password ok? %d\n", $is_ok;
  print $is_ok . "\n";;
  return $is_ok;
};
 
 
if (validate_password( $ARGV[0])) {
  print "$prompt: access granted\n";
} else {
  print "$prompt: access denied\n";
};
