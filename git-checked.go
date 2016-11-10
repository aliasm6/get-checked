package main

import (
	"fmt"
	// "net/http"
	// "html"
	// "html/template"
	"log"
	// "encoding/json"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gopkg.in/gorp.v1"
	"strconv"
)

type TestCenter struct {
	Id               int64
	Center_name      string
	Address          string
	Days_open        string
	Time_open        int
	Time_closed      int
	Website          string
	Need_appointment bool
}

var dbmap = initDb()

func initDb() *gorp.DbMap {
	db, err := sql.Open("postgres", "postgres://Alias:password@localhost/getchecked?sslmode=disable")
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func main() {
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.GET("/centers", GetCenters)
		v1.GET("/centers/:id", GetCenter)
		v1.POST("/centers", PostCenter)
	}

	r.Run(":8000")
}

func GetCenters(c *gin.Context) {
	var centers []TestCenter
	_, err := dbmap.Select(&centers, "SELECT * FROM testing_centers")

	if err == nil {
		c.JSON(200, centers)
	} else {
		c.JSON(404, centers)
		fmt.Println(err)
	}

	fmt.Println(centers)
}

func GetCenter(c *gin.Context) {
	id := c.Params.ByName("id")
	var center TestCenter
	err := dbmap.SelectOne(&center, "SELECT * FROM testing_centers WHERE id=$1", id)

	fmt.Println(id)
	fmt.Println(center)

	if err == nil {
		center_id, _ := strconv.ParseInt(id, 0, 64)

		content := &TestCenter{
			Id:               center_id,
			Center_name:      center.Center_name,
			Address:          center.Address,
			Days_open:        center.Days_open,
			Time_open:        center.Time_open,
			Time_closed:      center.Time_closed,
			Website:          center.Website,
			Need_appointment: center.Need_appointment,
		}
		c.JSON(200, content)
	} else {
		c.JSON(404, gin.H{"error": "center not found"})
	}

	// curl -i http://localhost:8080/api/v1/user/1
}
