package models

import (
//"fmt"
)

type GuruFeature int

const (
	GF_Humor GuruFeature = iota + 1
	GF_GoodWill
	GF_Understandability
)

func (gf GuruFeature) String() string {
	switch gf {
	case GF_Humor:
		return "Чувство юмора"
	case GF_GoodWill:
		return "Доброжелательность"
	case GF_Understandability:
		return "Понятность объяснений"
	}
	return ""
}

func TeachersGetList() ([]*Guru, error) {
	gurus := []*Guru{}
	e := x.Find(&gurus)
	return gurus, e
}

type TeacherVotes struct {
	HumorTotal    int
	HumorPositive int
	HumorVoted    bool

	GoodwillTotal    int
	GoodwillPositive int
	GoodwillVoted    bool

	UnderstandabilityTotal    int
	UnderstandabilityPositive int
	UnderstandabilityVoted    bool
}

func (t TeacherVotes) MarshalJSON() {

}

type Votes struct {
	Feature  GuruFeature
	Total    int
	Positive int
	Voted    bool
}

func (v Votes) PositiveWidth() int {
	Positive := v.Positive
	total := v.Total
	if Positive == 0 && total == 0 {
		return 100
	}
	return int((float64(Positive) / (float64(Positive) + float64(total))) * 100)
}

func (v Votes) Against() int {
	return v.Total - v.Positive
}

func (v Votes) AgainstWidth() int {
	return 100 - v.PositiveWidth()
}

func (v Votes) Sum() int {
	return v.Positive - (v.Total - v.Positive)
}

func (tv TeacherVotes) AsSlice() (res []Votes) {
	var (
		total    = 0
		positive = 0
		voted    = false
		feature  GuruFeature
	)
	for i := 1; i <= 3; i++ {
		switch i {
		case 1:
			total = tv.HumorTotal
			positive = tv.HumorPositive
			voted = tv.HumorVoted
			feature = GuruFeature(i)
		case 2:
			total = tv.GoodwillTotal
			positive = tv.GoodwillPositive
			voted = tv.GoodwillVoted
			feature = GuruFeature(i)
		case 3:
			total = tv.UnderstandabilityTotal
			positive = tv.UnderstandabilityPositive
			voted = tv.UnderstandabilityVoted
			feature = GuruFeature(i)
		}
		res = append(res, Votes{Feature: feature, Total: total, Positive: positive, Voted: voted})
	}
	return
}

func TeachersGetVotes(guruId int64, userIds ...int64) (*TeacherVotes, error) {
	var (
		arr          = []*TeacherVotes{}
		userId int64 = 0
	)

	if userIds != nil {
		userId = userIds[0]
	}

	e := x.Sql("select (select count(val) from old_guru_votes where guid = $1 and vote_type = 0) humor_total, "+
		"(select count(val) from old_guru_votes where guid = $1 and vote_type = 0 and val = 1) humor_positive, "+
		"(select count(*) from old_guru_votes where uid = $2 and vote_type = 0 and guid = $1) humor_voted, "+
		"(select count(val) from old_guru_votes where guid = $1 and vote_type = 1) goodwill_total, "+
		"(select count(val) from old_guru_votes where guid = $1 and vote_type = 1 and val = 1) goodwill_positive, "+
		"(select count(*) from old_guru_votes where uid = $2 and vote_type = 1 and guid = $1) goodwill_voted, "+
		"(select count(val) from old_guru_votes where guid = $1 and vote_type = 2) understandability_total, "+
		"(select count(val) from old_guru_votes where guid = $1 and vote_type = 2 and val = 1) understandability_positive, "+
		"(select count(*) from old_guru_votes where uid = $2 and vote_type = 2 and guid = $1) understandability_voted ",
		guruId, userId).
		Limit(1).
		Table("old_guru_votes").
		Find(&arr)
	if len(arr) == 1 {
		return arr[0], nil
	}
	return nil, e
}
