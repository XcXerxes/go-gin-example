/*
 * @Description: 123
 * @Author: leo
 * @Date: 2020-02-17 16:32:59
 * @LastEditors: leo
 * @LastEditTime: 2020-02-17 17:01:56
 */

package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
)

type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

func bookableDate(fl validator.FieldLevel) bool {
	today := time.Now()
	date := fl.Field().Interface().(time.Time)
	return date.Unix() > today.Unix()
}

func main() {
	r := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}
	r.GET("/bookable", func(c *gin.Context) {
		var b Booking
		if err := c.ShouldBind(&b); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "ok!", "booking": b})
	})

	r.Run()
}
