package nlp

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/nlp/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"google.golang.org/grpc"
)

func NewApp(c *config.AppConfig, logger *logger.Logger, rs *rpc.Server) (*app.Application, error) {
	// service
	s := service.NewNLP()
	err := rs.Register(func(gs *grpc.Server) error {
		pb.RegisterNLPServer(gs, s)
		return nil
	})
	if err != nil {
		return nil, err
	}

	a, err := app.New(c, logger, app.RPCServerOption(rs))
	if err != nil {
		return nil, err
	}

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp)