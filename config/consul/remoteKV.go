package consul

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
)

// ConsulKV -  拆出consul Key/Value
func ConsulKV(host, key string, ftype string) (*viper.Viper, error) {
	v := viper.New()

	c := &api.Config{
		Address: host,
		Scheme:  "http",
	}

	// connect consul server
	client, err := api.NewClient(c)

	if err != nil {
		fmt.Println("connect" + host + "failed")
		return nil, err
	}

	// Get a handle to the KV API
	kv := client.KV()

	pair, _, err := kv.Get(key, nil)

	if err != nil {
		fmt.Println("KV Data Not Found")
		return nil, err
	}

	// viper type set
	switch ftype {
	case "yaml":
		v.SetConfigType("yaml")
	case "json":
		v.SetConfigType("json")
	}

	err = v.ReadConfig(bytes.NewReader(pair.Value))

	if err != nil {
		fmt.Println("read consul config failed with key :" + key)
		return nil, err
	}
	return v, nil
}
