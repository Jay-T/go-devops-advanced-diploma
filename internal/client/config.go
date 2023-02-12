package client

type Config struct {
	Host      string
	Masterkey string
}

func NewConfig() (*Config, error) {
	c := &Config{}
	if err := runtimeViper.Unmarshal(c); err != nil {
		return nil, err
	}
	return c, nil
}
