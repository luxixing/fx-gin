package repo

import (
	"database/sql"

	"github.com/luxixing/fx-gin/internal/domain"
	"go.uber.org/zap"
)

type configRepo struct {
	db     *sql.DB
	logger *zap.Logger
}

func NewConfigRepo(db *sql.DB, logger *zap.Logger) domain.ConfigRepo {
	return &configRepo{
		db:     db,
		logger: logger,
	}
}
func (r *configRepo) Create(entity *domain.Config) error {
	//todo: impletation as your need
	r.logger.Info("create config", zap.Any("entity", entity))
	return nil
}
func (r *configRepo) Get(key string) (*domain.Config, error) {
	//todo: impletation as your need
	r.logger.Info("get config", zap.String("key", key))
	return &domain.Config{Key1: key}, nil
}
func (r *configRepo) List(filter *domain.ConfigFilter) ([]*domain.Config, error) {
	//todo: impletation as your need
	r.logger.Info("list config", zap.Any("filter", filter))
	return []*domain.Config{
		{Key1: "key1"},
		{Key1: "key2"},
	}, nil
}
