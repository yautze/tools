package mongo

import (
	"context"
	"time"

	"github.com/davecgh/go-spew/spew"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var m *mongo.Client

// M - Ping()
func M() (*mongo.Client, error) {
	if err := m.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}
	return m, nil
}

// Close - disconnect DB connection
func Close() {
	m.Disconnect(context.Background())
}

// Con -
func Con(c DBConfig) error {
	if err := con(c); err != nil {
		return err
	}

	if err := m.Ping(context.TODO(), readpref.Primary()); err != nil {
		spew.Dump(err.Error())
	}

	return nil
}

// con - connect MongoDB
func con(c DBConfig) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connect options set
	opts := options.Client()

	opts.ApplyURI(uri(c))

	// set db auth
	if c.User != "" && c.Password != "" {
		opts.SetAuth(options.Credential{
			Username:   c.User,
			Password:   c.Password,
			AuthSource: c.Database, // view which database auth
		})
	}

	// number of connections allowed in the driver's connection pool to each server
	if c.PoolSize > 0 {
		opts.SetMaxPoolSize(c.PoolSize)
	} else {
		opts.SetMaxPoolSize(10)
	}

	opts.SetMinPoolSize(5)
	opts.SetMaxConnIdleTime(time.Second * 300) // maximum amount of time that a connection will remain idle

	// Connect to MongoDB
	client, err := mongo.NewClient(opts)
	if err != nil {
		return err
	}

	if err := client.Connect(ctx); err != nil {
		return err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	m = client

	return nil
}

// mongodb://[username:password@]host1[:port1][,...hostN[:portN]][/[database][?options]]
func uri(c DBConfig) string {
	s := "mongodb://"

	if c.User != "" && c.Password != "" {
		s += c.User + ":" + c.Password + "@"
	}

	s += c.Host

	return s
}
