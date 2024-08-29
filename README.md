# go-zabbix7

## Introduction
go-zabbix7 is a Go client library for interacting with the Zabbix 7.0 API. It provides a convenient way to interact with the API using Go. Currently, this library is in the development stage, with ongoing updates and new features being added.

## Installation
To install the library, use the following command:

```sh
go get github.com/marsacotan/go-zabbix7/api
```

## Note
This library does not support Windows.

If you encounter issues while using this library on Windows, please consider using a Unix-based operating system such as Linux or macOS for development and deployment.

## Usage

### 1. Example of obtaining a token
```go
package main

import (
	"log"

	"github.com/marsacotan/go-zabbix7/api"
)

func main() {
	// Supports skipping TLS verification; you can use HTTP or HTTPS.
	// config := api.ConfigConn("https://192.168.88.77/api_jsonrpc.php").WithLoginCred("Admin", "zabbix").WithSkipTlsVerify(true)

	// If TLS verification is required, you need to provide valid certificates.
	config := api.ConfigConn("https://192.168.88.77/api_jsonrpc.php").WithLoginCred("Admin", "zabbix").WithCertPath("./zabbix.crt")

	client := api.NewClient(config)

	resp, err := client.User.GetToken()

	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}
	log.Printf("get token resp: %v", resp)
	log.Printf("token: %v", resp.Result)
}
```
### 2. Three ways to obtain and use a token, using host creation as an example.

#### You can obtain the token using the account credentials, then set the token like this: client.Config.AuthToken = token.
```go
package main

import (
	"fmt"
	"testing"

	"github.com/marsacotan/go-zabbix7/api"
)

func main() {
	config := api.ConfigConn("https://192.168.88.77/api_jsonrpc.php").WithSkipTlsVerify(true).WithLoginCred("Admin", "zabbix")
	client := api.NewClient(config)
	res, err := client.User.GetToken()
	if err != nil {
		fmt.Println(err)
	}
	token := res.Result
	client.Config.AuthToken = token
	result, err := client.Host.HostCreate("testhost", "2")
	if err != nil {
		fmt.Println(err)
	}
}
```

#### You can obtain the pre-set token from the environment variable and use the WithExistingToken method to utilize this token.
```go
package main

import (
	"fmt"
	"testing"

	"github.com/marsacotan/go-zabbix7/api"
)

func main() {
	token := api.WithEnvToken("ZABBIX7_API_TOKEN")
	config := api.ConfigConn("https://192.168.88.77/api_jsonrpc.php").WithSkipTlsVerify(true).WithExistingToken(token).WithExistingToken(token)
	client := api.NewClient(config)
	result, err := client.Host.HostCreate("testhost", "2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result.HostIDs)
}
```

#### You can also directly set the existing token to client.Config.AuthToken.
```go
package main

import (
	"fmt"
	"testing"

	"github.com/marsacotan/go-zabbix7/api"
)

func main() {
	config := api.ConfigConn("https://192.168.88.77/api_jsonrpc.php").WithSkipTlsVerify(true)
	client := api.NewClient(config)
	client.Config.AuthToken = "136405daf4b2f6ce3d0e8fa08dd765ad"
	result, err := client.Host.HostCreate("testhost", "2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result.HostIDs)
}
```


## License
This project is licensed under the MIT License - see the LICENSE file for details.

## Contributing
Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## Contact
For questions or support, please open an issue on GitHub.