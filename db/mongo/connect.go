package mongo

import (
	"context"
	"crypto/tls"
	"errors"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var m *mongo.Client


// M - return m
func M() (*mongo.Client, error) {
	if m == nil {
		return nil, errors.New("no connection")
	}

    // test connection is OK
	if err := m.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}

	return m, nil
}

// Close -
func Close() {
	m.Disconnect(context.Background())
}

// Con - 
func Con(c DBConfig) error {
	if err := con(c); err != nil {
		return err
	}

    // test connection is OK
	if err := m.Ping(context.TODO(), readpref.Primary()); err != nil {
		return err
	}

	return nil
}

func con(c DBConfig) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connect options set
	opts := options.Client()
	opts.ApplyURI(uri(c)).SetAppName(c.AppName)

	// set db auth
	if c.User != "" && c.Password != "" {
		opts.SetAuth(options.Credential{
			Username:   c.User,
			Password:   c.Password,
			AuthSource: c.Database, // view which database auth
		})
	}

	//  連線到mongo叢集
	if c.ReplicaSet != "" {
		opts.ReplicaSet = &c.ReplicaSet
	}

	// set TLS/SSL (做成加密連線)
	if c.SSL {
		opts.SetTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		})
	}

	// 設定連接池個數
	if c.PoolSize > 0 {
		opts.SetMaxPoolSize(c.PoolSize)
	} else {
		opts.SetMaxPoolSize(10)
	}

	opts.SetMinPoolSize(5)                     // 最少
	opts.SetMaxConnIdleTime(time.Second * 300) // 最久連接時間
	opts.SetHeartbeatInterval(10 * time.Second)

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
	host := strings.Split(c.Host, ",")

	for i, v := range host {
		if c.User != "" && c.Password != "" {
			s += c.User + ":" + c.Password + "@"
		}
		s += v
		if i != len(host)-1 {
			s += ","
		}
	}

	return s
}
