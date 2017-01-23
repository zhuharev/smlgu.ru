package models

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"

	"time"

	"models/object"
)

var (
	x *xorm.Engine
)

func NewContext() (e error) {
	e = NewBlobLogContext()
	if e != nil {
		return e
	}
	return NewEngine()
}

func NewEngine() (e error) {
	x, e = xorm.NewEngine("sqlite3", "db.sqlite")
	if e != nil {
		return
	}
	x.Exec("PRAGMA synchronous = OFF")
	x.Sync2(new(OldGuruVotes), new(Guru), new(User), new(Coment))
	return
}

/*func main() {
	x, e := xorm.NewEngine("sqlite3", "db.sqlite")
	if e != nil {
		panic(e)
	}
}*/

type OldGuruVotes struct {
	Uid      int64
	Guid     int64
	VoteType int64
	Val      int64
}

func InsertVote(v *OldGuruVotes) (e error) {
	_, e = x.Insert(v)
	return e
}

type User struct {
	Id   int64
	VkId int64 //`xorm:"unique"`

	FirstName  string
	LastName   string
	Patronymic string

	Username string `xorm:"unique"`

	AvatarBlobId int64
}

type Guru struct {
	Id     int64
	UserId int64 `xorm:"index"`

	ChairId     int64
	Rate        int64
	Initiales   string
	Description string
	FacultyId   int64

	User     *User         `xorm:"-"`
	Features *TeacherVotes `xorm:"-"`
}

func GuruCreate(g *Guru) (e error) {
	_, e = x.Insert(g)
	return e
}

func UserCreate(u *User) (e error) {
	_, e = x.Insert(u)
	return
}

func UserGet(id int64) (*User, error) {
	u := new(User)
	_, e := x.Id(id).Get(u)
	return u, e
}

type Coment struct {
	Id int64

	ComentType int64 // hidden, public etc

	ObjectType object.Object // photo, user, guru
	ObjectId   int64

	AuthorName string // if anonimus
	AuthorId   int64
	Body       string

	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
	Deleted time.Time `xorm:"deleted"`
}

func ComentCreate(c *Coment) (e error) {
	_, e = x.Insert(c)
	return
}

func VoteCreate(v *OldGuruVotes) (e error) {
	_, e = x.Insert(v)
	return
}

func guruGet(tok interface{}) (*Guru, error) {
	var (
		g = new(Guru)
	)

	switch tok.(type) {
	case string:
		slug := tok.(string)
		_, e := x.Sql("select * from guru where user_id = (select id from user where username = ?)",
			slug).Get(g)
		if e != nil {
			return nil, e
		}
	case int64:
		id := tok.(int64)
		_, e := x.Id(id).Get(g)
		if e != nil {
			return nil, e
		}
	}

	votes, e := TeachersGetVotes(g.Id)
	if e != nil {
		return nil, e
	}
	g.Features = votes

	g.User, e = UserGet(g.UserId)
	if e != nil {
		return nil, e
	}

	return g, nil

}

func GuruGetBySlug(slug string) (*Guru, error) {
	return guruGet(slug)
}

func GuruGet(id int64) (*Guru, error) {
	return guruGet(id)
}

func GuruList() ([]*Guru, error) {
	var (
		gurus = []*Guru{}
	)
	e := x.Find(&gurus)
	if e != nil {
		return nil, e
	}

	users, e := guruGetUsers()
	if e != nil {
		return nil, e
	}

	for i, guru := range gurus {
		for _, user := range users {
			if guru.UserId == user.Id {
				gurus[i].User = user
			}
		}
	}

	return gurus, e
}

func guruGetUsers() ([]*User, error) {
	var users []*User
	e := x.Sql("select * from user where id in (select user_id from guru)").Find(&users)
	return users, e
}
