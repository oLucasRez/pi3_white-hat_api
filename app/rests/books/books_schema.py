from typing import Optional

from pydantic import BaseModel, Field


class ListBooksResponse(BaseModel):
    title: str = Field(
        ...,
        title="The canonical title of the book",
        example="How to learn API development"
    )

    author: str = Field(
        ...,
        title="The main author of the book",
        example="Enzo Ferrari"
    )

    creation_date: Optional[str] = Field(
        None,
        title="The date the book was created",
        example="2021-09-07"
    )
