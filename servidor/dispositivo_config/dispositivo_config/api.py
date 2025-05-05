from fastapi import APIRouter, status
from dispositivo_config.schemas import DispositivoConfig

router = APIRouter()


@router.post("/api/dispositivos/configurar", status_code=status.HTTP_200_OK)
async def configurar_dispositivo(payload: DispositivoConfig):
    # No momento só retornamos
    return {"mensagem": "Comando aceito para processamento"}
