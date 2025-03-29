package config

type RedisConfig struct {
	Addr     string
	Password string
	Db       int
}

func NewRedisConfig(address string, password string, db int) *RedisConfig {
	return &RedisConfig{
		Addr:     address,
		Password: password,
		Db:       db,
	}
}
