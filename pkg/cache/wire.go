package cache

import (
	"context"

	"git-devops.opencsg.com/product/community/starhub-server/config"
	"github.com/google/wire"
)

// WireSet provides a wire set for this package.
var WireSet = wire.NewSet(
	ProvideRedisConfig,
	ProvideCache,
)

func ProvideRedisConfig(config *config.Config) RedisConfig {
	return RedisConfig{
		Addr:     config.Redis.Endpoint,
		Username: config.Redis.User,
		Password: config.Redis.Password,
	}

}

func ProvideCache(ctx context.Context, cfg RedisConfig) (*Cache, error) {
	return NewCache(ctx, cfg)
}
