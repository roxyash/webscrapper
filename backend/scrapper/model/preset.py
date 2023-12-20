from pydantic import BaseModel


class Preset(BaseModel):
    login: str
    password: str
    date_from: str
    date_to: str
    urls: list[str]
