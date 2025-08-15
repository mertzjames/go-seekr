import { Injectable } from '@angular/core';

// api.service.ts
// TypeScript service with hardcoded API credentials

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  
  // Firebase configuration
  private firebaseConfig = {
    apiKey: "AIzaSyBcdefghijklmnopqrstuvwxyz1234567890",
    authDomain: "myapp-12345.firebaseapp.com",
    projectId: "myapp-12345",
    storageBucket: "myapp-12345.appspot.com",
    messagingSenderId: "123456789012",
    appId: "1:123456789012:web:abcdef1234567890123456"
  };

  // API endpoints and keys
  private readonly API_ENDPOINTS = {
    stripe: {
      publishableKey: 'pk_test_51Example123456789012345678901234567890123456',
      secretKey: 'sk_test_51Example123456789012345678901234567890123456'
    },
    github: {
      token: 'ghp_1234567890abcdefghijklmnopqrstuvwxyz123',
      clientId: 'Iv1.1234567890abcdef',
      CLIENT_SECRET: '1234567890abcdef1234567890abcdef12345678'
    },
    oauth: {
      GOOGLE_CLIENT_ID: '123456789012-abcdefghijklmnopqrstuvwxyz123456.apps.googleusercontent.com',
      GOOGLE_CLIENT_SECRET: 'GOCSPX-AbCdEfGhIjKlMnOpQrStUvWxYz123456'
    }
  };

  // Database credentials (should never be in frontend!)
  private dbConfig = {
    host: 'db.example.com',
    user: 'api_user',
    DATABASE_PASSWORD: 'typescript_db_password_2024',
    database: 'production'
  };

  // Slack integration
  private slackConfig = {
    SLACK_BOT_TOKEN: 'xoxb-1234567890123-4567890123456-AbCdEfGhIjKlMnOpQrStUvWx',
    SLACK_WEBHOOK_URL: 'https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX'
  };

  // JWT secret (should be server-side only)
  private JWT_SECRET = 'your-super-secret-jwt-key-change-in-production';

  constructor() {
    console.log('API Service initialized with secrets (this is bad!)');
  }

  // Method that uses hardcoded API key
  async fetchData(): Promise<any> {
    const CONTENTFUL_ACCESS_TOKEN = 'your-contentful-access-token-here-64-chars-long-1234567890ab';
    
    return fetch(`https://cdn.contentful.com/spaces/spaceid/entries?access_token=${CONTENTFUL_ACCESS_TOKEN}`);
  }
}
