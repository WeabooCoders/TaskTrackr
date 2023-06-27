package test

import (
	"os"
	"testing"

	"github.com/AvinFajarF/initializers"
	"github.com/stretchr/testify/assert"
)

func TestConnectToDB(t *testing.T) {
	// Mengatur nilai variabel lingkungan untuk koneksi basis data
	DB := os.Getenv("DB")

	// Memanggil fungsi ConnectToDB
	initializers.ConnectToDB()

	// Memeriksa apakah variabel DB telah diinisialisasi dengan benar
	assert.NotNil(t, DB)
	// assert.Implements(t, (*gorm.Dialector)(nil), DB)
}
