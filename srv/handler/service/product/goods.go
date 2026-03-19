package product

import (
	"context"
	"gospacex-pengyilong/srv/basic/config"
	"gospacex-pengyilong/srv/handler/model"
	__ "gospacex-pengyilong/srv/proto/goods"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	__.UnimplementedGoodsServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) GoodsAdd(_ context.Context, in *__.GoodsAddReq) (*__.GoodsAddResp, error) {

	var goods model.Goods
	err := goods.FindGoods(config.DB, in.Name)
	if err != nil {
		return &__.GoodsAddResp{
			Msg:  "参数错误",
			Code: 400,
		}, nil
	}
	m := model.Goods{
		Name:  in.Name,
		Price: float64(in.Price),
		Num:   int(in.Num),
	}
	err = m.GoodsAdd(config.DB)
	if err != nil {
		return &__.GoodsAddResp{
			Msg:  "添加失败",
			Code: 400,
		}, nil
	}
	return &__.GoodsAddResp{
		Msg:  "添加成功",
		Code: 200,
	}, nil
}
