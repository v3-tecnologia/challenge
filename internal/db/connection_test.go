package db

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetConnection_WithMock(t *testing.T) {
	mockDB, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("erro ao criar mock do banco: %v", err)
	}
	defer mockDB.Close()

	SetConnection(mockDB)

	conn, err := GetConnection()
	if err != nil {
		t.Errorf("esperado sucesso, recebido erro: %v", err)
	}
	if conn != mockDB {
		t.Errorf("conexão retornada não é a mockada")
	}
}
