from abc import ABCMeta, abstractmethod
from typing import List

from app.entities import Book


class BooksRepositoryABC(metaclass=ABCMeta):

    @abstractmethod
    async def list_books() -> List[Book]:
        pass
