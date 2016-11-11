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
	// "strconv"
)

//TestCenter structure
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

//initialization of connection to database
var dbmap = initDb()

func initDb() *gorp.DbMap {
	db, err := sql.Open("postgres", "postgres://Alias:password@localhost/getchecked?sslmode=disable")
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	return dbmap
}

//check error function
func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

// main function where routes live
func main() {
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.GET("/centers", GetCenters)
		v1.GET("/centers/:id", GetCenter)
		// v1.POST("/centers", PostCenter)
	}

	r.Run(":8000")
}

//helper functions for routes
func GetCenters(c *gin.Context) {
	var centers []TestCenter
	_, err := dbmap.Select(&centers, "SELECT * FROM testing_centers")

	if err == nil {
		c.JSON(200, centers)
	} else {
		c.JSON(404)
		fmt.Println(err)
	}

	fmt.Println(centers)
}

func GetCenter(c *gin.Context) {

	id := c.Params.ByName("id")
	var center TestCenter
	err := dbmap.SelectOne(&center, "SELECT * FROM testing_centers WHERE id=$1", id)
	fmt.Println(id)
	if err == nil {
		c.JSON(200, center)
	} else {
		c.JSON(404, gin.H{"error": "center not found"})
	}

}

func PostUser(c *gin.Context) {
	var center TestCenter
	c.Bind(&center)
	if center.Name != "" && center.Address != "" && center.Days_open != nil && center.Time_open != nil && center.Website != "" && center.Need_appointment != nil  {
		insert, _ := dbmap.Exec(`INSERT INTO trial (Id, Center_name, Address, Days_open, Time_open, Time_closed, Website, Need_appointment)
    VALUES ($1, $2)`, center.Name, center.Email)

		if insert != nil {
			c.JSON(201, gin.H{"success": "Added User"})
		} else {
			c.JSON(501, gin.H{"failed": "Insert failed"})
		}
	} else {
		c.JSON(422, gin.H{"error": "fields are empty"})
	}
	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"name\": \"Beau\", \"email\": \"b@o.com\" }"
	// http://localhost:8080/api/v1/users
}
