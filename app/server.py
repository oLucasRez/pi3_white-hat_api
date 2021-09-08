import ptvsd
from fastapi import FastAPI

from .rests import routers

ptvsd.enable_attach(address=('0.0.0.0', 5678))

server = FastAPI()

for router in routers:
    server.include_router(router)
