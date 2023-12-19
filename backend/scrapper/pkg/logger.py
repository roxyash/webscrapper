import logging

# Create a custom logger
logger = logging.getLogger(__name__)

# Configure logger
logging.basicConfig(format='%(asctime)s - %(name)s - %(levelname)s - %(message)s', level=logging.INFO)


def get_logger() -> logging.Logger:
    return logger
