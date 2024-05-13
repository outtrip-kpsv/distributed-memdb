package midleware

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	cfg "team01/internal/config"
	"team01/internal/proto/node"
	"time"
)

//todo все гетноде переписать

// ClientRequestInterceptor клентский мидлваре выполняемый перед запросом
func (m *midelwares) ClientRequestInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {

	if cc.GetState() != connectivity.Ready {
		delete(m.bl.Node.GetUnit().KnowNodes, cc.Target())
		cfg.GetLogger().Info("Нет соединение с нодой", zap.String("address", cc.Target()))
		return errors.New("Server connection is not ready")
	}

	resultErr := invoker(ctx, method, req, reply, cc, opts...)
	// посе запроса
	if resultErr == nil && method != "/NodeCommunication/GetKnownNodes" {
		res := &node.KnownNodes{}
		err := cc.Invoke(ctx, "/NodeCommunication/GetKnownNodes", m.bl.Node.GetKnowNode(), res, opts...)
		if err != nil {
			cfg.GetLogger().Info("ERR", zap.Error(err))
			return err
		}
		cfg.GetLogger().Info("GetNodes: ", zap.Reflect("nodes: ", res))

		// todo найти самый стары запрос и обновить LastNode (+)
		oldTime := time.Now()
		for k, v := range res.Nodes {
			//TODO Address ->cfg
			if k == cfg.GetAddress() {
				continue
			}
			err := m.bl.Node.GetUnit().ConnectTo(k, v.Ts)
			if err != nil {
				return err
			}
			cfg.GetLogger().Info("find time for last node check " + k)
			if v.Ts.AsTime().Before(oldTime) {
				oldTime = v.Ts.AsTime()
				m.bl.Node.GetUnit().LastNode.Address = k
				cfg.GetLogger().Info("last node naw is " + k)
			} else {
				cfg.GetLogger().Info("ln old " + m.bl.Node.GetUnit().LastNode.Address)
				//fmt.Println()
			}
		}
		m.bl.Node.GetUnit().LastNode.Ticker.Stop()
		tmp := time.Second*5 - (time.Now().Sub(oldTime))
		cfg.GetLogger().Info("-------- " + tmp.String())
		//fmt.Println("t", tmp)
		if tmp <= 0 {
			tmp = time.Millisecond
		}
		m.bl.Node.GetUnit().LastNode.Ticker = time.NewTicker(tmp)

	}

	return resultErr
}
