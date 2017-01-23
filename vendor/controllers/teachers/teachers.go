package teachers

import (
	"models"
	"modules/middleware"

	"bytes"
	"fmt"
	"time"
)

//var users []*models.User

func List(c *middleware.Context) {
	gurus, e := models.GuruList()
	if e != nil {
		fmt.Println(e)
	}
	c.Data["gurus"] = gurus
	c.HTML(200, "teachers/list")
}

func Show(c *middleware.Context) {

	exGur := models.Guru{
		User: &models.User{
			FirstName: "Kirill",
			LastName:  "Zhuharev",
		},
	}

	c.Data["guru"] = exGur
	c.HTML(200, "teachers/show")
}

func Avatar(c *middleware.Context) {
	guru, err := models.GuruGetBySlug(c.Params(":slug"))
	if err != nil {
		fmt.Println(err)
		return
	}
	bts, err := models.GetBlob(guru.User.AvatarBlobId)
	if err != nil {
		fmt.Println(err)
		return
	}

	rdr := bytes.NewReader(bts)

	c.ServeContent(guru.User.Username, rdr, time.Now().Truncate(time.Hour*24*31))
}

func Get(c *middleware.Context) {
	guru, err := models.GuruGetBySlug(c.Params(":slug"))
	if err != nil {
		//todo
		return
	}
	c.Data["guru"] = guru
	c.HTML(200, "teachers/show")
}
