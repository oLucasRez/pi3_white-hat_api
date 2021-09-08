from abc import ABCMeta, abstractmethod


class BooksRepositoryABC(metaclass=ABCMeta):

    @abstractmethod
    async def list_books():
        pass
