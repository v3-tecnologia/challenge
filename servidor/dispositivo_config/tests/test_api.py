from fastapi.testclient import TestClient
from dispositivo_config.main import app

client = TestClient(app)


def test_configurar_dispositivo_sucesso():
    payload = {"placa": "ABC1234", "volume_alerta": 50}
    response = client.post("/api/dispositivos/configurar", json=payload)
    assert response.status_code == 200
    assert "mensagem" in response.json()


def test_configurar_dispositivo_falha_validacao():
    payload = {"placa": "###123", "volume_alerta": 150}
    response = client.post("/api/dispositivos/configurar", json=payload)
    assert response.status_code == 422
