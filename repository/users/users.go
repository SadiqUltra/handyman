package users

import (
	"database/sql"
	database "github.com/sadiqultra/handyman/db/migrations/postgres"
	"golang.org/x/crypto/bcrypt"

	"log"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	City     string `json:"city"`
	Password string `json:"password"`
}

func (user *User) Create() {
	statement, err := database.Db.Prepare(
		"INSERT INTO Users(Username,Phone,Address,City,Password) VALUES($1,$2,$3,$4,$5)")

	if err != nil {
		log.Fatal(err)
	}
	hashedPassword, err := HashPassword(user.Password)
	_, err = statement.Exec(user.Username, user.Phone, user.Address, user.City, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}
}

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (user *User) Authenticate() bool {
	statement, err := database.Db.Prepare("select Password from Users WHERE Username = $1")
	if err != nil {
		log.Fatal(err)
	}
	row := statement.QueryRow(user.Username)

	var hashedPassword string
	err = row.Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err)
		}
	}

	return CheckPasswordHash(user.Password, hashedPassword)
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//GetUserIdByUsername check if a user exists in database by given username
func GetUserIdByUsername(username string) (int, error) {
	statement, err := database.Db.Prepare("select ID from Users WHERE Username = $1")
	if err != nil {
		log.Fatal(err)
	}
	row := statement.QueryRow(username)

	var Id int
	err = row.Scan(&Id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return 0, err
	}

	return Id, nil
}
