package tencentcloud

type Option func(*client)

func WithSecret(id, key string) Option {
	return func(c *client) {
		c.secretID = id
		c.secretKey = key
	}
}
