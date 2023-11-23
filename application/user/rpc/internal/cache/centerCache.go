package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type CenterCache struct {
	rdb *redis.Client
}

func NewCenterCache(rdb *redis.Client) *CenterCache {
	return &CenterCache{
		rdb: rdb,
	}
}

func (r *CenterCache) Insert(ctx context.Context, token string) error {
	return r.rdb.Set(ctx, token, "", time.Hour*168).Err()
}

func (r *CenterCache) InsertUserCode(ctx context.Context, phoneNum, code string) error {
	return r.rdb.Set(ctx, phoneNum+"USO", code, time.Hour*168).Err()
}

func (r *CenterCache) GetUserCodeAgain(ctx context.Context, phoneNum string) string {
	return r.rdb.Get(ctx, phoneNum+"USO").String()
}
