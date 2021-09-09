import motor
import motor.core as motor
from motor.motor_asyncio import AsyncIOMotorClient

from .abc_database import DatabaseABC
from .collections import BookCollectionABC


class Mongo(DatabaseABC):
    def __init__(self):
        self.client = AsyncIOMotorClient('database', 27017)
        self.database: motor.AgnosticDatabase = self.client["whitehat"]

    @property
    def books(self) -> BookCollectionABC:
        return self.database.get_collection("books")
