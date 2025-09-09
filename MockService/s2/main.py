from fastapi import FastAPI

app = FastAPI()

@app.get("/")
def route():
    return {"message": "Hello from mock service"}

