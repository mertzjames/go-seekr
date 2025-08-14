-- config.lua
-- Lua configuration script with embedded secrets

local config = {}

-- Database configuration
config.database = {
    host = "prod-lua.example.com",
    port = 3306,
    username = "lua_user",
    DATABASE_PASSWORD = "lua_secret_password_2024!",
    database = "production_app",
    ssl = true
}

-- API Keys and authentication
config.api_keys = {
    GITHUB_TOKEN = "ghp_1234567890abcdefghijklmnopqrstuvwxyz123",
    STRIPE_SECRET_KEY = "sk_live_51Example123456789012345678901234567890123456",
    stripe_publishable = "pk_live_51Example123456789012345678901234567890123456"
}

-- AWS configuration
config.aws = {
    AWS_ACCESS_KEY_ID = "AKIAIOSFODNN7EXAMPLE",
    AWS_SECRET_ACCESS_KEY = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
    region = "us-west-2",
    s3_bucket = "my-secret-bucket"
}

-- Google services
config.google = {
    GOOGLE_CLIENT_ID = "123456789012-abcdefghijklmnopqrstuvwxyz123456.apps.googleusercontent.com",
    GOOGLE_CLIENT_SECRET = "GOCSPX-AbCdEfGhIjKlMnOpQrStUvWxYz123456",
    GOOGLE_MAPS_API_KEY = "AIzaSyBcdefghijklmnopqrstuvwxyz1234567890"
}

-- Email services
config.email = {
    service = "sendgrid",
    SENDGRID_API_KEY = "SG.AbCdEfGhIjKlMnOpQrStUvWxYz123456789.012345678901234567890abcdefghijklmnop",
    MAILGUN_API_KEY = "key-1234567890abcdef1234567890abcdef",
    smtp_password = "lua_email_password_123!"
}

-- Social media integration
config.social = {
    twitter = {
        TWITTER_CONSUMER_KEY = "AbCdEfGhIjKlMnOpQrStU",
        TWITTER_CONSUMER_SECRET = "1234567890abcdefghijklmnopqrstuvwxyz1234567890ab"
    },
    slack = {
        SLACK_BOT_TOKEN = "xoxb-1234567890123-4567890123456-AbCdEfGhIjKlMnOpQrStUvWx",
        webhook_url = "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX"
    }
}

-- Security settings
config.security = {
    JWT_SECRET = "lua-jwt-secret-key-for-token-verification",
    SESSION_SECRET = "lua-session-secret-for-auth-management",
    encryption_key = "32-byte-encryption-key-for-lua-app"
}

-- Firebase configuration
config.firebase = {
    FIREBASE_API_KEY = "AIzaSyBcdefghijklmnopqrstuvwxyz1234567890",
    project_id = "myapp-12345",
    FIREBASE_PROJECT_ID = "myapp-12345",
    auth_domain = "myapp-12345.firebaseapp.com"
}

-- Third-party services
config.services = {
    twilio = {
        TWILIO_ACCOUNT_SID = "ACabcdef1234567890abcdef1234567890ab",
        TWILIO_AUTH_TOKEN = "your_auth_token_here_32_chars_long"
    },
    contentful = {
        CONTENTFUL_ACCESS_TOKEN = "your-contentful-access-token-here-64-chars-long-1234567890ab",
        space_id = "lua-contentful-space"
    },
    oauth = {
        CLIENT_SECRET = "lua-oauth-client-secret-hardcoded",
        client_id = "lua-oauth-client-id-12345"
    }
}

-- Database connection string
config.connection_string = "mysql://lua_user:hardcoded_lua_password@prod.mysql.example.com:3306/myapp"

-- Redis configuration
config.redis = {
    host = "cache.example.com",
    port = 6379,
    password = "redis_lua_secret_2024"
}

-- Function to initialize services
function config.init_services()
    print("Lua application initialized with hardcoded secrets (security risk!)")
    
    -- Log database connection (bad practice)
    print("Connecting to database with password: " .. config.database.DATABASE_PASSWORD)
    
    -- API setup
    local api_token = config.api_keys.GITHUB_TOKEN
    print("Using GitHub token: " .. api_token)
end

return config
