from fastapi import FastAPI

app = FastAPI()

@app.get("/")
def getRoute():
    return "Hello"
