package db

import (
	"context"
	"fmt"
	"log/slog"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Storage struct {
	db     *gorm.DB
	logger *slog.Logger
}

func NewStorage(logger *slog.Logger) (*Storage, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Host{})

	return &Storage{db: db, logger: logger}, nil
}

func (s *Storage) CreateHost(ctx context.Context, host *Host) error {

	return gorm.G[Host](s.db).Create(ctx, host)
}

func (s *Storage) CreateHosts(ctx context.Context, hosts []Host) error {
	return gorm.G[Host](s.db).CreateInBatches(ctx, &hosts, 100)
}

func (s *Storage) UpdateHost(ctx context.Context, host *Host) error {
	affectedRows, err := gorm.G[Host](s.db).
		Where("hostname = ?", host.Hostname).
		Updates(ctx, *host)

	if err != nil {
		s.logger.Error("failed to update host", "error", err, "host", host)
		return err
	}
	if affectedRows == 0 {
		s.logger.Error("no host found to update", "host", host)
		return fmt.Errorf("no host found to update: %w", err)
	}
	return nil
}

func (s *Storage) GetHostIPMap(ctx context.Context) (map[string]string, error) {
	hosts, err := gorm.G[Host](s.db).Find(ctx)
	if err != nil {
		s.logger.Error("failed to get all hosts", "error", err)
		return nil, err
	}
	hostIPMap := make(map[string]string, len(hosts))
	for _, host := range hosts {
		hostIPMap[host.Hostname] = host.IP
	}
	return hostIPMap, nil
}

func (s *Storage) DeleteHost(ctx context.Context, hostname string) error {
	result, err := gorm.G[Host](s.db).
		Where("hostname = ?", hostname).
		Delete(ctx)
	if err != nil {
		s.logger.Error("failed to delete host", "error", err, "hostname", hostname)
		return err
	}
	if result == 0 {
		s.logger.Warn("no host found to delete", "hostname", hostname)
		return fmt.Errorf("no host found to delete")
	}
	return nil
}

func (s *Storage) DeleteHosts(ctx context.Context, hostnames []string) error {
	result, err := gorm.G[Host](s.db).
		Where("hostname IN ?", hostnames).
		Delete(ctx)
	if err != nil {
		s.logger.Error("failed to delete hosts", "error", err, "hostnames", hostnames)
		return err
	}
	
	if result == 0 {
		s.logger.Warn("no hosts found to delete", "hostnames", hostnames)
		return fmt.Errorf("no hosts found to delete")
	}
	return nil
}