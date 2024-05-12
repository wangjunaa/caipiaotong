package constant

import (
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
	MsgBadReqs     = "参数有误"
	MsgSuccess     = "成功"
	MsgInternalErr = "内部错误"
	MsgDatabaseErr = "数据库异常"
	MsgMiddleErr   = "中间件异常"

	//用于utils的常量
	EncodeCost       = 10
	SecretKey        = "wangjunzhenshuai"
	Issuer           = "www.caipiaotong.com"
	TokenExpiresTime = 7 * 24 * time.Hour

	//用于context的数据操作
	CData = "Data"
)

var (
	//错误类型
	ErrPasswordWrong = errors.New("密码错误")
	ErrTokenWrong    = errors.New("令牌错误")
)