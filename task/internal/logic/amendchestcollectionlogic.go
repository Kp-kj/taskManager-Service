package logic

import (
	"context"
	"database/sql"
	"fmt"
	"taskManager-Service-main/task/internal/model"

	"taskManager-Service-main/task/internal/svc"
	"taskManager-Service-main/task/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type AmendChestCollectionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAmendChestCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AmendChestCollectionLogic {
	return &AmendChestCollectionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AmendChestCollectionLogic) AmendChestCollection(in *task.AmendChestCollectionInput) (*task.Mistake, error) {
	// 查询是否存在个人宝箱领取度
	chestCollection, err := l.svcCtx.ChestCollectionsModel.FindIndividualAllowance(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	// 判断个人进度是否存在，存在更新，不存在创建
	if chestCollection.Id > 0 { // 存在进行更新
		// 需要获取redis 或者直接读数据库
		treasureTask, err := l.svcCtx.TreasureTaskModel.FindTreasureQuantity(l.ctx)
		if err != nil {
			return nil, err
		}
		if chestCollection.ChestAmount.Int64+in.Amount > treasureTask.ExperienceReward {
			return &task.Mistake{Err: fmt.Sprintf("经验超出预算数量")}, fmt.Errorf("宝箱已全部领取")
		}
		// 更新个人宝箱领取数
		err = l.svcCtx.ChestCollectionsModel.UpdateIndividualAllowance(l.ctx, in.UserId, in.Amount)
		if err != nil {
			return nil, err
		}
	} else {
		// 不存在进行创新
		chestCollection := model.ChestCollections{
			UserId:      in.UserId,
			ChestAmount: sql.NullInt64{Int64: in.Amount},
		}
		// 添加宝箱领取度
		_, err := l.svcCtx.ChestCollectionsModel.Insert(l.ctx, &chestCollection)
		if err != nil {
			return nil, err
		}
	}
	return &task.Mistake{}, nil
}
