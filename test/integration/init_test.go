// +build integration

package integration

import "github.com/JPZ13/service/client"

var testClient = makeTestClient()

func makeTestClient() client.Blog {
	return client.New(&client.Config{
		ServiceURI:     "http://localhost:8080",
		TimeoutSeconds: 10,
	})
}
