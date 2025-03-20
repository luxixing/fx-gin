package domain

import "time"

// Config 配置实体
type Config struct {
	ID        int64     `json:"id"`
	Key1      string    `json:"key1"`
	Key2      string    `json:"key2"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ConfigFilter 配置过滤条件
type ConfigFilter struct {
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
}

// ConfigRequest 创建配置请求
type ConfigRequest struct {
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
}

// ConfigResponse 配置响应
type ConfigResponse struct {
	ID        int64     `json:"id"`
	Key       string    `json:"key"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ConfigRepo 配置仓库接口
type ConfigRepo interface {
	Create(config *Config) error
	Get(key string) (*Config, error)
	List(filter *ConfigFilter) ([]*Config, error)
}

// ConfigService 配置服务接口
type ConfigService interface {
	List(filter *ConfigFilter) ([]*ConfigResponse, error)
	Get(key string) (*ConfigResponse, error)
	Create(req *ConfigRequest) (*ConfigResponse, error)
}
