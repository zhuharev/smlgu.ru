package v1

import (
	"fmt"
	"models"
	"modules/middleware"
)

func TeachersList(c *middleware.Context) {
	list, e := models.TeachersGetList()
	if e != nil {
		panic(e)
	}
	c.JSON(200, list)
}

func TeachersGetVotes(c *middleware.Context) {
	fmt.Println(c.QueryInt64("as_user"), "as_user")
	votes, e := models.TeachersGetVotes(c.ParamsInt64(":id"), c.QueryInt64("as_user"))
	if e != nil {
		panic(e)
	}
	c.JSON(200, votes)
}

func TeachersGet(c *middleware.Context) {
	guru, e := models.GuruGet(c.ParamsInt64(":id"))
	if e != nil {
		c.JSON(500, e)
		return
	}
	c.JSON(200, guru)
}
