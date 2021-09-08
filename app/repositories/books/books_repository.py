from typing import List

from app.entities import Book

from .abc_books_repository import BooksRepositoryABC


class BooksRepository(BooksRepositoryABC):

    books: List[Book] = [
        Book(
            title="How to create an API",
            author="Enzo Ferrari"
        ),
        Book(
            title="Learning how to create an API",
            author="Lucas Rezende"
        ),
        Book(
            title="What the f**** is Typescript",
            author="Riccardo Cafagna"
        )
    ]

    async def list_books(self):
        return self.books
