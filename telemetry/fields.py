from django.db import models
from cryptography.fernet import Fernet, InvalidToken
from core.settings import FERNET_KEY

fernet = Fernet(FERNET_KEY.encode())

class EncryptedTextField(models.TextField):
    def get_prep_value(self, value):
        if value is None:
            return value
        return fernet.encrypt(value.encode()).decode()

    def from_db_value(self, value, expression, connection):
        if value is None:
            return value
        try:
            return fernet.decrypt(value.encode()).decode()
        except InvalidToken:
            return value  # fallback se já estiver em texto plano (opcional)

    def to_python(self, value):
        if value is None:
            return value
        try:
            return fernet.decrypt(value.encode()).decode()
        except InvalidToken:
            return value