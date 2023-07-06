package service

import (
	"github.com/ariwanss/CvBackendGo/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertCvItem(cvItem *entity.CvItem) (*entity.CvItem, error) {
	res, err := cvItemCol.InsertOne(c, cvItem)
	if err != nil {
		return nil, err
	}
	insertedId := res.InsertedID.(primitive.ObjectID)
	cvItem.ID = insertedId
	return cvItem, nil
}

func GetCvItemsByUserId(userId primitive.ObjectID) (*[]entity.CvItem, error) {
	res, err := cvItemCol.Find(c, bson.M{"userId": userId})
	if err != nil {
		return nil, err
	}
	var cvItems []entity.CvItem
	err = res.All(c, &cvItems)
	if err != nil {
		return nil, err
	}
	return &cvItems, nil
}

func UpdateCvItem(cvItemId, userId primitive.ObjectID, update *entity.CvItem) (*entity.CvItem, error) {
	updateDoc := bson.M{"$set": update}
	res := cvItemCol.FindOneAndUpdate(c, bson.M{"_id": cvItemId, "userId": userId}, updateDoc, options.FindOneAndUpdate().SetReturnDocument(options.After))
	var updatedCvItem entity.CvItem
	err := res.Decode(&updatedCvItem)
	if err != nil {
		return nil, err
	}
	return &updatedCvItem, nil
}

func DeleteCvItem(cvItemId primitive.ObjectID) (int, error) {
	res, err := cvItemCol.DeleteOne(c, bson.M{"_id": cvItemId})
	return int(res.DeletedCount), err
}

func DeleteAllCvItemsByUserId(userId primitive.ObjectID) error {
	_, err := cvItemCol.DeleteMany(c, bson.M{"usreId": userId})
	return err
}

func DropCvItemCol() {
	err := cvItemCol.Drop(c)
	if err != nil {
		panic("failed to drop cvItem collection")
	}
}

func GetAllCvItems() (*[]entity.CvItem, error) {
	res, err := cvItemCol.Find(c, bson.M{})
	if err != nil {
		return nil, err
	}
	var cvItems []entity.CvItem
	err = res.All(c, &cvItems)
	if err != nil {
		return nil, err
	}
	return &cvItems, nil
}

func GetCvItem(id string) (*entity.CvItem, error) {
	res := cvItemCol.FindOne(c, bson.M{"_id": id})
	var cvItem entity.CvItem
	err := res.Decode(&cvItem)
	if err != nil {
		return nil, err
	}
	return &cvItem, nil
}
