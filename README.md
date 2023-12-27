# Bots-futures (Go)
API for Bots Futures (example)

Documentation
https://api2.cryptorg.net/documentation/api/v2

## Usage
In your go project, import the package as usual.

You must have an API key and secret to use this package (See [documentation](https://api2.cryptorg.net/documentation/api/v2)).

Place them to .env file.

### Example
```go

import (
	client "github.com/cryptorg-io/bots-futures-go"
	"fmt"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		return
	}

	c := client.New()
	res, err := c.GetBotDetails(9000000000)
	
	fmt.Printf("Bot: %v\n", res)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
```