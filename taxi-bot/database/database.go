package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var host = os.Getenv("HOST")
var port = os.Getenv("PORT")
var user = os.Getenv("USER")
var password = os.Getenv("PASSWORD")
var dbname = os.Getenv("DBNAME")

var connectionString = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func CheckTable() error {
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		fmt.Println(err)
		return err
	} else {
		if _, err := db.Exec("CREATE TABLE IF NOT EXISTS TaxiCars (Id SERIAL PRIMARY KEY, UserId BIGINT, Text TEXT)"); err != nil {
			fmt.Println(err)
			return err
		}
	}
	defer db.Close()
	return nil
}

func SetCar(userId int64, text string) error {
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		fmt.Println(err)
		return err
	} else {
		if _, err := db.Exec("INSERT INTO TaxiCars (userid, text) VALUES ('" + strconv.FormatInt(userId, 10) + "', '" + text + "')"); err != nil {
			fmt.Println(err)
			return err
		}
	}
	defer db.Close()

	return nil
}

func DeleteCar(userId int64, text string) error {
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		fmt.Println(err)
		return err
	} else {
		if _, err := db.Exec("DELETE FROM TaxiCars WHERE userid ='" + strconv.FormatInt(userId, 10) + "' AND text ='" + text + "'"); err != nil {
			fmt.Println(err)
			return err
		}
	}
	defer db.Close()

	return nil
}

func GetListCars(userId int64) ([]string, error) {
	db, err := sql.Open("postgres", connectionString)
	defer db.Close()

	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		rows, err := db.Query("SELECT * FROM TaxiCars WHERE userid = '" + strconv.FormatInt(userId, 10) + "'")
		defer rows.Close()

		if err != nil {
			fmt.Println(err)
			return nil, err
		} else {
			carss := []string{}

			for rows.Next() {
				t := Car{}
				err := rows.Scan(&t.id, &t.userId, &t.text)

				if err != nil {
					fmt.Println(err)
					continue
				}

				carss = append(carss, t.text)
			}

			return carss, nil
		}
	}
}

type Car struct {
	id     int
	userId int64
	text   string
}
