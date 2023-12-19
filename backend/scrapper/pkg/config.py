import os
from typing import Optional

from pydantic_settings import BaseSettings


class Settings(BaseSettings):
    mode: Optional[str] = os.getenv("MODE")
    port: Optional[int] = os.getenv("PORT")


class Config:
    env_file = "config/.env"
    extra = "ignore"


settings = Settings(_env_file='.env')


def get_config():
    return settings
