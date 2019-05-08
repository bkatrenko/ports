package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/bkatrenko/ports/app/portdomain/domain/model"
	"github.com/bkatrenko/ports/proto"
	wrapper "github.com/pkg/errors"
	"github.com/prometheus/common/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	database   = "ports_info"
	collection = "port"
)

// MongoStorage using to connect to Domain Service to store and retrieve ports info
type MongoStorage struct {
	client       *mongo.Client
	collection   *mongo.Collection
	mongoAddr    string
	mongoTimeout time.Duration
}

// NewMongoStorage is just constructor for MongoStorage instance
func NewMongoStorage(mongoAddr string, mongoTimeout time.Duration) (*MongoStorage, error) {
	log.Infof("start to connect to mongo on addr: %s", mongoAddr)

	if mongoAddr == "" {
		return nil, errors.New("mongo address is empty")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoAddr))
	if err != nil {
		return nil, wrapper.Wrap(err, "error while create mongo client")
	}

	ctx, cancel := context.WithTimeout(context.Background(), mongoTimeout)
	defer cancel()
	if err := client.Connect(ctx); err != nil {
		return nil, wrapper.Wrap(err, "error while connect to mongo")
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, wrapper.Wrap(err, "error while ping mongo")
	}

	collection := client.Database(database).Collection(collection)

	log.Info("successfully connected to mongo")
	return &MongoStorage{
		client:       client,
		collection:   collection,
		mongoAddr:    mongoAddr,
		mongoTimeout: mongoTimeout,
	}, nil
}

// GetPorts will retrieve ports from mongo storage with incoming offset and limit
func (ms *MongoStorage) GetPorts(offset, limit int64) (map[string]*proto.Port, error) {
	if offset < 0 {
		return nil, errors.New("bad offset param: < 0")
	}

	if limit <= 0 {
		return nil, errors.New("bad limit param: <= 0")
	}

	ctx, cancel := context.WithTimeout(context.Background(), ms.mongoTimeout)
	defer cancel()

	cursor, err := ms.collection.Find(ctx, bson.M{}, []*options.FindOptions{&options.FindOptions{Limit: &limit, Skip: &offset}}...)
	if err != nil {
		return nil, wrapper.Wrap(err, "error while find ports")
	}

	defer cursor.Close(ctx)

	res := map[string]*proto.Port{}
	for cursor.Next(ctx) {
		var decodeRes model.MongoDoc

		err := cursor.Decode(&decodeRes)
		if err != nil {
			return nil, wrapper.Wrap(err, "error while decode doc")
		}

		fmt.Println(decodeRes)
		res[decodeRes.PortID] = decodeRes.Doc
	}

	if err := cursor.Err(); err != nil {
		return nil, wrapper.Wrap(err, "mongo curson error")
	}

	return res, nil
}

// PutPorts do Upsert operation to ports mongo bucket
func (ms *MongoStorage) PutPorts(ports map[string]*proto.Port) error {
	ctx, cancel := context.WithTimeout(context.Background(), ms.mongoTimeout)
	defer cancel()

	for k, v := range ports {
		upsert := true
		_, err := ms.collection.UpdateOne(ctx, bson.M{"_port_id": k}, bson.D{{"$set", bson.M{"port_doc": v}}}, &options.UpdateOptions{Upsert: &upsert})
		if err != nil {
			return wrapper.Wrap(err, "error while update")
		}
	}

	return nil
}
