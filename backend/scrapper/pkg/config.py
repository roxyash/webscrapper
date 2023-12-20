import os
from typing import Optional

from pydantic_settings import BaseSettings


class Settings(BaseSettings):
    # Sharing settings
    mode: Optional[str] = os.getenv("MODE")

    # Scrapper settings
    scrapper_port: Optional[int] = os.getenv("SCRAPPER_PORT")
    scrapper_host: Optional[str] = os.getenv("SCRAPPER_HOST")
    scrapper_count_workers: Optional[int] = os.getenv("SCRAPPER_COUNT_WORKERS")

    class Config:
        env_file = "config/.env"
        extra = "ignore"


settings = Settings(_env_file='.env')


def get_config():
    return settings
