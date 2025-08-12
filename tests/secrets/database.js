// database.js
// Database connection utility.

const { Pool } = require('pg');

// WARNING: Hardcoded credentials are a major security risk.
const pool = new Pool({
  user: 'auth_service_user',
  host: 'db.internal.example.com',
  database: 'auth_db',
  password: 'Password123!', // A clearly hardcoded password.
  port: 5432,
});

module.exports = pool;