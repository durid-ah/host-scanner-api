package handler

import (
	"context"
	"github.com/durid-ah/nmap-api/db"
)

type HostResponse struct {
	Body []db.Host `json:"hosts"`
}

func GetAllHosts(storage *db.Storage) func(c context.Context, _ *struct{}) (*HostResponse, error) {
	
	return func(c context.Context, _ *struct{}) (*HostResponse, error) {
		hosts, err := storage.GetAllHosts(c)
		if err != nil {
			return nil, err
		}
		return &HostResponse{Body: hosts}, nil
	}
}
