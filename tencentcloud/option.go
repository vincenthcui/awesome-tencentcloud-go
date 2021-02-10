package tencentcloud

type Option func(*Client)

func WithSecret(id, key string) Option {
	return func(c *Client) {
		c.secretID = id
		c.secretKey = key
	}
}
