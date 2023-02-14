package restapi

import (
	"github.com/gin-gonic/gin"
	"example/internal/example/interfaces/dto"
	"example/internal/pkg/restresp"
	"net/http"
)


func Health(c *gin.Context) {
	data := dto.HealthRespDto{Status: "ok"}
	resp := restresp.SuccessWithDataResp(c, data)
	c.JSON(http.StatusOK, resp)
}

func Index(c *gin.Context) {
	resp := restresp.DefaultSuccessResp(c)
	c.JSON(http.StatusOK, resp)
}
