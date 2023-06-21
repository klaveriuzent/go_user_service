package test

import (
	"os"
	"testing"
)

func TestIsMainFileExists(t *testing.T) {
	// Cek apakah file main.go ada di direktori root
	_, err := os.Stat("./../../main.go")
	if err != nil {
		if os.IsNotExist(err) {
			t.Fatal("File main.go tidak ditemukan di direktori root")
		} else {
			t.Fatalf("Terjadi kesalahan saat memeriksa file main.go: %v", err)
		}
	}
}
