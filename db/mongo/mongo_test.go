package mongo

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_Con(t *testing.T) {
	db := DBConfig{
		Host:     "localhost:27017",
		User:     "",
		Password: "",
		Database: "test",
	}

	err := Con(db)
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
