package mongo

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/isutare412/hexago/payment/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	cli *mongo.Client
	db  *mongo.Database
}

func NewRepository(
	ctx context.Context,
	cfg *config.MongoDBConfig,
) (*Repository, error) {
	heartbeat := time.Duration(cfg.HeartbeatInterval)
	addrs := strings.Join(cfg.Addrs, ",")
	uri := fmt.Sprintf("mongodb://%s", addrs)

	cli, err := mongo.Connect(ctx, options.Client().
		ApplyURI(uri).
		SetAuth(options.Credential{
			AuthSource: cfg.AuthSource,
			Username:   cfg.Username,
			Password:   cfg.Password,
		}).
		SetHeartbeatInterval(heartbeat*time.Millisecond).
		SetMaxPoolSize(uint64(cfg.MaxConnectionPool)),
	)
	if err != nil {
		return nil, fmt.Errorf("connecting mongodb: %w", err)
	}

	if err := cli.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("ping MongoDB: %w", err)
	}

	return &Repository{
		cli: cli,
		db:  cli.Database(cfg.Database),
	}, nil
}

func (r *Repository) Shutdown(ctx context.Context) error {
	return r.cli.Disconnect(ctx)
}

func (r *Repository) Migrate(ctx context.Context) error {
	_, err := r.db.Collection(collUser).Indexes().
		CreateMany(ctx, []mongo.IndexModel{
			{
				Keys:    bson.M{"id": 1},
				Options: options.Index().SetUnique(true),
			},
			{
				Keys:    bson.M{"email": 1},
				Options: options.Index().SetUnique(true),
			},
		})
	if err != nil {
		return fmt.Errorf("creating user indexes: %w", err)
	}

	_, err = r.db.Collection(collDonationHistory).Indexes().
		CreateOne(ctx, mongo.IndexModel{
			Keys: bson.M{"timestamp": 1},
		})
	if err != nil {
		return fmt.Errorf("creating donation history indexes: %w", err)
	}
	return nil
}
