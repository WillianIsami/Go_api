package config

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	ConnectTestDatabase()
	code := m.Run()

	db, _ := TestDB.DB()
	db.Close()

	os.Exit(code)
}
