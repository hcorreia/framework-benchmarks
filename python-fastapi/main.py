import asyncio
import os
from typing import Union

from fastapi import Body, FastAPI, status
from fastapi.responses import JSONResponse
import requests

from settings import config

app = FastAPI()

hostname = os.uname().nodename


@app.get("/")
def home():
    return {
        "result": "Ok",
        "hostname": hostname,
    }


@app.get("/db/")
def db():
    return JSONResponse(
        status_code=status.HTTP_501_NOT_IMPLEMENTED,
        content={
            "result": "Not implemented!",
            "hostname": hostname,
        },
    )


@app.get("/chaos/")
async def chaos():
    loop = asyncio.get_event_loop()

    resp = None

    try:
        resp = await loop.run_in_executor(None, requests.get, config["chaos_endpoint"])
    except requests.exceptions.RequestException as e:
        print(f"Error connecting to chaos service: {e}")

    if resp == None or not resp.ok:
        return JSONResponse(
            status_code=status.HTTP_503_SERVICE_UNAVAILABLE,
            content={
                "result": "Chaos service is unavailable.",
                "hostname": hostname,
            },
        )

    data = resp.json()

    return {
        "result": "OK",
        "hostname": hostname,
        "sleep_time": data["sleep_time"],
    }


@app.get("/chaos-sync/")
def chaos_sync():
    resp = None

    try:
        resp = requests.get(config["chaos_endpoint"])
    except requests.exceptions.RequestException as e:
        print(f"Error connecting to chaos service: {e}")

    if resp == None or not resp.ok:
        return JSONResponse(
            status_code=status.HTTP_503_SERVICE_UNAVAILABLE,
            content={
                "result": "Chaos service is unavailable.",
                "hostname": hostname,
            },
        )

    data = resp.json()

    return {
        "result": "OK",
        "hostname": hostname,
        "sleep_time": data["sleep_time"],
    }


@app.get("/health/")
def health():
    return {
        "result": "OK",
        "hostname": hostname,
    }
