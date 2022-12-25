package courses

import (
	"context"
	"fmt"

	commons "github.com/sidharthchoudhary/lmsAuth/Commons"
	"github.com/sidharthchoudhary/lmsAuth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// function to handle the file and uplooad to the server
func CreateStringContent(course models.CourseContent, courseID string) commons.Response {
	//converting the course ID
	id, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return commons.Response{
			Status:  0,
			Message: "Invalid Course ID",
		}
	}
	fmt.Println("Id is ", id)
	fmt.Println(course)

	//adding the course ID to the course
	course.Id = id
	filter := bson.M{"_id": id}
	//inserting the course to the database
	updateFilter := bson.M{"$push": bson.M{"courseContent": course}}
	inserted, err := CollectionMongo.UpdateMany(context.TODO(), filter, updateFilter)
	if err != nil {
		return commons.Response{
			Status:  0,
			Message: "Error Adding Course Content",
		}
	}
	//returning the response
	return commons.Response{
		Status:  1,
		Message: "Course Content Added",
		Data:    inserted,
	}
}

// creating the course with the file content
func CreateFileContent() commons.Response {
	return commons.Response{
		Status:  1,
		Message: "Function Called",
	}
}

// UploadFile uploads an object
// func UploadFile(file multipart.File, object string) error {
// 	ctx := context.Background()

// 	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
// 	defer cancel()

// 	// Upload an object with storage.Writer.
// 	wc := c.cl.Bucket(c.bucketName).Object(c.uploadPath + object).NewWriter(ctx)
// 	if _, err := io.Copy(wc, file); err != nil {
// 		return fmt.Errorf("io.Copy: %v", err)
// 	}
// 	if err := wc.Close(); err != nil {
// 		return fmt.Errorf("Writer.Close: %v", err)
// 	}

// 	return nil
// }
