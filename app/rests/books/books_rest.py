from typing import List
from fastapi import APIRouter

from app.main import main

from .books_schema import ListBooksResponse


router = APIRouter(prefix="/books", tags=["Books"])


@router.get(
    "",
    name="List all books",
    description="List all available books",
    response_model=List[ListBooksResponse],
)
async def list_books():
    return await main.books_bo.list_books()
