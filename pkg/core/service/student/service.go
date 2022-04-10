package student

import (
	"context"
	"fmt"

	"github.com/isutare412/hexago/pkg/core/entity"
	"github.com/isutare412/hexago/pkg/core/port"
	"github.com/isutare412/hexago/pkg/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
	repo port.StudentRepository
}

func NewService(repo port.StudentRepository) port.StudentService {
	return &service{
		repo: repo,
	}
}

func (s *service) AddStudent(
	ctx context.Context,
	stu *entity.Student,
) (*entity.Student, error) {
	sCtx, err := s.repo.StartSession(ctx)
	if err != nil {
		logger.S().Error(err)
		return nil, err
	}
	defer sCtx.EndSession(ctx)

	stu, err = s.repo.InsertStudent(sCtx, stu)
	if err != nil {
		logger.S().Error(err)
		return nil, err
	}

	if err := sCtx.CommitTransaction(ctx); err != nil {
		logger.S().Error(err)
		return nil, err
	}
	return stu, nil
}

func (s *service) StudentById(
	ctx context.Context,
	id primitive.ObjectID,
) (*entity.Student, error) {
	stu, err := s.repo.FindStudentById(ctx, id)
	if err != nil {
		logger.S().Error(err)
		return nil, err
	}
	return stu, nil
}

func (s *service) RemoveStudentById(
	ctx context.Context,
	id primitive.ObjectID,
) error {
	cnt, err := s.repo.DeleteStudentById(ctx, id)
	if err != nil {
		logger.S().Error(err)
		return err
	}
	if cnt != 1 {
		err := fmt.Errorf("student[%v] does not exist", id)
		logger.S().Error(err)
		return err
	}
	return nil
}
