# config.py
# Application configuration file with sensitive data.

class Config:
    # A fake AWS Access Key ID. Real keys often start with "AKIA".
    AWS_ACCESS_KEY_ID = "AKIAIOSFODNN7EXAMPLE"

    # A fake AWS Secret Access Key.
    AWS_SECRET_ACCESS_KEY = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"

    # Hardcoded password in a database URI string.
    DATABASE_URL = "postgresql://user:dev_password_dont_use@prod.db.example.com:5432/mydatabase"

# Secret token for a third-party API.
STRIPE_API_KEY = "sk_test_51MexampleA1BexampleC2DexampleE3FexampleG4H"