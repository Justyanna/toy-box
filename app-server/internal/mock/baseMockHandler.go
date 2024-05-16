package mock

import (
	"database/sql"
)

type BaseMockHandler struct {
	db *sql.DB
}

func NewBaseMockHandler(db *sql.DB) *BaseMockHandler {
	return &BaseMockHandler{db: db}
}
