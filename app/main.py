from app.repositories import BooksRepository
from app.businesses import BooksBusiness
from app.databases import Mongo


class App:
    def __init__(self):
        self.init_databases()
        self.init_repositories()
        self.init_businesses()

    def init_databases(self):
        self.mongo = Mongo()

    def init_repositories(self):
        if self.mongo is None:
            self.init_databases()

        self.books_repository = BooksRepository(self.mongo)

    def init_businesses(self):
        if self.books_repository is None:
            self.init_repositories()

        self.books_bo = BooksBusiness(self.books_repository)


main = App()
