from fastapi import FastAPI
from dispositivo_config.api import router

app = FastAPI(title="API de Configuração de Dispositivos", version="1.0.0")

app.include_router(router)
