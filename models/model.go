package models

import (
	"time"
)

type Model struct {
	Created time.Time `bson:"created"`
	Updated time.Time `bson:"updated"`
}
