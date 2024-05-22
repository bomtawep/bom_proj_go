package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type File struct {
	Id primitive.ObjectID `json:"id,omitempty"`
	// Define the file as a buffer type
	FilePath string `json:"filepath,omitempty"`
	FileName string `json:"filename,omitempty"`
	FileSize int64  `json:"filesize,omitempty"`
	FileType string `json:"filetype,omitempty"`
	UserId   string `json:"userid,omitempty"`
}
