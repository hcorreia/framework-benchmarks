FROM python:3.13-alpine


WORKDIR /app

COPY ./requirements-freeze.txt /app/requirements-freeze.txt

RUN pip install --no-cache-dir --upgrade -r /app/requirements-freeze.txt

COPY . /app

EXPOSE 8080

ENV PORT=8080

CMD fastapi run main.py --workers 1 --port $PORT
