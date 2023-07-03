package logic

import (
	"context"
	"taskManager-Service/internal/model"
	"taskManager-Service/internal/svc"
	"taskManager-Service/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySubtaskStyleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySubtaskStyleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySubtaskStyleLogic {
	return &QuerySubtaskStyleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QuerySubtaskStyle 查询子任务样式列表
func (l *QuerySubtaskStyleLogic) QuerySubtaskStyle(in *task.TaskIDInquireInput) (*task.ReSubtaskStyle, error) {
	// 判断是否进行搜索
	var subtaskStyle []*model.SubtaskStyle
	var err error
	if in.Id > 0 {
		subtaskStyle, err = l.svcCtx.SubtaskStyleModel.FindSubtaskStyles(l.ctx, int64(in.Id))
		if err != nil {
			return nil, err
		}
	} else {
		subtaskStyle, err = l.svcCtx.SubtaskStyleModel.FindSubtaskStyleList(l.ctx)
		if err != nil {
			return nil, err
		}
	}
	// 赋值返回数据
	var subtaskStyleList []*task.SubtaskStyle
	for _, item := range subtaskStyle {
		subtaskStyleList = append(subtaskStyleList, &task.SubtaskStyle{
			TaskId:         item.TaskId.Int64,
			TaskName:       item.TaskName.String,
			TaskNameEng:    item.TaskNameEng.String,
			TaskDetails:    item.TaskDetails.String,
			TaskDetailsEng: item.TaskDetailsEng.String,
			TaskStatus:     item.TaskStatus.Int64,
		})
	}
	reSubtaskStyleList := &task.ReSubtaskStyle{SubtaskStyle: subtaskStyleList}
	return reSubtaskStyleList, nil
}
