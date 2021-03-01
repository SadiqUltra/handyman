package services

import (
	database "github.com/sadiqultra/handyman/db/migrations/postgres"
	"github.com/sadiqultra/handyman/repository/users"
	"log"
)

type Service struct {
	ID          string
	Title       string
	Description string
	Price       float64
	User        *users.User
}

func (service Service) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO Service(Title, Description, Price, UserID) VALUES($1,$2, $3)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(service.Title, service.Description, service.Price, service.User.ID)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")
	return id
}

func GetAll() []Service {
	stmt, err := database.Db.Prepare("select S.id, S.title, S.description, S.price, S.UserID, " +
		"U.Username, U.Phone, U.Address, U.City " +
		"from Services S inner join Users U on S.UserID = U.ID")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var services []Service
	var username string
	var phone string
	var address string
	var city string
	var id string
	for rows.Next() {
		var service Service
		err := rows.Scan(&service.ID, &service.Title, &service.Description,
			&service.Price, &id, &username, &phone, &address, &city)
		if err != nil {
			log.Fatal(err)
		}
		service.User = &users.User{
			ID:       id,
			Username: username,
			Phone:    phone,
			Address:  address,
			City:     city,
		}
		services = append(services, service)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return services
}
