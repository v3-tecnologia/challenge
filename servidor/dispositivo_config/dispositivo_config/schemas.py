from pydantic import BaseModel, Field, validator


class DispositivoConfig(BaseModel):
    placa: str = Field(..., min_length=1, max_length=10)
    volume_alerta: int = Field(..., ge=0, le=100)

    @validator("placa")
    def validar_placa(cls, v):
        if not v.isalnum():
            raise ValueError("Placa deve conter apenas letras e números")
        return v.upper()
