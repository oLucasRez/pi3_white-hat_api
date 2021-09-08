from typing import Optional
from pydantic import BaseModel


class Book(BaseModel):
    title: str
    author: str
    creation_date: Optional[str]
