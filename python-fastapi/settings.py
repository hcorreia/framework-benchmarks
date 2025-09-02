import os

from dotenv import load_dotenv

load_dotenv()


config = {
    "chaos_endpoint": os.getenv("CHAOS_ENDPOINT"),
}
