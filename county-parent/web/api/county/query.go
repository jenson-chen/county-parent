package county

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jenson/county/core/e"
	"github.com/jenson/county/core/models"
	"github.com/jenson/county/web/global"
	"time"
)

func (ca CountyApi) QueryDetailByEditor(c *gin.Context) {
	var req models.EditorRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		e.FailWithMessage("必要参数未输入或数据类型错误", c)
		return
	}

	ctx, cancelFunc := context.WithTimeout(c.Request.Context(), 1*time.Second)
	defer cancelFunc()

	client := global.CountyServiceClient
	detail, err := client.QueryDetailByEditor(ctx, &req)
	if err != nil {
		e.FailWithMessage(err.Error(), c)
		return
	}
	e.OKWithData(detail, c)

}
