package service

import (
	"time"

	"github.com/luxixing/fx-gin/internal/config"
	"github.com/luxixing/fx-gin/internal/domain"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ConfigServiceParams 配置服务参数
type ConfigServiceParams struct {
	fx.In
	Config *config.Config
	Logger *zap.Logger
	Repo   domain.ConfigRepo
}

// configService 配置服务实现
type configService struct {
	cfg    *config.Config
	logger *zap.Logger
	repo   domain.ConfigRepo
}

// NewConfigService 创建配置服务
func NewConfigService(p ConfigServiceParams) domain.ConfigService {
	return &configService{
		cfg:    p.Config,
		logger: p.Logger,
		repo:   p.Repo,
	}
}

// Create 创建配置
func (s *configService) Create(req *domain.ConfigRequest) (*domain.ConfigResponse, error) {
	s.logger.Info("创建配置", zap.Any("request", req))

	now := time.Now()
	entity := &domain.Config{
		Key1:      req.Key,
		Key2:      req.Value,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := s.repo.Create(entity)
	if err != nil {
		return nil, err
	}

	return &domain.ConfigResponse{
		ID:        entity.ID,
		Key:       entity.Key1,
		Value:     entity.Key2,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}, nil
}

// Get 获取配置
func (s *configService) Get(key string) (*domain.ConfigResponse, error) {
	s.logger.Info("获取配置", zap.String("key", key))

	entity, err := s.repo.Get(key)
	if err != nil {
		return nil, err
	}

	if entity == nil {
		return nil, nil
	}

	return &domain.ConfigResponse{
		ID:        entity.ID,
		Key:       entity.Key1,
		Value:     entity.Key2,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}, nil
}

// List 列出配置
func (s *configService) List(filter *domain.ConfigFilter) ([]*domain.ConfigResponse, error) {
	s.logger.Info("列出配置", zap.Any("filter", filter))

	entities, err := s.repo.List(filter)
	if err != nil {
		return nil, err
	}

	var result []*domain.ConfigResponse
	for _, entity := range entities {
		result = append(result, &domain.ConfigResponse{
			ID:        entity.ID,
			Key:       entity.Key1,
			Value:     entity.Key2,
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,
		})
	}

	return result, nil
}
