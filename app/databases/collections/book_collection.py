from typing import AsyncIterable
from abc import ABCMeta, abstractmethod


class BookCollectionABC(metaclass=ABCMeta):
    @abstractmethod
    def find(parameters) -> AsyncIterable[dict]:
        pass
