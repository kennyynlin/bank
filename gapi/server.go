package gapi

import (
	"fmt"
	db "github.com/kennyynlin/bank/db/sqlc"
	"github.com/kennyynlin/bank/pb"
	"github.com/kennyynlin/bank/token"
	"github.com/kennyynlin/bank/util"
	"github.com/kennyynlin/bank/worker"
)

type Server struct {
	pb.UnimplementedBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %w", err)
	}
	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}
	return server, nil
}
