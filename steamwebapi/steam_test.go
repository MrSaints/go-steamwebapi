package steamwebapi

var (
	client *Client
)

func setup() {
	if client == nil {
		client = NewClient("")
	}
}
