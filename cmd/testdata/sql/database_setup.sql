-- database_setup.sql
-- SQL script with embedded credentials and sensitive data

-- Database user creation with hardcoded passwords
CREATE USER app_user WITH PASSWORD 'sql_secret_password_2024!';
CREATE USER admin_user WITH PASSWORD 'AdminPassword123!';
CREATE USER backup_user WITH PASSWORD 'BackupSecret2024';

-- Database configuration
CREATE DATABASE production_app;
GRANT ALL PRIVILEGES ON DATABASE production_app TO app_user;

-- Application configuration table with secrets
CREATE TABLE app_config (
    id SERIAL PRIMARY KEY,
    config_key VARCHAR(255) NOT NULL,
    config_value TEXT NOT NULL
);

-- Inserting API keys and secrets (bad practice!)
INSERT INTO app_config (config_key, config_value) VALUES 
('AWS_ACCESS_KEY_ID', 'AKIAIOSFODNN7EXAMPLE'),
('AWS_SECRET_ACCESS_KEY', 'wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY'),
('GITHUB_TOKEN', 'ghp_1234567890abcdefghijklmnopqrstuvwxyz123'),
('STRIPE_SECRET_KEY', 'sk_live_51Example123456789012345678901234567890123456'),
('DATABASE_PASSWORD', 'postgres_connection_password'),
('JWT_SECRET', 'your-jwt-secret-key-in-database'),
('SENDGRID_API_KEY', 'SG.AbCdEfGhIjKlMnOpQrStUvWxYz123456789.012345678901234567890abcdefghijklmnop'),
('GOOGLE_CLIENT_SECRET', 'GOCSPX-AbCdEfGhIjKlMnOpQrStUvWxYz123456'),
('SLACK_BOT_TOKEN', 'xoxb-1234567890123-4567890123456-AbCdEfGhIjKlMnOpQrStUvWx'),
('MAILGUN_API_KEY', 'key-1234567890abcdef1234567890abcdef');

-- Connection string examples in comments
-- postgresql://username:password@localhost:5432/dbname
-- mysql://user:secret_password@prod.mysql.example.com:3306/myapp
-- mongodb://admin:mongo_password_2024@cluster.mongodb.net/production

-- Service account credentials
CREATE TABLE service_accounts (
    service_name VARCHAR(100),
    CLIENT_SECRET VARCHAR(500),
    api_key VARCHAR(500)
);

INSERT INTO service_accounts VALUES 
('firebase', 'firebase-admin-sdk-credentials-json-key', 'AIzaSyBcdefghijklmnopqrstuvwxyz1234567890'),
('twilio', 'your_auth_token_here_32_chars_long', 'ACabcdef1234567890abcdef1234567890ab'),
('oauth_provider', 'oauth-client-secret-should-not-be-in-db', 'client-id-12345');

-- Backup script with embedded credentials
-- pg_dump -h prod.db.example.com -U backup_user -W myapp > backup.sql
-- Password: BackupSecret2024

-- SSL certificate references
-- /etc/ssl/private/server.key
-- /etc/ssl/certs/server.crt
