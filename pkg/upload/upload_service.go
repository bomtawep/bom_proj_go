package upload

import (
	"bom_proj_go/pkg/common/configs"
	"bom_proj_go/pkg/common/database"
	"bom_proj_go/pkg/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"mime/multipart"
	"os"
)

var upload models.File

func insertFile(context *fiber.Ctx) (*mongo.InsertOneResult, error) {
	file, _ := context.FormFile("fileUpload")
	uploadType := context.FormValue("uploadType")
	userId := context.FormValue("userId")
	uploadCollection := database.GetCollection("files")
	ctx, cancel := configs.CtxWithTimout()
	defer cancel()

	mapPath := map[string]string{
		"user":    "/uploads/user",
		"product": "/uploads/product",
	}

	path, _ := mapPath[uploadType]

	fmt.Println("Path: ", path)

	err := os.MkdirAll("."+path, os.ModePerm)
	if err != nil {
		return nil, err
	}

	// Get Buffer from file
	buffer, err := file.Open()
	fmt.Println("File: ", buffer)
	if err != nil {
		panic("Error opening file")
	}
	fmt.Println("Buffer: ", buffer)
	defer func(buffer multipart.File) {
		err := buffer.Close()
		if err != nil {
			panic(err)
		}
	}(buffer)
	filePath := fmt.Sprintf(path+"/%s", file.Filename)
	fmt.Println("File Path: ", filePath)
	if err := context.SaveFile(file, "."+filePath); err != nil {
		return nil, err
	}
	// Assign NewObjectID to user
	upload.Id = primitive.NewObjectID()
	upload.FilePath = filePath
	upload.FileName = file.Filename
	upload.FileSize = file.Size
	upload.FileType = file.Header.Get("Content-Type")
	upload.UserId = userId

	// Insert user to database
	result, insertError := uploadCollection.InsertOne(ctx, &upload)
	if insertError != nil {
		panic("Error inserting file")
	}
	return result, insertError
}

func getFiles() ([]models.File, error) {
	uploadCollection := database.GetCollection("files")
	ctx, cancel := configs.CtxWithTimout()
	defer cancel()

	cursor, err := uploadCollection.Find(ctx, primitive.D{})
	if err != nil {
		panic("Error getting files")
	}
	defer cursor.Close(ctx)

	var files []models.File
	if err = cursor.All(ctx, &files); err != nil {
		panic("Error getting files")
	}
	return files, nil
}
