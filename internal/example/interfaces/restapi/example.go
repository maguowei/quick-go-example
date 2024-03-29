package restapi

import (
    "github.com/gin-gonic/gin"
    "example/internal/example/application/service"
    "example/internal/example/interfaces/dto"
    "example/internal/pkg/restresp"
    "net/http"
)

type ExampleApi struct {
    exampleAppService service.ExampleAppServiceInterface
}

func NewExampleApi(exampleAppService service.ExampleAppServiceInterface) *ExampleApi {
    return &ExampleApi{
        exampleAppService: exampleAppService,
    }
}

func (api *ExampleApi) CreateExample(c *gin.Context) {
    var req dto.ExampleCreateReqDto
    if err := c.ShouldBindJSON(&req); err != nil {
        resp := restresp.DefaultErrWithMsgResp(c, "参数错误!")
        c.JSON(http.StatusBadRequest, resp)
        return
    }

    _, err := api.exampleAppService.CreateExample(c.Request.Context(), req.ToDO())
    if err != nil {
        resp := restresp.DefaultErrWithMsgResp(c, err.Error())
        c.JSON(http.StatusOK, resp)
        return
    }

    respDto := dto.ExampleCreateRespDto{}
    c.JSON(http.StatusOK, restresp.SuccessWithDataResp(c, respDto))
}
