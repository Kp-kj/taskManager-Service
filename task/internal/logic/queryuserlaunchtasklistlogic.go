package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"task/internal/svc"
)

type QueryUserLaunchTaskListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryUserLaunchTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserLaunchTaskListLogic {
	return &QueryUserLaunchTaskListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//func (l *QueryUserLaunchTaskListLogic) QueryUserLaunchTaskList(in *task.UserLaunchTaskListInput) (*task.RePublishTask, error) {
//	switch in.Status {
//	case 1:
//		// 获取数据总量
//		totalAmount, err := l.svcCtx.PublishTaskModel.FindPublishTaskAmount(l.ctx, int(in.Status), "creator")
//		if err != nil {
//			return nil, err
//		}
//		// 计算分页
//		startLine, err := MathPagination(totalAmount, &in.MaxNum, &in.CurrPage)
//		if err != nil {
//			return nil, err
//		}
//		// 查询个人发起任务列表
//		taskList, err := l.svcCtx.PublishTaskModel.FindPublishTaskList(l.ctx, in.Status, in.MaxNum, startLine, "creator")
//		if err != nil {
//			return nil, err
//		}
//	case 2:
//		// 查询个人参与任务
//
//	default:
//		return nil, fmt.Errorf("查询类型未知")
//	}
//	return &task.RePublishTask{}, nil
//}
