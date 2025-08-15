# app_config.rb
# Ruby configuration file with sensitive credentials

require 'openssl'

class AppConfig
  # Database configuration
  DATABASE_CONFIG = {
    host: 'prod-rails.example.com',
    username: 'rails_user',
    DATABASE_PASSWORD: 'ruby_secret_password_2024!',
    database: 'production_app',
    adapter: 'postgresql'
  }
  
  # API Keys and tokens
  API_KEYS = {
    stripe: {
      publishable: 'pk_live_51Example123456789012345678901234567890123456',
      STRIPE_SECRET_KEY: 'sk_live_51Example123456789012345678901234567890123456'
    },
    github: {
      GITHUB_TOKEN: 'ghp_1234567890abcdefghijklmnopqrstuvwxyz123',
      client_id: 'Iv1.1234567890abcdef',
      CLIENT_SECRET: '1234567890abcdef1234567890abcdef12345678'
    },
    aws: {
      AWS_ACCESS_KEY_ID: 'AKIAIOSFODNN7EXAMPLE',
      AWS_SECRET_ACCESS_KEY: 'wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY'
    }
  }
  
  # Social media integration
  SOCIAL_CONFIG = {
    twitter: {
      TWITTER_CONSUMER_KEY: 'AbCdEfGhIjKlMnOpQrStU',
      TWITTER_CONSUMER_SECRET: '1234567890abcdefghijklmnopqrstuvwxyz1234567890ab'
    },
    slack: {
      SLACK_BOT_TOKEN: 'xoxb-1234567890123-4567890123456-AbCdEfGhIjKlMnOpQrStUvWx',
      webhook_url: 'https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX'
    }
  }
  
  # Email service
  EMAIL_CONFIG = {
    service: 'sendgrid',
    SENDGRID_API_KEY: 'SG.AbCdEfGhIjKlMnOpQrStUvWxYz123456789.012345678901234567890abcdefghijklmnop',
    MAILGUN_API_KEY: 'key-1234567890abcdef1234567890abcdef'
  }
  
  # Security settings
  JWT_SECRET = 'your-rails-jwt-secret-key-should-be-in-env'
  SESSION_SECRET = 'another-secret-for-rails-sessions'
  SECRET_KEY_BASE = 'your-rails-secret-key-base-64-chars-long-1234567890abcdef'
  
  def self.setup_services
    # Firebase setup
    firebase_config = {
      FIREBASE_API_KEY: 'AIzaSyBcdefghijklmnopqrstuvwxyz1234567890',
      project_id: 'myapp-12345'
    }
    
    # Third-party services
    services = {
      twilio: {
        TWILIO_ACCOUNT_SID: 'ACabcdef1234567890abcdef1234567890ab',
        TWILIO_AUTH_TOKEN: 'your_auth_token_here_32_chars_long'
      },
      google: {
        GOOGLE_CLIENT_ID: '123456789012-abcdefghijklmnopqrstuvwxyz123456.apps.googleusercontent.com',
        GOOGLE_CLIENT_SECRET: 'GOCSPX-AbCdEfGhIjKlMnOpQrStUvWxYz123456',
        GOOGLE_MAPS_API_KEY: 'AIzaSyBcdefghijklmnopqrstuvwxyz1234567890'
      }
    }
    
    puts "Services configured with hardcoded secrets (bad practice!)"
  end
end

# Database connection with hardcoded password
db_url = "postgresql://user:hardcoded_password@prod.db.example.com:5432/myapp"

# OAuth client setup
oauth_client = OAuth2::Client.new(
  'your-client-id',
  'your-client-secret-should-not-be-hardcoded',
  site: 'https://api.example.com'
)
