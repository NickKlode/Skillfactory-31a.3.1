package mongodb

import (
	"context"
	"gonews/pkg/storage"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	Client *mongo.Client
	DB     *mongo.Database
}

const (
	databaseName   = "news"
	collectionName = "posts"
)

func New(constr string) (*Storage, error) {
	mongoOpts := options.Client().ApplyURI(constr)
	client, err := mongo.Connect(context.Background(), mongoOpts)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	s := Storage{
		Client: client,
		DB:     client.Database(databaseName),
	}
	return &s, nil
}

func (s *Storage) Posts() ([]storage.Post, error) {
	collection := s.Client.Database(databaseName).Collection(collectionName)
	filter := bson.D{}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	var data []storage.Post
	for cur.Next(context.Background()) {
		var t storage.Post
		err := cur.Decode(&t)
		if err != nil {
			return nil, err
		}
		data = append(data, t)
	}
	return data, cur.Err()
}

func (s *Storage) AddPost(p storage.Post) error {
	collection := s.Client.Database(databaseName).Collection(collectionName)
	_, err := collection.InsertOne(context.Background(), p)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) UpdatePost(p storage.Post) error {
	collection := s.Client.Database(databaseName).Collection(collectionName)
	update := bson.M{
		"$set": p,
	}
	_, err := collection.UpdateOne(context.Background(), bson.M{"id": p.ID}, update)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) DeletePost(p storage.Post) error {
	collection := s.Client.Database(databaseName).Collection(collectionName)
	_, err := collection.DeleteOne(context.Background(), bson.M{"id": p.ID})
	if err != nil {
		return err
	}
	return nil
}
