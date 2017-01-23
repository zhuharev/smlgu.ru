package controllers

import (
	"modules/middleware"
)

func Test(c *middleware.Context) {

	c.HTML(200, "test")
}
