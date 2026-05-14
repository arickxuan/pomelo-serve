package service

import (
	"context"
	"pomeloServe/common/biz"
	"pomeloServe/core/dao"
	"pomeloServe/core/models/entity"
	"pomeloServe/core/models/requests"
	"pomeloServe/core/repo"
	"pomeloServe/framework/msError"
	"pomeloServe/proto/pd"
	"time"
)

//创建账号

type AccountService struct {
	accountDao *dao.AccountDao
	redisDao   *dao.RedisDao
	pd.UnimplementedUserServiceServer
}

func NewAccountService(manager *repo.Manager) *AccountService {
	return &AccountService{
		accountDao: dao.NewAccountDao(manager),
		redisDao:   dao.NewRedisDao(manager),
	}
}

func (a *AccountService) Register(ctx context.Context, req *pd.RegisterRequest) (*pd.RegisterResponse, error) {
	//写注册的业务逻辑
	if req.LoginPlatform == requests.WeiXin {
		ac, err := a.wxRegister(req)
		if err != nil {
			return &pd.RegisterResponse{}, msError.GrpcError(err)
		}
		return &pd.RegisterResponse{
			Uid: ac.Uid,
		}, nil
	}
	return &pd.RegisterResponse{}, nil
}

func (a *AccountService) wxRegister(req *pd.RegisterRequest) (*entity.Account, *msError.Error) {
	//1.封装一个account结构 将其存入数据库  mongo 分布式id objectID
	ac := &entity.Account{
		WxAccount:  req.Account,
		CreateTime: time.Now(),
	}
	//2.需要生成几个数字做为用户的唯一id  redis自增
	uid, err := a.redisDao.NextAccountId()
	if err != nil {
		return ac, biz.SqlError
	}
	ac.Uid = uid
	err = a.accountDao.SaveAccount(context.TODO(), ac)
	if err != nil {
		return ac, biz.SqlError
	}
	return ac, nil
}
