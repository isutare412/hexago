package repo

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/isutare412/hexago/gateway/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	cli *mongo.Client
	db  *mongo.Database
}

func NewMongoDB(
	ctx context.Context,
	cfg *config.MongoDBConfig,
) (*MongoDB, error) {
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

	return &MongoDB{
		cli: cli,
		db:  cli.Database(cfg.Database),
	}, nil
}

func (mdb *MongoDB) StartSession(
	ctx context.Context,
) (mongo.SessionContext, error) {
	sess, err := mdb.cli.StartSession()
	if err != nil {
		return nil, fmt.Errorf("start session: %w", err)
	}
	if err := sess.StartTransaction(); err != nil {
		return nil, fmt.Errorf("start trasaction: %w", err)
	}

	return mongo.NewSessionContext(ctx, sess), nil
}

func (mdb *MongoDB) Shutdown(ctx context.Context) error {
	return mdb.cli.Disconnect(ctx)
}

func (mdb *MongoDB) Migrate(ctx context.Context) error {
	_, err := mdb.db.Collection(collUser).Indexes().
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
	return nil
}

func isErrDupKey(err error) bool {
	if err == nil {
		return false
	}

	var exc mongo.WriteException
	if !errors.As(err, &exc) {
		return false
	}
	for _, err := range exc.WriteErrors {
		if err.Code == errCodeDuplicateKey {
			return true
		}
	}
	return false
}

func isErrNotFound(err error) bool {
	return errors.Is(err, mongo.ErrNoDocuments)
}
