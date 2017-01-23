package models

import (
	"github.com/Unknwon/com"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/sisteamnik/guseful/chpu"

	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"

	"models/object"
)

func Migrate1() error {

	var engine *xorm.Engine
	var err error
	engine, err = xorm.NewEngine("mysql", "zhu:pass@tcp(localhost:3306)/smolgu?charset=utf8")
	if err != nil {
		return err
	}

	type OldGuru struct {
		Id          int64
		FirstName   string
		LastName    string
		Patronymic  string
		ChairId     int64
		Rate        int64
		Img         string
		Initiales   string
		Description string
		FacultyId   int64
	}

	gurus := []OldGuru{}
	err = engine.Table("university_gurus").Find(&gurus)
	if err != nil {
		return err
	}

	var unique = map[string]struct{}{}

	for _, v := range gurus {
		p := filepath.Join("./img", v.Img)
		//fmt.Print("check", p, " ")
		if !com.IsExist(p) {
			fmt.Println(p)
		}
		if _, has := unique[chpu.Chpu(v.LastName+" "+v.Initiales)]; has {
			panic("uniquer failed " + chpu.Chpu(v.LastName+" "+v.Initiales))
		} else {
			unique[chpu.Chpu(v.LastName+" "+v.Initiales)] = struct{}{}
		}

	}

	return nil
}

func Migrate() error {

	var engine *xorm.Engine

	type OldUser struct {
		Id   int64
		VkId int64
	}

	type OldComent struct {
		Id      int64
		Text    string
		Created time.Time
		Author  string
		Uid     int64
		Guid    int64
		Vis     int64
	}

	type OldVote struct {
		Uid      int64
		GuruId   int64
		Votetype int64
		Val      int64
	}

	type OldGuru struct {
		Id          int64
		FirstName   string
		LastName    string
		Patronymic  string
		ChairId     int64
		Rate        int64
		Img         string
		Initiales   string
		Description string
		FacultyId   int64
	}

	type OldImg struct {
		Id          int64
		Name        string
		Description string
		TId         int64 `xorm:"t_id"`
	}

	var err error
	engine, err = xorm.NewEngine("mysql", "zhu:pass@tcp(localhost:3306)/smolgu?charset=utf8")
	if err != nil {
		return err
	}
	users := []OldUser{}
	err = engine.Table("users").Find(&users)
	if err != nil {
		return err
	}

	coments := []OldComent{}
	err = engine.Table("university_gurus_coments").Find(&coments)
	if err != nil {
		return err
	}

	votes := []OldVote{}
	err = engine.Table("university_gurus_votes").Find(&votes)
	if err != nil {
		return err
	}

	gurus := []OldGuru{}
	err = engine.Table("university_gurus").Find(&gurus)
	if err != nil {
		return err
	}

	images := []OldImg{}
	err = engine.Table("imgs").Find(&images)
	if err != nil {
		return err
	}

	_ = users
	_ = coments
	_ = votes
	_ = gurus
	_ = images

	for _, v := range users {
		user := User{
			Id:       v.Id,
			VkId:     v.VkId,
			Username: fmt.Sprint(v.VkId),
		}
		err = UserCreate(&user)
		if err != nil {
			return err
		}
	}

	for _, v := range coments {
		coment := Coment{
			Id: v.Id,

			ObjectType: object.Guru,
			ObjectId:   v.Guid,

			AuthorId: v.Id,
			Body:     v.Text,
			Created:  v.Created,
		}

		err = ComentCreate(&coment)
		if err != nil {
			return err
		}
	}

	for _, v := range gurus {

		p := filepath.Join("./img", v.Img)
		bts, e := ioutil.ReadFile(p)
		if e != nil {
			return e
		}

		blobId, e := InsertBlob(bts)
		if e != nil {
			return e
		}

		user := User{
			LastName:     v.LastName,
			FirstName:    v.Initiales,
			AvatarBlobId: blobId,
			Username:     chpu.Chpu(v.LastName + " " + v.Initiales),
		}

		e = UserCreate(&user)
		if e != nil {
			return e
		}

		guru := Guru{
			Id: v.Id,

			FacultyId: v.FacultyId,
			ChairId:   v.ChairId,
			Rate:      v.Rate,

			UserId: user.Id,
		}

		err = GuruCreate(&guru)
		if err != nil {
			return err
		}
	}

	for _, v := range votes {
		vote := OldGuruVotes{
			Uid:      v.Uid,
			Guid:     v.GuruId,
			VoteType: v.Votetype,
			Val:      v.Val,
		}
		err = VoteCreate(&vote)
		if err != nil {
			return err
		}
	}

	return nil
}
