package models

import (
	"github.com/zhuharev/bloblog"
)

var (
	bl *bloblog.BlobLog
)

func NewBlobLogContext() (e error) {
	bl, e = bloblog.Open("data/data.bl")
	return
}

func InsertBlob(bytes []byte) (int64, error) {
	return bl.Insert(bytes)
}

func GetBlob(id int64) ([]byte, error) {
	return bl.Get(id)
}
