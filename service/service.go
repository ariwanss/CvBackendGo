package service

import (
	"context"

	"github.com/ariwanss/CvBackendGo/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCol *mongo.Collection
var cvItemCol *mongo.Collection
var c context.Context

func StartService() {
	userCol = repository.Database.Collection("users")
	cvItemCol = repository.Database.Collection("cvItems")
	c = context.TODO()
}