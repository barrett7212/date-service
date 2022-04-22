package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/datetime", getDateTime)

	r.GET("/day", getDay)
	r.GET("/month", getMonth)
	r.GET("/year", getYear)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getDateTime(c *gin.Context) {
	dateTime := time.Now()
	fmt.Println("dateTime: ", dateTime)

	c.JSON(200, gin.H{
		"datetime": dateTime,
	})
}

func getDay(c *gin.Context) {
	day := time.Now().Day()
	fmt.Println("day: ", day)

	c.JSON(200, gin.H{
		"day": day,
	})
}

func getMonth(c *gin.Context) {
	month := time.Now().Month()
	fmt.Println("month: ", month)

	c.JSON(200, gin.H{
		"month": month,
	})
}

func getYear(c *gin.Context) {
	year := time.Now().Year()
	fmt.Println("year: ", year)

	c.JSON(200, gin.H{
		"year": year,
	})
}
