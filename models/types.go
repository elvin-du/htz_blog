package models

import (
	"time"
)

type Nav struct {
	Id   int
	Name string
}

type Article struct {
	Id              int
	Title           string
	Content         string
	ContentAbstract string
	ReadCount       uint
	NavName         string
	CreatedAt       time.Time
}

type File struct {
	Id        int
	Mime      string
	Size      int
	CreatedAt time.Time
}
