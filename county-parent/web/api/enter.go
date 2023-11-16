package api

import (
	"github.com/jenson/county/web/api/county"
	"github.com/jenson/county/web/api/user"
)

type ApiGroup struct {
	UserApi   user.UserApi
	CountyApi county.CountyApi
}

var ApiGroupInstance = new(ApiGroup)
