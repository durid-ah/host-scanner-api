package handler

import (
	"context"
	"github.com/durid-ah/host-scanner-api/db"
)

type HostsResponse struct {
	Body []db.Host `json:"hosts"`
}

func GetAllHosts(storage *db.Storage) func(c context.Context, _ *struct{}) (*HostsResponse, error) {
	
	return func(c context.Context, _ *struct{}) (*HostsResponse, error) {
		hosts, err := storage.GetAllHosts(c)
		if err != nil {
			return nil, err
		}
		return &HostsResponse{Body: hosts}, nil
	}
}

type HostResponse struct {
	Body db.Host `json:"host"`
}

type HostParams struct {
	Hostname string `path:"hostname"`
}

func GetHost(storage *db.Storage) func(c context.Context, params *HostParams) (*HostResponse, error) {
	return func(c context.Context, params *HostParams) (*HostResponse, error) {
		host, err := storage.GetHost(c, params.Hostname)
		if err != nil {
			return nil, err
		}
		return &HostResponse{Body: *host}, nil
	}
}