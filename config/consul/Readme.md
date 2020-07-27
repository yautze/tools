# Config



## ConsulKV
```golang test code
package main

import (
	"fmt"

	"github.com/yautz/go-tools/config/consul"
)

func main() {
	v, err := consul.ConsulKV("consulhost" , "key" , "ftype{json/yaml}")
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(v.GetString("consulKey"))
}
```
