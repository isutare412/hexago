package main

import (
	"context"
	"time"

	"github.com/isutare412/hexago/pkg/config"
	"github.com/isutare412/hexago/pkg/core/entity"
	"github.com/isutare412/hexago/pkg/logger"
)

func main() {
	cfg := config.Config{
		Mode: "development",
		MongoDB: &config.MongoDBConfig{
			Uri:      "mongodb://hexagoer:hexagoer@localhost:27017/?authSource=hexago&replicaSet=replicaset",
			Database: "hexago",
		},
	}
	logger.Initialize(!cfg.IsProduction())
	defer logger.S().Sync()

	diCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	logger.S().Info("Start dependency injection")
	beans, err := dependencyInjection(diCtx, &cfg)
	if err != nil {
		logger.S().Fatalf("Failed to inject dependencies: %v", err)
	}
	logger.S().Info("Done dependency injection")

	appCtx := context.Background()
	stu, err := beans.stuSvc.AddStudent(appCtx, &entity.Student{
		GivenName:  "Suhyuk",
		FamilyName: "Lee",
		Birth:      time.Date(1993, 9, 25, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		logger.S().Fatalf("Inserting student: %v", err)
	}

	stu, err = beans.stuSvc.StudentById(appCtx, stu.Id)
	if err != nil {
		logger.S().Fatalf("Finding student: %v", err)
	}
	logger.S().Debugf("Student[%v]", *stu)

	err = beans.stuSvc.RemoveStudentById(appCtx, stu.Id)
	if err != nil {
		logger.S().Fatalf("Removing student: %v", err)
	}

	logger.S().Info("Start graceful shutdown")
	shutdown(beans)
	logger.S().Info("Done graceful shutdown")
}

func shutdown(beans *beans) {
	ctx := context.Background()
	shutdownCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if err := beans.mongoRepo.Close(shutdownCtx); err != nil {
		logger.S().Error("Failed to close mongodb: %v", err)
	}
}
