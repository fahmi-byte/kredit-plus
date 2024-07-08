package main

import (
	"context"
	"database/sql"
	"github.com/go-playground/assert/v2"
	_ "github.com/lib/pq"
	"kredit-plus/helper"
	"log"
	"testing"
	"time"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:123@localhost/kredit_plus_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func TestCreateUser(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	pin := "030201"
	hashPin := helper.Hash(pin) // Ganti dengan hasil hash dari helper.Hash(pin)

	SQL := "INSERT INTO customers(user_id, full_name, birth_place, birth_date, salary, identity_card, selfie_photo, pin) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	_, err := db.ExecContext(context.Background(), SQL, 2, "anwar_maulana", "bekasi", time.Now(), 2000000, "idcard.png", "seldie.png", hashPin)
	if err != nil {
		log.Fatal(err)
	}

	// Tes berhasil
	t.Logf("Data pengguna berhasil dimasukkan ke dalam database.")
}

func TestPin(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	pin := "030201"
	hashPin := helper.Hash(pin)
	var hashedPin string

	rows, _ := db.QueryContext(context.Background(), "SELECT pin from customers where id = $1", 1)

	if rows.Next() {
		rows.Scan(&hashedPin)
	}

	assert.Equal(t, hashPin, hashedPin)
}

func TestCreateTenorCustomer(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	SQL := "INSERT INTO tenor_customers(customer_id, tenor_1, tenor_2, tenor_3, tenor_4) VALUES ($1, $2, $3, $4, $5)"
	_, err := db.ExecContext(context.Background(), SQL, 1, 320000, 500000, 730000, 950000)
	if err != nil {
		log.Fatal(err)
	}

	// Tes berhasil
	t.Logf("Data tenor customer berhasil dimasukkan ke dalam database.")
}

func TestCreateMerchants(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	apiKey, err := helper.GenerateAPIKey(32)

	SQL := "INSERT INTO merchants(user_id, merchant_name, bank_account, api_key) VALUES ($1, $2, $3, $4)"
	_, err = db.ExecContext(context.Background(), SQL, 3, "toyota_showroom", "2212932832912", apiKey)
	if err != nil {
		log.Fatal(err)
	}

	// Tes berhasil
	t.Logf("Data merchants berhasil dimasukkan ke dalam database.")
}
