#! /usr/bin/env perl

use strict;
use warnings;

my $image = `exercise38`;
$image =~ s/IMAGE://;
print $image;
