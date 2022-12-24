package registermodels

import (
	"database/sql"
	"fmt"
	"log"

	models "backend/models"

	_ "github.com/lib/pq"
)

var Db *sql.DB
var err error

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "12345"
	dbname   = "register"
)

func CheckError(e error) {
	if e != nil {
		log.Fatal(e.Error())
	}
}

func Initialize() {
	ConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	Db, err = sql.Open(user, ConnectionString)
	CheckError(err)

}

func GetFromRegister(user_info models.User) {
	res, err := Db.Exec("INSERT INTO register(name, email, password) VALUES($1, $2, $3)", user_info.Name, user_info.Email, user_info.Password)
	CheckError(err)
	resAffected, _ := res.RowsAffected()
	fmt.Printf("Rows affected -> %d", resAffected)
}

func TakePasswordWithEmail(login_info models.LoginUser) string {
	var pass string
	err := Db.QueryRow("SELECT password FROM register WHERE email=$1", login_info.Email).Scan(&pass)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No users with that Email")

	case err != nil:
		log.Fatal(err)

	default:
		fmt.Println("Success!")
	}

	return string(pass)
}