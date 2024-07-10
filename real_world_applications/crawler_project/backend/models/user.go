package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

type User struct {
	ID       string `bson:"_id,omitempty" json:"id"`
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
	Email    string `bson:"email" json:"email"`
}

func (u *User) Insert() error {
	collection := client.Database("users").Collection("users")
	_, err := collection.InsertOne(context.TODO(), User{
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
	})

	if err != nil {
		return fmt.Errorf("error while inserting user: %v", err)
	}
	return nil
}

func (u *User) GetUserByEmail(email string) (*User, error) {
	collection := client.Database("users").Collection("users")
	filter := bson.M{"email": email}
	var user_detail User
	err := collection.FindOne(context.TODO(), filter).Decode(&user_detail)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("user not found")
		}
		fmt.Println(err)
		return nil, fmt.Errorf("error while getting user: %v", err)
	}
	return &user_detail, nil
}

func (u *User) GetUserByID(ID string) (*User, error) {
	collection := client.Database("users").Collection("users")
	userId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, fmt.Errorf("error while converting id: %v", err)
	}
	filter := bson.M{"_id": userId}
	var user User
	err = collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("error while getting user: %v", err)
	}
	return &user, nil
}

func (u *User) Update(ID string) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("users").Collection("users")
	userid, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, fmt.Errorf("error while converting id: %v", err)
	}
	filter := bson.M{"_id": userid}
	update := bson.M{"$set": bson.M{
		"username": u.Username,
	}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("error while updating user: %v", err)
	}

	return result, nil
}

func (u *User) Delete(ID string) error {
	collection := client.Database("users").Collection("users")
	userId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return fmt.Errorf("error while converting id: %v", err)
	}
	filter := bson.M{"id": userId}
	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("user not found")
		}
		return fmt.Errorf("error while deleting user: %v", err)
	}
	return nil
}
