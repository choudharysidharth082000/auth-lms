package courses

import (
	"context"
	"fmt"
	"log"

	commons "github.com/sidharthchoudhary/lmsAuth/Commons"
	"github.com/sidharthchoudhary/lmsAuth/contollers/Auth"
	"github.com/sidharthchoudhary/lmsAuth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddWishListController(courseId string, userId string) commons.Response {
	//converting the types
	course, err := primitive.ObjectIDFromHex(courseId)
	if err != nil {
		return commons.Response{
			Status:  0,
			Message: "Invalid Course Id",
		}
	}
	user, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		fmt.Println(err)
		return commons.Response{
			Status:  0,
			Message: "Invalid User Id",
		}
	}
	//adding the course to the wishlist
	filter := bson.M{"_id": user}
	//pushing update to wishlist
	update := bson.M{"$push": bson.M{"wishlist": course}}
	//updating the wishlist
	inserted, err := Auth.CollectionMongo.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
		return commons.Response{
			Status:  0,
			Message: "Error Adding Course to Wishlist",
		}
	}
	fmt.Println(inserted)
	return commons.Response{
		Status:  1,
		Message: "Course Added to Wishlist",
	}

}

// add the product to te cart
func AddProductToCartController(productId string, userId string) commons.Response {
	//converting the types
	product, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return commons.Response{
			Status:  0,
			Message: "Invalid Product Id",
		}
	}
	user, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return commons.Response{
			Status:  0,
			Message: "Invalid User Id",
		}
	}
	//adding the course to the wishlist
	filter := bson.M{"_id": user}
	//pushing update to wishlist
	update := bson.M{"$push": bson.M{"cart": product}}
	//updating the wishlist
	inserted, err := Auth.CollectionMongo.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
		return commons.Response{
			Status:  0,
			Message: "Error Adding Product to Cart",
		}
	}
	fmt.Println(inserted)
	return commons.Response{
		Status:  1,
		Message: "Product Added to Cart",
	}
}

func GetAllCoursesController() commons.Response {
	var courseVal []models.Course
	//finding all the values from the database
	result, err := CollectionMongo.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	for result.Next(context.TODO()) {
		var singleCourse models.Course
		err := result.Decode(&singleCourse)
		if err != nil {
			log.Fatal(err)
		}
		courseVal = append(courseVal, singleCourse)
	}
	if err := result.Err(); err != nil {
		log.Fatal(err)
	}
	result.Close(context.TODO())
	return commons.Response{
		Status: 1,
		Data:   courseVal,
	}

}

func InsertCourse(val models.Course) commons.Response {
	//inserting the course to the database
	inserted, err := CollectionMongo.InsertOne(context.TODO(), val)
	if err != nil {
		log.Fatal(err)
		return commons.Response{
			Status:  0,
			Message: "Error Adding Course",
		}
	}
	fmt.Println(inserted)
	return commons.Response{
		Status:  1,
		Message: "Course Added",
	}	
}
//add Course
func AddCourse(course models.Course) commons.Response {
	//inserting the course to the database
	inserted, err := CollectionMongo.InsertOne(context.TODO(), course)
	if err != nil {
		log.Fatal(err)
		return commons.Response{
			Status:  0,
			Message: "Error Adding Course",
		}
	}
	fmt.Println(inserted)
	return commons.Response{
		Status:  1,
		Message: "Course Added",
	}
}

//edit a single course
func UpdateCourse(course models.Course, courseID string) commons.Response {
	//converting the types
	id, err := primitive.ObjectIDFromHex(courseID)	
	if err != nil {
		return commons.Response{
			Status:  0,
			Message: "Invalid Course Id",
		}
	}
	//filtering the course
	filter := bson.M{"_id": id}
	//updating the course
	update := bson.M{"$set": bson.M{"name": course.Name, "description": course.Description, "price": course.Price, "image": course.Image}}
	//updating the course
	inserted, err := CollectionMongo.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
		return commons.Response{
			Status:  0,
			Message: "Error Updating Course",
		}
	}
	fmt.Println(inserted)
	return commons.Response{
		Status:  1,
		Message: "Course Updated",
	}
}

//getting single course
func GetSingleCourse(courseId string) commons.Response {
	//converting the types
	id, err := primitive.ObjectIDFromHex(courseId)
	if err != nil {
		return commons.Response{
			Status:  0,
			Message: "Invalid Course Id",
		}
	}
	//filtering the course
	filter := bson.M{"_id": id}
	//getting the course
	var course models.Course
	err = CollectionMongo.FindOne(context.TODO(), filter).Decode(&course)
	if err != nil {
		return commons.Response{
			Status:  0,
			Message: "Error Getting Course",
		}
	}
	return commons.Response{
		Status:  1,
		Message: "Course Retrieved",
		Data:    course,
	}
}

//deleting the course from the database
func DeleteCourseFunction(courseId string) commons.Response {
	//converting the types
	course, err := primitive.ObjectIDFromHex(courseId)
	if err != nil {
		return commons.Response{
			Status:  0,
			Message: "Invalid Course Id",
		}
	}
	//deleting the course from the database
	filter := bson.M{"_id": course}
	//deeleting the course
	deleted, err := CollectionMongo.DeleteOne(context.TODO(), filter);
	if err != nil {
		log.Fatal(err)
		return commons.Response{
			Status:  0,
			Message: "Error Deleting Course",
		}
	}
	fmt.Println(deleted)
	return commons.Response{
		Status:  1,
		Message: "Course Deleted",
		Data:   deleted,
	}
}
//delete all the courses from the database of the particular user id 
func DeleteManyCourses(courseID string) commons.Response{
	//converting the types
	course, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return commons.Response{
			Status:  0,
			Message: "Invalid Course Id",
		}
	}
	//deleting the course from the database
	filter := bson.M{"_id": course}
	//deeleting the course
	deleted, err := CollectionMongo.DeleteMany(context.TODO(), filter);
	if err != nil {
		log.Fatal(err)
		return commons.Response{
			Status:  0,
			Message: "Error Deleting Course",
		}
	}
	fmt.Println(deleted)
	return commons.Response{
		Status:  1,
		Message: "Course Deleted",
		Data:   deleted,
	}
}
//deleting all the records based on created by
func DeleteManyCoursesByUser(createdBy string) commons.Response{
	//deleting the course from the database
	filter := bson.M{"createdBy": createdBy}
	//deeleting the course
	deleted, err := CollectionMongo.DeleteMany(context.TODO(), filter);
	if err != nil {
		log.Fatal(err)
		return commons.Response{
			Status:  0,
			Message: "Error Deleting Course",
		}
	}
	fmt.Println(deleted)
	return commons.Response{
		Status:  1,
		Message: "Course Deleted",
		Data:   deleted,
	}
}