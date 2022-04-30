package repo

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/isutare412/hexago/pkg/config"
	"github.com/isutare412/hexago/pkg/core/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (mdb *MongoDB) Close(ctx context.Context) error {
	return mdb.cli.Disconnect(ctx)
}

func (mdb *MongoDB) InsertStudent(
	ctx context.Context,
	stu *entity.Student,
) (*entity.Student, error) {
	coll := mdb.db.Collection(collStudent)

	stu.Id = primitive.NewObjectID()
	_, err := coll.InsertOne(ctx, stu)
	if err != nil {
		return nil, fmt.Errorf("insert student: %w", err)
	}
	return stu, nil
}

func (mdb *MongoDB) FindStudentById(
	ctx context.Context,
	id primitive.ObjectID,
) (*entity.Student, error) {
	coll := mdb.db.Collection(collStudent)

	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	res := coll.FindOne(ctx, filter)
	if err := res.Err(); err != nil {
		return nil, fmt.Errorf("find student: %w", err)
	}

	stu := entity.Student{}
	if err := res.Decode(&stu); err != nil {
		return nil, fmt.Errorf("decode student: %w", err)
	}
	return &stu, nil
}

func (mdb *MongoDB) DeleteStudentById(
	ctx context.Context,
	id primitive.ObjectID,
) (int64, error) {
	coll := mdb.db.Collection(collStudent)

	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	res, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return 0, fmt.Errorf("delete student: %w", err)
	}
	return res.DeletedCount, nil
}
