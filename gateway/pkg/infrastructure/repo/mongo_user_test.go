package repo

import (
	"context"
	"fmt"
	"testing"
	"time"

	centity "github.com/isutare412/hexago/common/pkg/entity"
	"github.com/isutare412/hexago/gateway/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestMongoUserRepo(t *testing.T) {
	var (
		id        = "isutare412"
		email     = "foo@bar.com"
		nickname  = "redshore"
		birthDate = time.Date(1993, 9, 25, 0, 0, 0, 0, time.UTC)
	)

	cfg, err := loadMongoTestConfig()
	assert.NoError(t, err)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	mongoRepo, err := NewMongoDB(ctx, cfg.MongoDB)
	assert.NoError(t, err)

	err = mongoRepo.InsertUser(ctx, &centity.User{
		Id:         id,
		Email:      email,
		Nickname:   nickname,
		GivenName:  "Suhyuk",
		FamilyName: "Lee",
		Birth:      birthDate,
	})
	assert.NoError(t, err)

	user, err := mongoRepo.FindUserById(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, user.Id, id)
	assert.Equal(t, user.Email, email)
	assert.Equal(t, user.Nickname, nickname)
	assert.Equal(t, user.Birth, birthDate)

	err = mongoRepo.DeleteUserById(ctx, id)
	assert.NoError(t, err)
}

func loadMongoTestConfig() (*config.Config, error) {
	cfg, err := config.Load("../../../configs/local/config.yaml")
	if err != nil {
		return nil, fmt.Errorf("loading config: %w", err)
	}

	cfg.MongoDB.Database = "hexago_test"
	return cfg, nil
}
