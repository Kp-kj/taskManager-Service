package logic

import (
	"context"
	"fmt"

	"task/internal/svc"
	"task/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeTreasureTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangeTreasureTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeTreasureTaskLogic {
	return &ChangeTreasureTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ChangeTreasureTask 上架+删除宝箱样式
// @param id     uint 宝箱样式ID
// @param status int  更新状态，1-删除，0-上架
func (l *ChangeTreasureTaskLogic) ChangeTreasureTask(in *task.TreasureTaskInput) (*task.Mistake, error) {
	// 判断是上架还是删除
	switch in.Status {
	case 0:
		err := GroundingTreasureTask(l, in)
		if err != nil {
			return &task.Mistake{Msg: fmt.Sprintf("删除宝箱样式失败:%v", err)}, err
		}
	case 1:
		err := DeleteTreasureTask(l, in)
		if err != nil {
			return &task.Mistake{Msg: err.Error()}, err
		}
	default:
		return nil, fmt.Errorf("更新类型不合规")
	}
	return &task.Mistake{Msg: "succeed"}, nil
}

// GroundingTreasureTask 上架宝箱样式
func GroundingTreasureTask(l *ChangeTreasureTaskLogic, in *task.TreasureTaskInput) error {
	// 查看宝箱是否存在
	treasureTask, err := l.svcCtx.TreasureTaskModel.FindOne(l.ctx, int64(in.Id))
	if err != nil && err.Error() != "sql: no rows in result set" {
		return err
	}
	// 判断是否存在
	if treasureTask == nil {
		return fmt.Errorf("宝箱样式不存在")
	}
	// 更新redis上架宝箱的数量（未完）
	//err := tools.AmendTreasureRedis("treasureAmount", treasureTask.ExperienceReward, 0)
	//if err != nil {
	//	logs.Error("上架宝箱样式_更新宝箱redis数量失败，err:%v", err)
	//	return err
	//}
	// 下架旧样式
	err = l.svcCtx.TreasureTaskModel.UpdateOldStyle(l.ctx, in.Id)
	if err != nil {
		return err
	}
	// 上架
	err = l.svcCtx.TreasureTaskModel.UpdateNewStyle(l.ctx, in.Id)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTreasureTask 删除宝箱样式
func DeleteTreasureTask(l *ChangeTreasureTaskLogic, in *task.TreasureTaskInput) error {
	treasureTask, err := l.svcCtx.TreasureTaskModel.FindOne(l.ctx, int64(in.Id))
	fmt.Printf("删除宝箱样式:%v", err)
	if err != nil {
		return err
	}
	// 判断是否上架
	if treasureTask.Employ == 1 {
		return fmt.Errorf("上架任务不允许删除")
	}
	err = l.svcCtx.TreasureTaskModel.Delete(l.ctx, int64(in.Id))
	if err != nil {
		return err
	}
	return nil
}

//// AmendTreasureRedis 修改redis宝箱数量
//func AmendTreasureRedis(key string, value int, expiration time.Duration) error {
//	date, err := conf.Rdb.Set(key, value, expiration).Result()
//	if err != nil {
//		err = fmt.Errorf("add redis err key: %v, token: %v, err: %v", date, value, err)
//		logs.Error(err)
//		return err
//	}
//
//	logs.Info("add redis success key: %s, token: %s\r\n", date, value)
//	return nil
//}
