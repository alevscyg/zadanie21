package sqlstorage_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = "user=postgres password=qwe host=localhost dbname=db_golang_test sslmode=disable"

	os.Exit(m.Run())
}
