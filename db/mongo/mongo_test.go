package mongo

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_OneCon(t *testing.T) {
	config := DBConfig{
		Host:     "localhost:27017",
		User:     "",
		Password: "",
		Database: "TEST",
		AppName:  "test",
	}

	err := Con(config)
	assert.NoError(t, err)

	client, err := M()
	assert.NoError(t, err)

	c := client.Database("TEST").Collection("user")

	cur, err := c.Find(context.Background(), bson.M{})
	assert.NoError(t, err)

	for cur.Next(context.Background()) {
		res := bson.M{}
		cur.Decode(&res)
		spew.Dump(res)
	}
}

// Test_Con - 叢集連線
func Test_Con(t *testing.T) {
	config := DBConfig{
		Host: "localhost:27017,localhost:27018,localhost:27019",
		//User:     "root",
		//Password: "9527",
		Database:   "csofans",
		AppName:    "user",
		ReplicaSet: "rs0",
	}

	err := Con(config)
	assert.NoError(t, err)

	client, err := M()
	assert.NoError(t, err)

	c := client.Database("csofans").Collection("user")

	cur, err := c.Find(context.Background(), bson.M{})
	assert.NoError(t, err)

	for cur.Next(context.Background()) {
		res := bson.M{}
		cur.Decode(&res)
		spew.Dump(res)
	}
}
