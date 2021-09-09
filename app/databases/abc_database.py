from abc import ABCMeta, abstractmethod


from .collections import BookCollectionABC


class DatabaseABC(metaclass=ABCMeta):
    @property
    @abstractmethod
    def books() -> BookCollectionABC:
        pass
