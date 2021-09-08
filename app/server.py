from fastapi import FastAPI

from .rests import routers

server = FastAPI()

for router in routers:
    server.include_router(router)
