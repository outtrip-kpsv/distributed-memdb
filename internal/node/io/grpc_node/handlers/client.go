package handlers

import (
	"context"
	"fmt"
	"github.com/frankenbeanies/uuid4"
	"go.uber.org/zap"
	cfg "team01/internal/config"
	"team01/internal/models"
	"team01/internal/node/bl"
	"team01/internal/proto/node"
)

type CliController struct {
	node.UnimplementedClientCommunicationServer
	Bl *bl.BL
}

func NewCliController(bl *bl.BL) *CliController {
	return &CliController{
		Bl: bl,
	}
}

func (n *CliController) GetInfoNode(ctx context.Context, req *node.PingRequest) (*node.Info, error) {
	cfg.GetLogger().Info("GetInfoNode")
	res := n.Bl.Node.GetInfo()
	return res, nil
}

func (n *CliController) Request(ctx context.Context, req *node.CliReq) (*node.NodeResp, error) {
	//comm := strings.Split(req.Value, " ")

	switch n.Bl.Node.GetInfo().Status {
	case node.NodeStatus_FOLLOWER:
		//todo proxy to master
	case node.NodeStatus_LISTENER, node.NodeStatus_LEADER:
		var uuid uuid4.UUID4
		var err error

		if req.Req.Uuid == "" {
			return models.ErrorResponse("пустой запрос или неизвестная команда"), nil
		} else {
			uuid, err = uuid4.ParseString(req.Req.Uuid)
			if err != nil {
				return models.ErrorResponse("проблема распознавания uuid4"), nil
			}
		}

		switch req.Req.Comm {
		case "GET":
			return models.OkResponse("Get " + uuid.String()), nil
		case "SET":
			if req.Req.Value == "" {
				return models.ErrorResponse("SET: недостаточно аргументов для выполнения"), nil
			}

			//TODO логика занесения в хранилище

			resp := node.NodeResp{
				Value:       "",
				Code:        0,
				CountResult: 0,
				Result:      make(map[string]bool)}
			for _, address := range req.Addresses {
				if address == cfg.GetAddress() {
					_, resp.Result[address] = n.Bl.Node.GetUnit().Vault.SetArtifact(req.Req.Uuid, req.Req.Value, req.Req.HashCluster)
					resp.CountResult++
					continue
				}

				cfg.GetLogger().Info(address, zap.Reflect("req.Value", req))

				resFromNode, err := n.Bl.Node.GetUnit().KnowNodes[address].Client.Set(context.Background(), req.Req)
				if err != nil {
					return nil, err
				}
				//if resFromNode.Result[address] {
				resp.Result[address] = resFromNode.Result[address]
				resp.CountResult++
				//}
			}

			///////
			var text string
			if resp.Result[req.Addresses[0]] {
				text = "добавлена"
			} else {
				text = "переписана"

			}
			if int(resp.CountResult) == len(req.Addresses) {
				resp.Value = fmt.Sprintf("Запись %s: (%d replicas)", text, resp.CountResult)
			} else {
				resp.Value = fmt.Sprintf("RRRR %s: (%d replicas)", text, resp.CountResult)

			}

			return &resp, nil
		case "DELETE":
			return models.OkResponse("Del"), nil
		}
	}

	return models.ErrorResponse("Errrrr"), nil
}
