package service

import (
	"github.com/ariwanss/CvBackendGo/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func verifyPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func CreateUser(user *entity.User) (*entity.User, error) {
	hashedPassword, err := hashPassword(user.Password)

	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword
	res, err := userCol.InsertOne(c, user)

	if err != nil {
		return nil, err
	}

	insertedId := res.InsertedID.(primitive.ObjectID)
	user.ID = insertedId

	return user, nil
}

func GetUserByUsername(username string) (*entity.User, error) {
	res := userCol.FindOne(c, bson.M{"username": username})
	var user entity.User
	err := res.Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func LoginUser(username, password string) (*entity.User, error) {
	user, err := GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	err = verifyPassword(user.Password, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(userId primitive.ObjectID, update *entity.User) (*entity.User, error) {
	if update.Password != "" {
		hashedPassword, err := hashPassword(update.Password)
		if err != nil {
			return nil, err
		}
		update.Password = hashedPassword
	}
	updateDoc := bson.M{
		"$set": update,
	}
	res := userCol.FindOneAndUpdate(c, bson.M{"_id": userId}, updateDoc, options.FindOneAndUpdate().SetReturnDocument(options.After))
	var updatedUser entity.User
	err := res.Decode(&updatedUser)
	if err != nil {
		return nil, err
	}
	return &updatedUser, nil
}

func DeleteUser(userId primitive.ObjectID) (int, error) {
	res, err := userCol.DeleteOne(c, bson.M{"_id": userId})
	return int(res.DeletedCount), err
}

func GetAllUser() (*[]entity.User, error) {
	res, err := userCol.Find(c, bson.D{{}})
	if err != nil {
		return nil, err
	}
	var users []entity.User
	err = res.All(c, &users)
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func DropUserCol() {
	err := userCol.Drop(c)
	if err != nil {
		panic("failed to drop user collection")
	}
}
