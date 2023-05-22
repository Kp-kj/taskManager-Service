package logic

import (
	"context"

	"task/internal/svc"
	"task/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *task.Request) (*task.Response, error) {
	// todo: add your logic here and delete this line
	// 测试连接
	return &task.Response{
		Pong: "pong",
	}, nil
	return &task.Response{}, nil
}
