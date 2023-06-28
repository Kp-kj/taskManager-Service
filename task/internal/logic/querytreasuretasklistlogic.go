package logic

import (
	"context"
	"fmt"
	"strings"
	"task/internal/model"
	"task/internal/svc"
	"task/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryTreasureTaskListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryTreasureTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTreasureTaskListLogic {
	return &QueryTreasureTaskListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryTreasureTaskList 查询+搜索宝箱样式列表
func (l *QueryTreasureTaskListLogic) QueryTreasureTaskList(in *task.TreasureTaskListInput) (*task.ReTreasureTaskSrt, error) {
	// 判断参数
	if strings.Count(in.Name, "") >= 21 {
		return nil, fmt.Errorf("任务名称长度不合规")
	}
	// 校验奖励个数
	if in.Reward > 4 {
		return nil, fmt.Errorf("经验奖励个数不合规")
	}
	// 条件搜索
	treasureTask, total, err := ConditionalSearch(l, in)
	if err != nil {
		return nil, err
	}
	// 赋值数据+搜索宝箱样式
	reTreasureTask, err := SearchTreasureChestStyles(l, treasureTask)
	if err != nil {
		return nil, err
	}
	// 赋值分页数据
	paginationData := task.PaginationData{
		Total:   total,
		Page:    in.CurrPage,
		PerPage: in.MaxNum,
	}
	reTreasureTaskSrt := task.ReTreasureTaskSrt{
		PaginationData:       &paginationData,
		TreasureTaskSrtInput: reTreasureTask,
	}
	return &reTreasureTaskSrt, nil
}

// ConditionalSearch 条件搜索
func ConditionalSearch(l *QueryTreasureTaskListLogic, in *task.TreasureTaskListInput) (treasureTask []*model.TreasureTask, total int64, err error) {
	// 名称+数据搜索
	if in.Name != "" && in.Reward > 0 {
		// 统计数量
		total, err = l.svcCtx.TreasureTaskModel.FindNameAndQuantityCount(l.ctx, in.Reward, in.Name)
		if err != nil {
			return nil, 0, err
		}
		// 计算分页
		startLine, err := MathPagination(total, &in.MaxNum, &in.CurrPage)
		if err != nil {
			return nil, total, err
		}
		// 获取数量
		treasureTask, err = l.svcCtx.TreasureTaskModel.FindNameAndQuantity(l.ctx, in.Reward, in.Name, startLine, in.MaxNum)
		if err != nil {
			return nil, 0, err
		}
	}
	// 查看全部
	if in.Name == "" && in.Reward == 0 {
		// 查看全部统计数量
		total, err = l.svcCtx.TreasureTaskModel.FindAllCount(l.ctx)
		if err != nil {
			return nil, 0, err
		}
		// 计算分页
		startLine, err := MathPagination(total, &in.MaxNum, &in.CurrPage)
		if err != nil {
			return nil, total, err
		}
		// 查看全部a
		treasureTask, err = l.svcCtx.TreasureTaskModel.FindAll(l.ctx, startLine, in.MaxNum)
		if err != nil {
			return nil, 0, err
		}
	}
	// 按名称搜索
	if in.Name != "" && in.Reward == 0 {
		// 按名称搜索统计数量
		total, err = l.svcCtx.TreasureTaskModel.FindNameCount(l.ctx, in.Name)
		if err != nil {
			return nil, 0, err
		}
		// 计算分页
		startLine, err := MathPagination(total, &in.MaxNum, &in.CurrPage)
		if err != nil {
			return nil, total, err
		}
		// 按名称搜索
		treasureTask, err = l.svcCtx.TreasureTaskModel.FindName(l.ctx, in.Name, startLine, in.MaxNum)
		if err != nil {
			return nil, 0, err
		}
	}
	// 按奖励个数搜索
	if in.Name == "" && in.Reward > 0 {
		// 按名称搜索统计数量
		total, err = l.svcCtx.TreasureTaskModel.FindQuantityCount(l.ctx, in.Reward)
		if err != nil {
			return nil, 0, err
		}
		// 计算分页
		startLine, err := MathPagination(total, &in.MaxNum, &in.CurrPage)
		if err != nil {
			return nil, total, err
		}
		// 按奖励个数搜索
		treasureTask, err = l.svcCtx.TreasureTaskModel.FindQuantity(l.ctx, in.Reward, startLine, in.MaxNum)
		if err != nil {
			return nil, 0, err
		}
	}
	return treasureTask, total, nil
}

// SearchTreasureChestStyles 赋值数据+搜索宝箱样式
func SearchTreasureChestStyles(l *QueryTreasureTaskListLogic, treasureTask []*model.TreasureTask) (reTreasureTask []*task.TreasureTaskSrtInput, err error) {
	// 赋值数据+搜索宝箱样式
	for _, item := range treasureTask {
		treasureTaskStage, err := l.svcCtx.TreasureTaskStageModel.FindTreasureInformation(l.ctx, item.Id, "treasure")
		if err != nil {
			return nil, err
		}
		// 循环赋值
		var treasureTaskStageSrt []*task.TreasureTaskStage
		for _, item1 := range treasureTaskStage {
			treasureTaskStageSrt = append(treasureTaskStageSrt, &task.TreasureTaskStage{
				ID:               item1.Id,
				Treasure:         item1.Treasure,
				TreasureSequence: item1.TreasureSequence,
				StageExperience:  item1.StageExperience,
				StageReward:      item1.StageReward,
			})
		}
		// 赋值数据
		reTreasureTask = append(reTreasureTask, &task.TreasureTaskSrtInput{
			Id:                uint64(item.Id),
			Name:              item.Name,
			DemandIntegral:    item.DemandIntegral,
			TaskReward:        item.TaskReward,
			ExperienceReward:  item.ExperienceReward,
			RewardQuantity:    item.RewardQuantity,
			TreasureTaskStage: treasureTaskStageSrt,
		})
	}
	return
}
