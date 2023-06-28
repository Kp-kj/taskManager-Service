package logic

import (
	"context"
	"fmt"
	"taskManager-Service/task/internal/model"

	"taskManager-Service/task/internal/svc"
	"taskManager-Service/task/task"

	"github.com/zeromicro/go-zero/core/logx"
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

// QueryUserLaunchTaskList 查询个人发起任务列表+参与任务
func (l *QueryUserLaunchTaskListLogic) QueryUserLaunchTaskList(in *task.UserLaunchTaskListInput) (*task.RePublishTask, error) {
	// 获取任务列表
	publishTaskList, totalAmount, err := GetTheTaskList(l, in)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	// 获取任务要求
	userTaskList, err := ObtainTaskRequirements(l, publishTaskList)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	// 赋值分页数据
	PaginationData := task.PaginationData{
		Total:   totalAmount,
		Page:    in.CurrPage,
		PerPage: in.MaxNum,
	}
	// 赋值返回数据
	rePublishTask := task.RePublishTask{
		PaginationData:   &PaginationData,
		RePublishTaskBak: userTaskList,
	}
	return &rePublishTask, nil
}

// GetTheTaskList 获取任务列表
func GetTheTaskList(l *QueryUserLaunchTaskListLogic, in *task.UserLaunchTaskListInput) ([]*model.PublishTask, int64, error) {
	var publishTaskList []*model.PublishTask
	var totalAmount int64
	var err error
	switch in.Status {
	case 1:
		// 查询个人发起任务列表
		// 获取数据总量
		totalAmount, err = l.svcCtx.PublishTaskModel.FindPublishTaskAmount(l.ctx, in.UserId, "creator")
		if err != nil && err.Error() != "sql: no rows in result set" {
			return nil, totalAmount, err
		}
		// 计算分页
		startLine, err := MathPagination(totalAmount, &in.MaxNum, &in.CurrPage)
		if err != nil && err.Error() != "sql: no rows in result set" {
			return nil, totalAmount, err
		}
		// 查询个人发起任务列表
		publishTaskList, err = l.svcCtx.PublishTaskModel.FindPublishTaskList(l.ctx, in.UserId, in.MaxNum, startLine, "creator")
		if err != nil && err.Error() != "sql: no rows in result set" {
			return nil, totalAmount, err
		}
	case 2:
		// 查询个人参与任务
		// 查询参与列表数量

		totalAmount, err = l.svcCtx.ParticipantModel.FindParticipantAmount(l.ctx, in.UserId, "user_id") // 改为userID，查询sql需要重写
		if err != nil {
			return nil, totalAmount, err
		}
		// 计算分页
		startLine, err := MathPagination(totalAmount, &in.MaxNum, &in.CurrPage)
		if err != nil {
			return nil, totalAmount, err
		}
		// 获取任务ID
		participatingIdSrt, err := l.svcCtx.ParticipantModel.GetListIndividualParticipating(l.ctx, in.UserId, in.MaxNum, startLine)
		if err != nil && err.Error() != "sql: no rows in result set" {
			return nil, totalAmount, err
		}
		var participatingId string
		for _, item := range participatingIdSrt {
			if participatingId == "" {
				participatingId = fmt.Sprintf("'%v'%v", item.Id, participatingId)
			} else {
				participatingId = fmt.Sprintf("'%v',%v", item.Id, participatingId)
			}

		}
		participatingId = fmt.Sprintf("(%v)", participatingId)
		// 根据ID获取任务信息
		publishTaskList, err = l.svcCtx.PublishTaskModel.FindTaskInformationBasedID(l.ctx, participatingId)
		if err != nil && err.Error() != "sql: no rows in result set" {
			return nil, totalAmount, err
		}
	default:
		return nil, totalAmount, fmt.Errorf("查询类型未知")
	}
	return publishTaskList, totalAmount, err
}

// ObtainTaskRequirements 获取任务要求
func ObtainTaskRequirements(l *QueryUserLaunchTaskListLogic, publishTaskList []*model.PublishTask) ([]*task.RePublishTaskBak, error) {
	var userTaskList []*task.RePublishTaskBak
	for _, item := range publishTaskList {
		taskDemand, err := l.svcCtx.TaskDemandModel.FindList(l.ctx, item.Id)
		if err != nil {
			return nil, err
		}
		// 获取推特详情信息//（未完）
		// 查询用户信息 （未完）
		// 赋值返回数据
		var reTaskDemand []*task.TaskDemand
		for _, item1 := range taskDemand {
			reTaskDemand = append(reTaskDemand, &task.TaskDemand{
				TaskId:     uint64(item1.TaskId.Int64),
				TaskName:   int32(item1.TaskName.Int64),
				TaskDemand: item1.TaskDemand.String,
				Article:    item1.Article.String,
			})
		}
		// 赋值返回
		userTaskList = append(userTaskList, &task.RePublishTaskBak{
			TaskId:        uint64(item.Id),
			CreatedAt:     item.CreatedAt.Time.String(),
			Creator:       item.Creator,
			CreatorName:   "", // 未完，用户名称
			CreatorNick:   "", // 未完，用户昵称
			CreatorAvatar: "", // 未完，用户头像
			Status:        int32(item.Status),
			TweetDetails:  "", // 未完，推特详情
			TweetPicture:  "", // 未完，推特文章图片
			Label:         item.Label,
			AwardBudget:   item.AwardBudget,
			MaxUser:       item.MaxUser,
			AwardAmount:   item.AwardAmount,
			EndTime:       item.EndTime.String(),
			Accomplish:    item.Accomplish,
			TaskDemand:    reTaskDemand,
		})
	}
	return userTaskList, nil
}
