package config

import (
    "context"
    "errors"
    "fmt"
    "time"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"

    mg "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"

    "abc-user-backend/models"
    mysqlRepo "abc-user-backend/repositories/mysql"
    mongoRepo "abc-user-backend/repositories/mongo"
    "abc-user-backend/repositories"
)

func ConnectDB(cfg *Config) (repositories.UserRepository, error) {
    switch cfg.DBType {
    case "mysql":
        dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
            cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)

        db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
        if err != nil {
            return nil, err
        }

        // Auto migrate the User model
        if err := db.AutoMigrate(&models.User{}); err != nil {
            return nil, err
        }

        return mysqlRepo.NewUserRepo(db), nil

    case "mongo":
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()

        uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort)
        clientOpts := options.Client().ApplyURI(uri)

        client, err := mg.Connect(ctx, clientOpts)
        if err != nil {
            return nil, err
        }

        err = client.Ping(ctx, nil)
        if err != nil {
            return nil, err
        }

        db := client.Database(cfg.DBName)
        return mongoRepo.NewUserRepo(db), nil

    default:
        return nil, errors.New("unsupported DB_TYPE in config")
    }
}
