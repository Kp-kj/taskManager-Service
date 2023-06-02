package logic

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"task/internal/model"
	"task/internal/svc"
	"task/task"
	"time"

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
	if strings.Contains(in.TweetAddress, "twitter.com/") {
		return nil, fmt.Errorf("推特文章地址不正确")
	}
	// 判断金额是否合规（是否小于0.01）
	if decimal.NewFromFloat(in.AwardAmount).LessThan(decimal.NewFromFloat(0.01)) {
		return nil, fmt.Errorf("奖励金额最小为0.01")
	}
	// 判断用户数是否合规
	if in.MaxUser < 10 {
		return nil, fmt.Errorf("奖励用户数量不能少于10人")
	}
	// 查询用户是否存在
	userInfo, err := QueryUser(in.Creator)
	if err != nil {
		logx.Error(fmt.Sprintf("查询用户是否存在失败，err:%v", err))
		return nil, err
	}
	// 查询资金数量
	Amount, err := UserTokenAmount(in.Creator)
	if err != nil {
		logx.Error("查询用户资金是否充足失败，err:%v", err)
		return nil, err
	}
	// 判断资金是否充足
	if decimal.NewFromFloat(in.AwardBudget).GreaterThan(Amount) {
		return nil, fmt.Errorf("用户资金不足")
	}
	// 获取循环数据
	var label string
	for _, item := range in.Label {
		label = label + item
	}
	mistake, err := CreateTaskInformation(l, in, userInfo, label)
	if err != nil {
		return mistake, err
	}
	return &task.Mistake{}, nil
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
	return decimal.Decimal{}, nil
}

// CreateTaskInformation 执行创建任务信息
func CreateTaskInformation(l *CreateCuratorialTaskLogic, in *task.CreatePublishTaskInput, userInfo, label string) (*task.Mistake, error) {
	// 附值
	publishTask := &model.PublishTask{
		Creator:      userInfo,
		TweetAddress: in.TweetAddress,
		Label:        label,
		AwardBudget:  in.AwardBudget,
		MaxUser:      in.MaxUser,
		AwardAmount:  in.AwardAmount,
		EndTime:      time.Now().Add(time.Hour * 24 * 3),
	}
	// 添加用户
	result, err := l.svcCtx.PublishTaskModel.Insert(l.ctx, publishTask)
	if err != nil {
		return &task.Mistake{Err: fmt.Sprintf("%v", result)}, err
	}
	// 添加任务类型
	for _, item := range in.TaskBak {
		taskDemand := &model.TaskDemand{
			TaskId:     sql.NullInt64{Int64: int64(item.TaskId), Valid: true},
			TaskName:   sql.NullInt64{Int64: int64(item.TaskName), Valid: true},
			TaskDemand: sql.NullString{String: item.TaskDemand, Valid: true},
			Article:    sql.NullString{String: item.Article, Valid: true},
		}
		result1, err := l.svcCtx.TaskDemandModel.Insert(l.ctx, taskDemand)
		if err != nil {
			return &task.Mistake{Err: fmt.Sprintf("%v", result1)}, err
		}
	}
	return &task.Mistake{}, nil
}
