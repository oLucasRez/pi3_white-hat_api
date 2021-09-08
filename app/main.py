from app.repositories import BooksRepository
from app.businesses import BooksBusiness


class App:
    def __init__(self):
        self.init_repositories()
        self.init_businesses()

    def init_repositories(self):
        self.books_repository = BooksRepository()

    def init_businesses(self):
        self.books_bo = BooksBusiness(self.books_repository)


main = App()
