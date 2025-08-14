<?php
// secrets.php
// PHP configuration file with embedded secrets

define('DB_HOST', 'localhost');
define('DB_USER', 'webapp');
define('DATABASE_PASSWORD', 'php_secret_password_2024');
define('DB_NAME', 'production_db');

// API Configuration
$config = array(
    'paypal_client_id' => 'AeB1234567890abcdef1234567890abcdef1234567890abcdef',
    'PAYPAL_CLIENT_SECRET' => 'EHijklmnopqrstuvwxyz1234567890abcdef1234567890abcd',
    'STRIPE_PUBLISHABLE_KEY' => 'pk_live_51Example123456789012345678901234567890123456',
    'mailchimp_api_key' => '1234567890abcdef1234567890abcdef-us1',
    'GOOGLE_MAPS_API_KEY' => 'AIzaSyBcdefghijklmnopqrstuvwxyz1234567890'
);

// Social Media API Keys
$social_config = [
    'TWITTER_CONSUMER_KEY' => 'AbCdEfGhIjKlMnOpQrStU',
    'TWITTER_CONSUMER_SECRET' => '1234567890abcdefghijklmnopqrstuvwxyz1234567890ab',
    'facebook_app_secret' => '1234567890abcdef1234567890abcdef',
    'LINKEDIN_CLIENT_SECRET' => 'AbCdEfGh'
];

// Cloud Storage
$aws_credentials = [
    'AWS_ACCESS_KEY_ID' => 'AKIAIOSFODNN7EXAMPLE',
    'AWS_SECRET_ACCESS_KEY' => 'wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY',
    'S3_BUCKET_NAME' => 'my-secret-bucket'
];

// Email service
$mail_config = [
    'SENDGRID_API_KEY' => 'SG.AbCdEfGhIjKlMnOpQrStUvWxYz123456789.012345678901234567890abcdefghijklmnop',
    'MAILGUN_API_KEY' => 'key-1234567890abcdef1234567890abcdef',
    'smtp_password' => 'email_password_123!'
];

// JWT and session secrets
define('JWT_SECRET', 'your-256-bit-secret');
define('SESSION_SECRET', 'another-secret-for-sessions');

// Database connection with embedded password
$pdo = new PDO("mysql:host=db.example.com;dbname=myapp", "username", "hardcoded_password_bad!");

?>
