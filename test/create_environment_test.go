package test

import (
	"context"
	"database/sql"
	"fmt"
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

func TestDelete(t *testing.T) {
	db := setupTestDB()
	// Menggunakan db.Exec untuk mengeksekusi banyak pernyataan SQL
	defer db.Close()
	tables := []string{
		"payment_merchants",
		"installment_details",
		"transactions",
		"tenor_customers",
		"customers",
		"merchants",
		"users",
		"role",
		"interest_rates",
	}

	for _, table := range tables {
		deleteStatement := fmt.Sprintf("DELETE FROM %s", table)
		_, err := db.Exec(deleteStatement)
		if err != nil {
			fmt.Printf("Error deleting data from table %s: %v\n", table, err)
			panic(err)
		}
	}

	fmt.Println("Semua data telah dihapus.")
}

func TestCreateRole(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	SQL := "INSERT INTO role(id, name) VALUES ($1, $2)"
	_, err := db.ExecContext(context.Background(), SQL, 1, "admin")
	if err != nil {
		log.Fatal(err)
	}

	SQL2 := "INSERT INTO role(id, name) VALUES ($1, $2)"
	_, err = db.ExecContext(context.Background(), SQL2, 2, "customer")
	if err != nil {
		log.Fatal(err)
	}

	SQL3 := "INSERT INTO role(id, name) VALUES ($1, $2)"
	_, err = db.ExecContext(context.Background(), SQL3, 3, "merchant")
	if err != nil {
		log.Fatal(err)
	}

	// Tes berhasil
	t.Logf("Data role berhasil dimasukkan ke dalam database.")
}

func TestCreateUser(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	SQL := "INSERT INTO users(id, username, password, phone_number, role_id) VALUES ($1, $2, $3, $4, $5)"
	_, err := db.ExecContext(context.Background(), SQL, 1, "anwar", "123", "089832872", 2)
	if err != nil {
		log.Fatal(err)
	}

	SQL2 := "INSERT INTO users(id, username, password, phone_number, role_id) VALUES ($1, $2, $3, $4, $5)"
	_, err = db.ExecContext(context.Background(), SQL2, 2, "budi", "123", "087737823", 3)
	if err != nil {
		log.Fatal(err)
	}

	// Tes berhasil
	t.Logf("Data user berhasil dimasukkan ke dalam database.")
}

func TestCreateInterestRate(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	SQL := "INSERT INTO interest_rates(id, interest_rate, valid_date) VALUES ($1, $2, $3)"
	_, err := db.ExecContext(context.Background(), SQL, 1, 4.5, time.Now())
	if err != nil {
		log.Fatal(err)
	}

	// Tes berhasil
	t.Logf("Data user berhasil dimasukkan ke dalam database.")
}

func TestCreateCustomer(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	pin := "030201"
	hashPin := helper.Hash(pin)

	SQL := "INSERT INTO customers(id, user_id, full_name, birth_place, birth_date, salary, identity_card, selfie_photo, pin) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	_, err := db.ExecContext(context.Background(), SQL, 1, 1, "anwar maulana", "bekasi", time.Now(), 2000000, "idcard.png", "seldie.png", hashPin)
	if err != nil {
		log.Fatal(err)
	}

	// Tes berhasil
	t.Logf("Data customer berhasil dimasukkan ke dalam database.")
}

func TestCreateTenorCustomer(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	SQL := "INSERT INTO tenor_customers(id, customer_id, tenor_1, tenor_2, tenor_3, tenor_4) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := db.ExecContext(context.Background(), SQL, 1, 1, 320000, 500000, 730000, 950000)
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

	SQL := "INSERT INTO merchants(id, user_id, merchant_name, bank_account, api_key) VALUES ($1, $2, $3, $4, $5)"
	_, err = db.ExecContext(context.Background(), SQL, 1, 2, "toyota_showroom", "2212932832912", apiKey)
	if err != nil {
		log.Fatal(err)
	}

	// Tes berhasil
	t.Logf("Data merchants berhasil dimasukkan ke dalam database.")
}

//func TestUpdate(t *testing.T) {
//	db := setupTestDB()
//	defer db.Close()
//	SQLUpdate := "UPDATE tenor_customers SET tenor_2 = 50000 WHERE customer_id = $1"
//	_, _ = db.Exec(SQLUpdate, 1)
//}
