import uvicorn
from fastapi import FastAPI
from orders.database.database import Database
from typing import List

app = FastAPI()
_ = Database()

if __name__ == "__main__":
    uvicorn.run("app:app", host="0.0.0.0", port=8080, reload=True,
                timeout_keep_alive=3600, debug=True, workers=10)
