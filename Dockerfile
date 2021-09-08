FROM python:3.9.6-slim-buster

WORKDIR /usr/app

COPY requirements.txt .

RUN pip3 install -r requirements.txt

COPY . .

CMD ["uvicorn", "app.server:server", "--reload", "--host", "0.0.0.0"]