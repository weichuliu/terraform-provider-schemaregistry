package schemaregistry

import (
	"fmt"
	"log"

	"github.com/weichuliu/schema-registry"
)

// Config defines the configuration options for the Schema Registry client
type Config struct {
	// Url for Schema Registry instance
	URL string

	Username string
	Password string
}

// Client returns a new Schema Registry client
func (c *Config) Client() (*schemaregistry.Client, error) {
	// Validate that the Schema Registry url is set
	if c.URL == "" {
		return nil, fmt.Errorf("Please set `url` of Schema Registry service")
	}

	if c.Username == "" {
		log.Printf("[INFO] Schema registry client initialized")

		// return client, nil
		client, _ := schemaregistry.NewClient(c.URL)

		return client, nil
	} else {
		log.Printf("[INFO] Schema registry client initialized with auth")

		// return client, nil
		client, _ := schemaregistry.NewClient(c.URL)

		client.SetBasicAuth(c.Username, c.Password)

		return client, nil
	}

}
