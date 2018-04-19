#!/usr/bin/env perl

my $app = sub {
      my $env = shift;
	  my $cpu;
		open(STAT, '/proc/stat') or die "WTF: $!";
        while (<STAT>) {
           $cpu = $cpu . $_;    
        }
      return [
          '200',
          [ 'Content-Type' => 'text/html' ],
          [ "<html><head><meta http-equiv=\"refresh\" content=\"1\"></head><body bgcolor='#eee'>Hello from Unit / Perl, CPU stat: <pre>$cpu</pre>" ],
      ];
};
