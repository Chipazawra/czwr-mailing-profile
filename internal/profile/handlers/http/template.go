package httpHandler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Chipazawra/czwr-mailing-profile/internal/profile/model"
	"github.com/gin-gonic/gin"
)

type TemplateHandler struct {
	usercase model.ITemplatesUserCase
}

func NewTemplateHandler(uc model.ITemplatesUserCase) *TemplateHandler {
	return &TemplateHandler{usercase: uc}
}

func (th *TemplateHandler) Register(g *gin.RouterGroup) *gin.RouterGroup {
	templates := g.Group("/templates")
	templates.POST("/upload", th.UploadTemplateHandler)

	return templates
}

// profile receivers godoc
// @Summary upload template data
// @Tags profile
// @Description upload data
// @Accept  json
// @Produce  json
// @Success 200
// @Router /profile/upload_template [delete]
func (th *TemplateHandler) UploadTemplateHandler(c *gin.Context) {

	raw, err := c.GetRawData()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": fmt.Errorf("TemplateHandler: %v", err),
		})
		return
	}

	if res, err := th.usercase.UploadTemplate(context.TODO(), string(raw)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": fmt.Errorf("TemplateHandler: %v", err),
		})
		return
	} else {
		c.JSON(http.StatusOK, res)
	}

}
