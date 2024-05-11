package constant

import "time"

const (
	CacheExpiresTime = 60 * 60 * 24 * 7 * time.Second
	SecretKey        = "wangjunzhenshuai"
	Issuer           = "www.caipiaotong.com"
	TokenExpiresTime = 604800

	UserCachePrefix = "users"
	BillCachePrefix = "bills"
)
