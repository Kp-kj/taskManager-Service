package logic

import (
	"context"
	"fmt"
	"taskManager-Service/internal/model"

	"taskManager-Service/internal/svc"
	"taskManager-Service/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryTaskListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTaskListLogic {
	return &QueryTaskListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryTaskList 查询策展任务列表
func (l *QueryTaskListLogic) QueryTaskList(in *task.PublishTaskInput) (*task.RePublishTask, error) {
	// 获取数据总量
	totalAmount, err := l.svcCtx.PublishTaskModel.FindPublishTaskAmount(l.ctx, int(in.Status), "status")
	if err != nil && err.Error() != "sql: no rows in result set" {
		fmt.Printf("1err:%v\n", err)
		return nil, err
	}
	// 计算分页
	startLine, err := MathPagination(totalAmount, &in.MaxNum, &in.CurrPage)
	if err != nil {
		fmt.Printf("2err:%v\n", err)
		return nil, err
	}
	// 查询数据
	taskList, err := l.svcCtx.PublishTaskModel.FindPublishTaskList(l.ctx, in.Status, in.MaxNum, startLine, "status")
	if err != nil && err.Error() != "sql: no rows in result set" {
		fmt.Printf("3err:%v\n", err)
		return nil, err
	}
	// 查询任务要求
	rePublishTaskBak, err := QueryTaskRequiresAndAssignsValue(l, taskList)
	if err != nil && err.Error() != "sql: no rows in result set" {
		fmt.Printf("4err:%v\n", err)
		return nil, err
	}
	// 赋值分页
	paginationData := task.PaginationData{
		Total:   totalAmount,
		Page:    in.CurrPage,
		PerPage: in.MaxNum,
	}
	// 赋值返回数据
	rePublishTask := task.RePublishTask{
		PaginationData:   &paginationData,
		RePublishTaskBak: rePublishTaskBak,
	}
	return &rePublishTask, nil
}

// MathPagination 分页计算
// @params total    int64   数据总量
// @params maxNum   *int64  每页最大输出数量
// @params currPage *int64 当前页码
// @return          int64   总页数
// @return          int64   偏移量
// @return          error   错误信息
func MathPagination(total int64, maxNum *int64, currPage *int64) (startLine int64, err error) {
	// 捕获异常
	defer func() {
		if errs := recover(); errs != nil {
			err = fmt.Errorf("分页计算发生错误：%s", errs)
			fmt.Print(err.Error())
		}
	}()
	// 限制每页最大数据量
	if *maxNum > 100 {
		*maxNum = 100
	} else if *maxNum <= 0 {
		*maxNum = 30
	}
	// 计算偏移查询量
	startLine = (*currPage - 1) * (*maxNum)
	return startLine, nil
}

// QueryTaskRequiresAndAssignsValue 查询任务要求并赋值
func QueryTaskRequiresAndAssignsValue(l *QueryTaskListLogic, taskList []*model.PublishTask) (rePublishTaskBak []*task.RePublishTaskBak, err error) {
	for _, itemSrt := range taskList {
		taskDemand, err := l.svcCtx.TaskDemandModel.FindList(l.ctx, itemSrt.Id)
		if err != nil {
			return nil, err
		}

		// 查询推特文章详细信息(未完)+用户信息

		// 赋值任务要求
		var taskDemandSrt []*task.TaskDemand
		for _, item := range taskDemand {
			taskDemandSrt = append(taskDemandSrt, &task.TaskDemand{
				TaskId:     uint64(item.TaskId.Int64),
				TaskName:   int32(item.TaskName.Int64),
				TaskDemand: item.TaskDemand.String,
				Article:    item.Article.String,
			})
		}
		// 赋值数据
		rePublishTaskBak = append(rePublishTaskBak, &task.RePublishTaskBak{
			TaskId:        uint64(itemSrt.Id),
			CreatedAt:     itemSrt.CreatedAt.Time.String(),
			Creator:       itemSrt.Creator,
			CreatorName:   itemSrt.Creator, // 未完
			CreatorNick:   itemSrt.Creator, // 未完
			CreatorAvatar: itemSrt.Creator, // 未完
			Status:        int32(itemSrt.Status),
			TweetDetails:  itemSrt.TweetAddress, // 未完
			TweetPicture:  itemSrt.TweetAddress, // 未完
			Label:         itemSrt.Label,
			AwardBudget:   itemSrt.AwardBudget,
			MaxUser:       itemSrt.MaxUser,
			AwardAmount:   itemSrt.AwardAmount,
			EndTime:       itemSrt.EndTime.String(),
			Accomplish:    itemSrt.Accomplish,
			TaskDemand:    taskDemandSrt,
		})
	}
	return
}
