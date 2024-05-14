package constant

import (
	"context"
	"errors"
	"time"
)

const (

	//用于Dao中数据操作
	UserCachePrefix  = "users"
	BillCachePrefix  = "bills"
	TokenCachePrefix = "tokens"
	CacheExpiresTime = 7 * 24 * time.Hour

	//用于返回数据时消息提示
	MsgBadReqs       = "参数有误"
	MsgSuccess       = "成功"
	MsgInternalErr   = "内部错误"
	MsgDatabaseErr   = "数据库异常"
	MsgMiddleErr     = "中间件异常"
	MsgUserNotExist  = "用户不存在"
	MsgBillNotExist  = "账单不存在"
	MsgPasswordWrong = "密码错误"
	MsgTokenWrong    = "令牌错误"
	MsgHasLogin      = "已经登录"
	MsgPageOut       = "已到最后"

	//用于utils的常量
	EncodeCost       = 10
	SecretKey        = "wangjunzhenshuai"
	Issuer           = "www.caipiaotong.com"
	TokenExpiresTime = 7 * 24 * time.Hour

	//用于gin存取数据
	DUser         = "user"
	DUsername     = "username"
	DPassword     = "password"
	DPhone        = "phone"
	DNewPassword  = "newPassword"
	DNewUsername  = "newUsername"
	DToken        = "token"
	DBillType     = "type"
	DBillCost     = "cost"
	DBillName     = "name"
	DBillState    = "state"
	DBillMinCost  = "minCost"
	DBillMaxCost  = "maxCost"
	DBillPageSize = "pageSize"
	DBillPageNum  = "pageNum"
)

var (
	//错误类型
	ErrPasswordWrong = errors.New(MsgPasswordWrong)
	ErrTokenWrong    = errors.New(MsgTokenWrong)
	ErrHasLogin      = errors.New(MsgHasLogin)
	ErrUserNotExist  = errors.New(MsgUserNotExist)
	ErrPageOut       = errors.New(MsgPageOut)
	//context
	CtxTimeout = 5 * time.Second
	CtxBg      = context.Background()
)
