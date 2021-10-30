package httpHandler

import (
	"context"
	"net/http"

	"github.com/Chipazawra/czwr-mailing-profile/internal/profile/model"
	"github.com/gin-gonic/gin"
)

type ReceiverHandler struct {
	usercase model.IReceiverUserCase
}

func NewReceiverHandler(uc model.IReceiverUserCase) *ReceiverHandler {
	return &ReceiverHandler{usercase: uc}
}

func (rh *ReceiverHandler) Register(g *gin.RouterGroup) *gin.RouterGroup {

	receivers := g.Group("/receivers")
	receivers.POST("/:usr/:receiver", rh.CreateHandler)
	receivers.GET("/:usr", rh.ReadHandler)
	receivers.PATCH("/:id/:receiver", rh.UpdateHandler)
	receivers.DELETE("/:id", rh.DeleteHandler)

	return receivers
}

// profile receivers godoc
// @Summary create receiver in receivers list
// @Tags profile
// @Description create receiver
// @Accept  json
// @Produce  json
// @Success 200
// @Param usr string query string true "USR"
// @Param receiver string query string true "RECEIVER"
// @Router /profile/receivers/{usr}/{receiver} [post]
func (rh *ReceiverHandler) CreateHandler(c *gin.Context) {

	usr := c.Param("usr")
	receiver := c.Param("receiver")
	id, err := rh.usercase.Create(context.TODO(), usr, receiver)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

// profile receivers godoc
// @Summary get receivers list
// @Tags profile
// @Description get receivers
// @Accept  json
// @Produce  json
// @Success 200
// @Param usr string query string true "USR"
// @Router /profile/receivers/{usr} [get]
func (rh *ReceiverHandler) ReadHandler(c *gin.Context) {

	usr := c.Param("usr")

	receivers, err := rh.usercase.Read(context.TODO(), usr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"receivers": receivers,
	})
}

// profile receivers godoc
// @Summary update receiver in receiver list
// @Tags profile
// @Description update receiver
// @Accept  json
// @Produce  json
// @Success 200
// @Param usr string query string true "USR"
// @Param id path int true "ID"
// @Param receiver string query string true "RECEIVER"
// @Router /profile/receivers/{usr}/{id}/{receiver} [patch]
func (rh *ReceiverHandler) UpdateHandler(c *gin.Context) {

	id := c.Param("id")
	usr := c.Param("usr")
	receiver := c.Param("receiver")

	_, err := rh.usercase.Update(context.TODO(), id, usr, receiver)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "updated",
	})
}

// profile receivers godoc
// @Summary delete receiver from receiver list
// @Tags profile
// @Description delete receiver
// @Accept  json
// @Produce  json
// @Success 200
// @Param usr string query string true "USR"
// @Param receiver string query string true "RECEIVER"
// @Router /profile/receivers/{id} [delete]
func (rh *ReceiverHandler) DeleteHandler(c *gin.Context) {

	id := c.Param("id")
	err := rh.usercase.Delete(context.TODO(), id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "deleted",
	})
}
