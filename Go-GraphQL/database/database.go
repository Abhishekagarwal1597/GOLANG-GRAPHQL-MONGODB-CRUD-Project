package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Abhishekagarwal1597/Go-GraphQL-Project/graph/model"
)

// var connectionString string = "mongodb://localhost:27017"
var connectionString string = "mongodb://localhost:27017"

type DB struct {
	client *mongo.Client
}

func Connect() *DB {

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, options.Client().ReadPreference)
	if err != nil {
		log.Fatal(err)
	}
	return &DB{client: client}
}

func (db *DB) GetJobs() []*model.JobListing {
	// dbInstance := Connect()
	collection := db.client.Database("graphql-jobs").Collection("jobListing")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var jobListings []*model.JobListing
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(context.TODO(), &jobListings); err != nil {
		panic(err)
	}

	// for cursor.Next(ctx) {
	// 	var jobListing *model.JobListing
	// 	err := cursor.Decode(&jobListing)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	jobListings = append(jobListings, jobListing)
	// }
	return jobListings
}

func (db *DB) GetJob(id string) *model.JobListing {
	jobCollec := db.client.Database("graphql-jobs").Collection("jobListing")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var jobListing model.JobListing
	err := jobCollec.FindOne(ctx, filter).Decode(&jobListing)
	if err != nil {
		log.Fatal(err)
	}
	return &jobListing

	// dbInstance := Connect()
	// collection := dbInstance.client.Database("graphql-jobs").Collection("jobListing")

	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()

	// var jobListing model.JobListing
	// _id, _ := primitive.ObjectIDFromHex(*jobID)
	// fmt.Println("id:", _id)
	// err := collection.FindOne(ctx, bson.M{"id": _id}).Decode(&jobListing)
	// if err != nil {
	// 	fmt.Println("err in finding collection ....")
	// 	log.Fatal(err)
	// }
	// // cur.Decode(&jobListing)

	// return &jobListing
}

func (db *DB) CreateJobListing(jobInfo model.CreateJobListingInput) *model.JobListing {
	// dbInstance := Connect()
	collection := db.client.Database("graphql-jobs").Collection("jobListing")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	inserg, err := collection.InsertOne(ctx, bson.M{"title": jobInfo.Title, "description": jobInfo.Description, "url": jobInfo.URL, "company": jobInfo.Company})
	if err != nil {
		log.Fatal(err)
	}
	insertedID := inserg.InsertedID.(primitive.ObjectID).Hex()

	jobListingResp := model.JobListing{ID: insertedID, Title: jobInfo.Title, Description: jobInfo.Description, Company: jobInfo.Company, URL: jobInfo.URL}
	return &jobListingResp

	// log.Println("Job Listing Created")

}

func (db *DB) DeleteJobListing(jobId string) *model.DeleteJobResponse {
	// dbInstance := Connect()
	collection := db.client.Database("graphql-jobs").Collection("jobListing")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(jobId)
	fmt.Println("ID:", _id)
	_, err := collection.DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		fmt.Println("err in deleting doc...")
		log.Fatal(err)
	}

	DeletejobResp := model.DeleteJobResponse{DeletedJobID: jobId}
	return &DeletejobResp

	// log.Println("Job Listing Created")

}

func (db *DB) UpdateJobListing(jobId string, jobInfo model.UpdateJobListingInput) *model.JobListing {
	fmt.Println("jobID:", jobId)
	fmt.Println("jobInfo:", jobInfo)
	jobCollec := db.client.Database("graphql-jobs").Collection("jobListing")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	updateJobInfo := bson.M{}

	if jobInfo.Title != nil {
		updateJobInfo["title"] = jobInfo.Title
	}
	// if jobInfo.Description != nil {
	// 	updateJobInfo["description"] = jobInfo.Description
	// }
	// if jobInfo.URL != nil {
	// 	updateJobInfo["url"] = jobInfo.URL
	// }

	_id, _ := primitive.ObjectIDFromHex(jobId)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updateJobInfo}

	results := jobCollec.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var jobListing model.JobListing

	if err := results.Decode(&jobListing); err != nil {
		log.Fatal(err)
	}

	return &jobListing
}
