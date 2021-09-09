from typing import List

from app.entities import Book
from app.databases import DatabaseABC

from .abc_books_repository import BooksRepositoryABC


class BooksRepository(BooksRepositoryABC):

    def __init__(self, database: DatabaseABC):
        self.collection = database.books

    async def list_books(self) -> List[Book]:
        iterable_cursor = self.collection.find({})

        books = []

        async for book in iterable_cursor:
            """
             This conversion should be inside an adapter
             But it was late in the night
             And this is just an example =)
            """
            books.append(Book(**book))

        return books
