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
```go
package main

import (
	"log"

	"github.com/marsacotan/go-zabbix7/api"
)

func main() {
	// Supports environment variables; you need to set an available token.
	// api.WithLocalEnvToken("ZABBIX7_API_TOKEN")

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


## License
This project is licensed under the MIT License - see the LICENSE file for details.

## Contributing
Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## Contact
For questions or support, please open an issue on GitHub.