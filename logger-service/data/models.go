package data

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func New(mongo *mongo.Client) Models {
	client = mongo
	ret := Models{
		LogEntry: LogEntry{},
	}
	return ret
}

type Models struct {
	LogEntry LogEntry
}

type LogEntry struct {
	ID     string    `bson:"_id,omitempty" json:"id,omitempty"`
	UserId string    `bson:"user_id" json:"user_id"`
	Action string    `bson:"action" json:"action"`
	Time   time.Time `bson:"time" json:"time"`
}

func (l LogEntry) Insert(entry LogEntry) error {
	collection := client.Database("logs").Collection("entries")

	entryForLog := LogEntry{
		UserId: entry.UserId,
		Action: entry.Action,
		Time:   time.Now(),
	}

	_, err := collection.InsertOne(context.Background(), entry)
	if err != nil {
		log.Print("error inserting log entry ", err)
		return err
	}
}

func (l LogEntry) All() ([]LogEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Database("logs").Collection("entries")

	ops := options.Find()
	ops.SetSort(bson.D{{"time", -1}})

	cursor, err := collection.Find(context.TODO(), bson.D{}, ops)
	if err != nil {
		log.Print("error getting log entries ", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var ret []LogEntry

	for cursor.Next(ctx) {
		var entry LogEntry
		err := cursor.Decode(&entry)
		if err != nil {
			log.Print("error decoding log entry ", entry, err)
			return nil, err
		}

		ret = append(ret, &entry)
	}

	return ret, nil
}

func (l *LogEntry) FetONe(id string) (*LogEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Database("logs").Collection("entries")

	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Print("error getting primitive id from hex ", err)
		return nil, err
	}

	var entry LogEntry
	err := collection.FindOne(ctx, bson.M{{"_id", docID}}).Decode(&entry)
	if err != nil {
		log.Print("error getting log entry ", err)
		return nil, err
	}

	return &entry, nil
}

func (l *LongEntry) DropCollection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Database("logs").Collection("entries")

	err := collection.Drop(ctx)
	if err != nil {
		log.Print("error dropping collection ", err)
		return err
	}

	return nil
}
