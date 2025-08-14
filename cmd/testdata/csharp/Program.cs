using System;
using System.Configuration;

// Program.cs
// C# application with hardcoded secrets

namespace MyApp.Configuration
{
    public class AppConfig
    {
        // Database connection strings
        public static readonly string DatabaseConnectionString = 
            "Server=prod-sql.example.com;Database=MyApp;User Id=app_user;PASSWORD=CSharpSecret2024!;";
        
        // API Keys
        public const string GITHUB_TOKEN = "ghp_1234567890abcdefghijklmnopqrstuvwxyz123";
        public const string STRIPE_SECRET_KEY = "sk_live_51Example123456789012345678901234567890123456";
        public const string SENDGRID_API_KEY = "SG.AbCdEfGhIjKlMnOpQrStUvWxYz123456789.012345678901234567890abcdefghijklmnop";
        
        // Cloud service credentials
        public static class AWS
        {
            public const string AWS_ACCESS_KEY_ID = "AKIAIOSFODNN7EXAMPLE";
            public const string AWS_SECRET_ACCESS_KEY = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY";
            public const string AWS_DEFAULT_REGION = "us-east-1";
        }
        
        // OAuth configuration
        public static class OAuth
        {
            public const string GOOGLE_CLIENT_ID = "123456789012-abcdefghijklmnopqrstuvwxyz123456.apps.googleusercontent.com";
            public const string GOOGLE_CLIENT_SECRET = "GOCSPX-AbCdEfGhIjKlMnOpQrStUvWxYz123456";
            public const string CLIENT_SECRET = "your-oauth-client-secret-here";
        }
        
        // JWT and security
        public const string JWT_SECRET = "your-super-secret-jwt-signing-key-256-bits";
        public const string ENCRYPTION_KEY = "32-byte-encryption-key-for-sensitive";
    }
    
    public class Program
    {
        public static void Main(string[] args)
        {
            // Reading configuration (bad practice - hardcoded)
            string twilioSid = "ACabcdef1234567890abcdef1234567890ab";
            string TWILIO_AUTH_TOKEN = "your_auth_token_here_32_chars_long";
            
            // Email service setup
            var emailConfig = new
            {
                MAILGUN_API_KEY = "key-1234567890abcdef1234567890abcdef",
                SmtpPassword = "email_password_123!"
            };
            
            // Firebase configuration
            var firebaseConfig = new
            {
                FIREBASE_API_KEY = "AIzaSyBcdefghijklmnopqrstuvwxyz1234567890",
                FIREBASE_PROJECT_ID = "myapp-12345"
            };
            
            Console.WriteLine("Application started with embedded secrets (security risk!)");
        }
    }
}
