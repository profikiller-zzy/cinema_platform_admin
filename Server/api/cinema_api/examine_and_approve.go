package cinema_api

import (
	"AfterEnd/model/response"
	"github.com/gin-gonic/gin"
)

func (CinemaApi) ExamineAndApprove(c *gin.Context) {
	var id ID
	err := c.ShouldBindQuery(&id)
	if err != nil {
		response.FailWithCode(response.ParameterError, c)
		return
	}

}
