#!/usr/bin/env perl
# secrets.pl
# Perl script with embedded secrets

use strict;
use warnings;
use DBI;

# Database configuration with hardcoded credentials
my $DATABASE_PASSWORD = 'perl_secret_password_2024!';
my $db_host = 'prod-perl.example.com';
my $db_user = 'perl_app_user';

# API Keys and tokens
my %api_config = (
    'GITHUB_TOKEN' => 'ghp_1234567890abcdefghijklmnopqrstuvwxyz123',
    'STRIPE_SECRET_KEY' => 'sk_live_51Example123456789012345678901234567890123456',
    'AWS_ACCESS_KEY_ID' => 'AKIAIOSFODNN7EXAMPLE',
    'AWS_SECRET_ACCESS_KEY' => 'wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY',
    'SENDGRID_API_KEY' => 'SG.AbCdEfGhIjKlMnOpQrStUvWxYz123456789.012345678901234567890abcdefghijklmnop'
);

# Social media configuration
my $TWITTER_CONSUMER_KEY = 'AbCdEfGhIjKlMnOpQrStU';
my $TWITTER_CONSUMER_SECRET = '1234567890abcdefghijklmnopqrstuvwxyz1234567890ab';
my $SLACK_BOT_TOKEN = 'xoxb-1234567890123-4567890123456-AbCdEfGhIjKlMnOpQrStUvWx';

# Google services
my $GOOGLE_CLIENT_ID = '123456789012-abcdefghijklmnopqrstuvwxyz123456.apps.googleusercontent.com';
my $GOOGLE_CLIENT_SECRET = 'GOCSPX-AbCdEfGhIjKlMnOpQrStUvWxYz123456';
my $GOOGLE_MAPS_API_KEY = 'AIzaSyBcdefghijklmnopqrstuvwxyz1234567890';

# Email service configuration
my %email_config = (
    'MAILGUN_API_KEY' => 'key-1234567890abcdef1234567890abcdef',
    'smtp_password' => 'email_password_perl_123!'
);

# Security settings
my $JWT_SECRET = 'perl-jwt-secret-key-for-token-signing';
my $SESSION_SECRET = 'perl-session-secret-for-auth';

# Database connection with embedded password
my $dsn = "DBI:Pg:dbname=myapp;host=prod.db.example.com;port=5432";
my $dbh = DBI->connect($dsn, $db_user, $DATABASE_PASSWORD, {
    RaiseError => 1,
    AutoCommit => 1
}) or die "Cannot connect to database: $DBI::errstr";

# Firebase configuration
my %firebase_config = (
    'FIREBASE_API_KEY' => 'AIzaSyBcdefghijklmnopqrstuvwxyz1234567890',
    'project_id' => 'myapp-12345',
    'FIREBASE_PROJECT_ID' => 'myapp-12345'
);

# Third-party services
my $TWILIO_ACCOUNT_SID = 'ACabcdef1234567890abcdef1234567890ab';
my $TWILIO_AUTH_TOKEN = 'your_auth_token_here_32_chars_long';
my $CONTENTFUL_ACCESS_TOKEN = 'your-contentful-access-token-here-64-chars-long-1234567890ab';

# OAuth configuration
my $CLIENT_SECRET = 'perl-oauth-client-secret-hardcoded';

print "Perl application started with hardcoded secrets (security risk!)\n";

# Function that uses API key
sub call_external_api {
    my $api_key = $api_config{'GITHUB_TOKEN'};
    print "Making API call with token: $api_key\n";
    # This would normally make an actual API call
}

# Close database connection
$dbh->disconnect();

1;
