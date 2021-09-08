from typing import List

from app.entities import Book
from app.repositories import BooksRepositoryABC


class BooksBusiness:
    def __init__(self, books_repository: BooksRepositoryABC):
        self.books_repository = books_repository

    async def list_books(self) -> List[Book]:
        return await self.books_repository.list_books()
