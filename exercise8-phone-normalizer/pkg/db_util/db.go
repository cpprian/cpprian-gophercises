package dbutil

import (
	"database/sql"
	"fmt"

	pkg_err "github.com/cpprian/cpprian-gophercises/exercise8-phone-normalizer/pkg/error"
	norm "github.com/cpprian/cpprian-gophercises/exercise8-phone-normalizer/pkg/phone_normalizer"
	_ "github.com/lib/pq"
)

type DB_Handle struct {
	db *sql.DB
}

type PhoneNumber struct {
	id    int
	phone string
}

func InitDbHandle() *DB_Handle {
	return &DB_Handle{}
}

func (d *DB_Handle) OpenDB() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s sslmode=disable", Host, Port, User)

	d.db, err = sql.Open("postgres", psqlInfo)
	pkg_err.DBError(err)
	d.DeleteDB()
	d.CreateDB()
	d.db.Close()

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, Dbname)
	d.db, err = sql.Open("postgres", psqlInfo)
	pkg_err.DBError(err)

	d.CreateTable()
}

func (d *DB_Handle) CloseDB() {
	d.db.Close()
}

func (d *DB_Handle) PingDB() {
	err := d.db.Ping()
	pkg_err.DBError(err)
}

func (d *DB_Handle) CreateDB() {
	_, err := d.db.Exec(Create_database)
	pkg_err.DBError(err)
}

func (d *DB_Handle) CreateTable() {
	_, err := d.db.Exec(Create_table)
	pkg_err.DBError(err)
}

func (d *DB_Handle) DeleteDB() {
	_, err := d.db.Exec(Delete_database)
	pkg_err.DBError(err)
}

func (d *DB_Handle) InsertPhoneNumber(phone_number string) {
	// regexp to check if phone number is valid and normalize it
	reg_phone := norm.PhoneNormalize(phone_number)

	// remove previous an identical phone number
	d.DeletePhoneNumber(reg_phone)

	_, err := d.db.Exec(Insert_phone_number, reg_phone)
	pkg_err.DBError(err)
}

func (d *DB_Handle) DeletePhoneNumber(phone_number string) {
	_, err := d.db.Exec(Delete_phone_number, phone_number)
	pkg_err.DBError(err)
}

func (d *DB_Handle) PrintPhoneNumbers() {
	rows, err := d.db.Query(Select_phone_number)
	pkg_err.DBError(err)

	var phone PhoneNumber
	for rows.Next() {
		err = rows.Scan(&phone.id, &phone.phone)
		pkg_err.DBError(err)
		fmt.Printf("#%d:	%s\n", phone.id, phone.phone)
	}
}
