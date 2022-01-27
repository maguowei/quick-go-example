package restapi

import (
    "github.com/gin-gonic/gin"
    "github.com/maguowei/example/internal/example/application/service"
    "github.com/maguowei/example/internal/example/interfaces/dto"
    "github.com/maguowei/example/internal/pkg/restresp"
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
        resp := restresp.DefaultErrWithMsgResp("参数错误!")
        c.JSON(http.StatusBadRequest, resp)
        return
    }

    _, err := api.exampleAppService.CreateExample(req.ToDO())
    if err != nil {
        resp := restresp.DefaultErrWithMsgResp(err.Error())
        c.JSON(http.StatusOK, resp)
        return
    }

    respDto := dto.ExampleCreateRespDto{}
    c.JSON(http.StatusOK, restresp.SuccessWithDataResp(respDto))
}