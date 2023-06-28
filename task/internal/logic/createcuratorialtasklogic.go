package logic

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"task/internal/model"
	"time"

	"task/internal/svc"
	"task/task"

	"github.com/shopspring/decimal"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCuratorialTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCuratorialTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCuratorialTaskLogic {
	return &CreateCuratorialTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateCuratorialTask 创建策展任务
func (l *CreateCuratorialTaskLogic) CreateCuratorialTask(in *task.CreatePublishTaskInput) (*task.Mistake, error) {
	// 判断推特文章地址是否正确
	if !strings.Contains(in.TweetAddress, "twitter.com/") {
		return &task.Mistake{Msg: "推特文章地址不正确"}, nil
	}
	// 判断金额是否合规（是否小于0.01）
	if decimal.NewFromFloat(in.AwardAmount).LessThan(decimal.NewFromFloat(0.01)) {
		return &task.Mistake{Msg: "奖励金额最小为0.01"}, nil
	}
	// 判断用户数是否合规
	if in.MaxUser < 10 {
		return &task.Mistake{Msg: "奖励用户数量不能少于10人"}, nil
	}
	// 查询用户是否存在
	_, err := QueryUser(in.Creator)
	if err != nil {
		logx.Error(fmt.Sprintf("查询用户是否存在失败，err:%v", err))
		return &task.Mistake{Msg: err.Error()}, nil
	}
	// 查询资金数量
	Amount, err := UserTokenAmount(in.Creator)
	if err != nil {
		logx.Error("查询用户资金是否充足失败，err:", err)
		return &task.Mistake{Msg: err.Error()}, nil
	}
	// 判断资金是否充足
	if decimal.NewFromFloat(in.AwardBudget).GreaterThan(Amount) {
		return &task.Mistake{Msg: "用户资金不足"}, nil
	}
	// 获取循环数据
	var label string
	for _, item := range in.Label {
		// 判断标签是否存在
		labelSrt, err := l.svcCtx.LabelModel.FindOne(l.ctx, item, in.Creator)
		if err != nil {
			return &task.Mistake{Msg: err.Error()}, err
		}
		if labelSrt.Id == 0 {
			return &task.Mistake{Msg: "标签不存在"}, nil
		}
		label = label + "#" + item
	}
	// 执行创建任务信息
	_, err = CreateTaskInformation(l, in, in.Creator, label)
	if err != nil {
		return &task.Mistake{Msg: err.Error()}, err
	}
	// 每日任务—校验发布策展任务
	exist, err := verifyPublishCuration(l, in.Creator, "发布策展任务")
	if err != nil || !exist {
		return &task.Mistake{Msg: err.Error()}, err
	}
	return &task.Mistake{Msg: "succeed"}, nil
}

// QueryUser 查询用户是否存在(未完)
// @param userId uint 用戶ID   profile *model.Profile,
func QueryUser(userId string) (profile string, err error) {
	// if err := conf.MyDB.Model(model.Profile{}).
	// 	Where("user_id = ?", userId).
	// 	Find(&profile).Error; err != nil {
	// 	logs.Error("查询用户是否存在失败，err:%v", err)
	// 	return nil, err
	// }
	return
}

// UserTokenAmount 查询用户可用代币数量(未完)
// @param userId uint 用戶ID
func UserTokenAmount(userId string) (decimal.Decimal, error) {
	return decimal.NewFromFloat(100000), nil
}

// CreateTaskInformation 执行创建任务信息
func CreateTaskInformation(l *CreateCuratorialTaskLogic, in *task.CreatePublishTaskInput, userInfo, label string) (*task.Mistake, error) {
	// 附值
	publishTask := &model.PublishTask{
		CreatedAt:    sql.NullTime{Time: time.Now(), Valid: true},
		Creator:      userInfo,
		TweetAddress: in.TweetAddress,
		Label:        label,
		Status:       1,
		AwardBudget:  in.AwardBudget,
		MaxUser:      in.MaxUser,
		AwardAmount:  in.AwardAmount,
		EndTime:      time.Now().Add(time.Hour * 24 * 3),
	}
	// 添加用户
	_, err := l.svcCtx.PublishTaskModel.Insert(l.ctx, publishTask)
	if err != nil {
		return &task.Mistake{Msg: fmt.Sprintf("执行创建任务信息失败：%v", err)}, err
	}
	// 获取刚刚插入数据的ID（后期改为UUID,未完）
	publishTaskSrt, _ := l.svcCtx.PublishTaskModel.FindNewlyAddedData(l.ctx, userInfo, in.TweetAddress, label)
	// 添加任务类型
	for _, item := range in.TaskBak {
		taskDemand := &model.TaskDemand{
			CreatedAt:  sql.NullTime{Time: time.Now(), Valid: true},
			TaskId:     sql.NullInt64{Int64: publishTaskSrt.Id, Valid: true},
			TaskName:   sql.NullInt64{Int64: int64(item.TaskName), Valid: true},
			TaskDemand: sql.NullString{String: item.TaskDemand, Valid: true},
			Article:    sql.NullString{String: item.Article, Valid: true},
		}
		_, err := l.svcCtx.TaskDemandModel.Insert(l.ctx, taskDemand)
		if err != nil {
			return &task.Mistake{Msg: fmt.Sprintf("添加任务类型%v", err)}, err
		}
	}
	return &task.Mistake{Msg: "succeed"}, nil
}
