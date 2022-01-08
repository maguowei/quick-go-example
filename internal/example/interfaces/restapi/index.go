package restapi

import (
	"github.com/gin-gonic/gin"
	"github.com/maguowei/example/internal/pkg/restresp"
	"github.com/maguowei/example/internal/example/interfaces/dto"
	"net/http"
)


func Health(c *gin.Context) {
	data := dto.HealthRespDto{Status: "ok"}
	resp := restresp.SuccessWithDataResp(data)
	c.JSON(http.StatusOK, resp)
}

func Index(c *gin.Context) {
	resp := restresp.DefaultSuccessResp()
	c.JSON(http.StatusOK, resp)
}