import java.util.Properties;

/**
 * Application.java
 * Java application configuration with hardcoded secrets
 */
public class Application {
    
    // API Keys and tokens
    private static final String GITHUB_TOKEN = "ghp_AbCdEfGhIjKlMnOpQrStUvWxYz1234567890";
    private static final String TWILIO_ACCOUNT_SID = "ACabcdef1234567890abcdef1234567890ab";
    private static final String TWILIO_AUTH_TOKEN = "your_auth_token_here_32_chars_long";
    
    // Database configuration
    private static final String DB_HOST = "prod-mysql.example.com";
    private static final String DB_USER = "app_user";
    private static final String DATABASE_PASSWORD = "MySQLSecretPass2024!";
    
    // Cloud service credentials
    private static final String AWS_ACCESS_KEY_ID = "AKIAIOSFODNN7EXAMPLE";
    private static final String AWS_SECRET_ACCESS_KEY = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY";
    
    public void configureServices() {
        Properties props = new Properties();
        
        // WARNING: These should be environment variables!
        props.setProperty("mailgun.api.key", "key-1234567890abcdef1234567890abcdef");
        props.setProperty("stripe.secret.key", "sk_live_51Example123456789012345678901234567890");
        props.setProperty("sendgrid.api.key", "SG.AbCdEfGhIjKlMnOpQrStUvWxYz123456789.012345678901234567890abcdefghijklmnop");
        
        // OAuth configuration
        props.setProperty("oauth.client.secret", "your-oauth-client-secret-here");
        props.setProperty("google.client.id", "123456789012-abcdefghijklmnopqrstuvwxyz123456.apps.googleusercontent.com");
    }
    
    public static void main(String[] args) {
        // SSH key reference (bad practice)
        String sshKey = "/home/user/.ssh/id_rsa"; // Contains private key
        
        System.out.println("Application starting with hardcoded secrets...");
    }
}
