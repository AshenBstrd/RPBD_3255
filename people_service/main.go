package main

import (
	"fmt"

	"github.com/RyabovNick/databasecourse_2/golang/tasks/people_service/service/store"
)

func main() {
	conn := "postgresql://postgre:testbd:1234/postgre"
	s := store.NewStore(conn)
	fmt.Println(s.GetPeopleByID("400"))
	fmt.Println(s.ListPeople())
}
