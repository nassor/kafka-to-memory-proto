package device

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/nassor/kafka-compact-pipeline/api"
)

var t = true

// MongoStore implements a MongoDB store
type MongoStore struct {
	mc *mongo.Collection
}

// NewMongoStore creates a new instance of MongoStore
func NewMongoStore(ctx context.Context, mc *mongo.Collection) (*MongoStore, error) {
	ms := &MongoStore{mc: mc}
	if err := ms.updateIndexes(ctx, mc); err != nil {
		return nil, err
	}
	return ms, nil
}

func (st *MongoStore) updateIndexes(ctx context.Context, mc *mongo.Collection) error {
	o := &options.IndexOptions{
		Unique: &t,
	}
	_, err := mc.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bsonx.Doc{{Key: "device_id", Value: bsonx.Int32(-1)}},
		Options: o,
	})
	return err
}

// Upsert insert or update store info data
func (st *MongoStore) Upsert(ctx context.Context, d *api.Device) error {
	_, err := st.mc.UpdateOne(ctx,
		bson.M{"device_id": d.Id},
		bson.M{"$set": bson.M{
			"device_id":       d.Id,
			"organization_id": d.OrganizationId,
			"tz":              d.Tz,
			"installed_at":    d.InstalledAt,
			"updated_at":      d.UpdatedAt,
		}},
		&options.UpdateOptions{Upsert: &t})
	return err
}
