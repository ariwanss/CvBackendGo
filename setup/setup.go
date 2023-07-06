package setup

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ariwanss/CvBackendGo/entity"
	"github.com/ariwanss/CvBackendGo/service"
)

func purgeDb() {
	service.DropUserCol()
	service.DropCvItemCol()
}

func populateDb() {
	users := []entity.User{
		{Username: "user1", Password: "user1", Name: "User One", DOB: time.Date(1997, 8, 29, 0, 0, 0, 0, time.Local), Email: "user1@gmail.com", Phone: "082133712080", Address: "Los Angeles, California", Summary: "I am user one"},
		{Username: "user2", Password: "user2", Name: "User Two", DOB: time.Date(2001, 4, 2, 0, 0, 0, 0, time.Local), Email: "user2@gmail.com", Phone: "082356789874", Address: "Boston, Massachussetts", Summary: "I am user two"},
		{Username: "user3", Password: "user3", Name: "User Three", DOB: time.Date(1999, 7, 27, 0, 0, 0, 0, time.Local), Email: "user3@gmail.com", Phone: "081234567890", Address: "Denver, Colorado", Summary: "I am user three"},
	}

	cvItemsOne := []entity.CvItem{
		{Section: "Experience",
			Title: "Database Administrator", DateStart: time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local), DateEnd: time.Date(2012, 1, 1, 0, 0, 0, 0, time.Local)},
		{Section: "Experience",
			Title: "Data Engineer", DateStart: time.Date(2012, 1, 1, 0, 0, 0, 0, time.Local), DateEnd: time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)},
	}
	cvItemsTwo := []entity.CvItem{
		{Section: "Experience",
			Title: "Associate Software Engineer", DateStart: time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local), DateEnd: time.Date(2012, 1, 1, 0, 0, 0, 0, time.Local)},
		{Section: "Experience",
			Title: "Software Engineer", DateStart: time.Date(2012, 1, 1, 0, 0, 0, 0, time.Local), DateEnd: time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)},
	}
	cvItemsThree := []entity.CvItem{
		{Section: "Experience",
			Title: "Accountant", DateStart: time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local), DateEnd: time.Date(2012, 1, 1, 0, 0, 0, 0, time.Local)},
		{Section: "Experience",
			Title: "Accounting Manager", DateStart: time.Date(2012, 1, 1, 0, 0, 0, 0, time.Local), DateEnd: time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)},
	}

	cvItemsAll := [][]entity.CvItem{cvItemsOne, cvItemsTwo, cvItemsThree}

	userCount := 0
	cvItemCount := 0

	for i, user := range users {
		inserted, err := service.CreateUser(&user)
		if err != nil {
			panic("failed to insert a user")
		}
		userCount++
		for _, cvItem := range cvItemsAll[i] {
			cvItem.UserID = inserted.ID
			_, err = service.InsertCvItem(&cvItem)
			if err != nil {
				panic("failed to insert a cv item")
			}
			cvItemCount++
		}
	}
	fmt.Println("inserted " + strconv.Itoa(userCount) + " users")
	fmt.Println("inserted " + strconv.Itoa(cvItemCount) + " cv items")
}

func RunSetup() {
	purgeDb()
	populateDb()
}
