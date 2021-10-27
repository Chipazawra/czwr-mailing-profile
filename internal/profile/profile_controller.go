package profile

import (
	"context"
	"net/http"

	"github.com/Chipazawra/czwr-mailing-auth/pkg/jwtmng"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (p *Profile) pHandler(c *gin.Context, fn func(val string) (jwt.Claims, error)) {

	ac, err := c.Request.Cookie("access")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "access token not found.",
		})
		return
	}

	data, err := fn(ac.Value)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "invalid token",
		})
		return
	} else {
		c.JSON(http.StatusOK, data)
	}
}

// profile i godoc
// @Summary show user info
// @Tags profile
// @Description get auth data
// @Accept  json
// @Produce  json
// @Success 200
// @Router /profile/i [get]
func (p *Profile) iHandler(c *gin.Context) {
	p.pHandler(c, jwtmng.ParseToken)
}

// profile me godoc
// @Summary valid jwt token and show user info
// @Tags profile
// @Description get auth data
// @Accept  json
// @Produce  json
// @Success 200
// @Router /profile/me [get]
func (p *Profile) meHandler(c *gin.Context) {
	p.pHandler(c, jwtmng.ValidToken)
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
// @Router /profile/reciviers/{usr}/{receiver} [post]
func (p *Profile) CreateHandler(c *gin.Context) {

	usr := c.Param("usr")
	receiver := c.Param("receiver")
	id, err := p.receivers.Create(context.TODO(), usr, receiver)

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
// @Router /profile/reciviers/{usr} [get]
func (p *Profile) ReadHandler(c *gin.Context) {

	usr := c.Param("usr")

	receivers, err := p.receivers.Read(context.TODO(), usr)

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
// @Router /profile/reciviers/{usr}/{id}/{receiver} [patch]
func (r *Profile) UpdateHandler(c *gin.Context) {

	receiver := c.Param("receiver")
	id := c.Param("id")
	err := r.receivers.Update(context.TODO(), id, receiver)

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
// @Router /profile/reciviers/{usr}/{id} [delete]
func (r *Profile) DeleteHandler(c *gin.Context) {

	usr := c.Param("usr")
	id := c.Param("id")
	err := r.receivers.Delete(context.TODO(), usr, id)

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

// profile receivers godoc
// @Summary upload template data
// @Tags profile
// @Description upload data
// @Accept  json
// @Produce  json
// @Success 200
// @Router /profile/upload_template [delete]
func (r *Profile) UploadTemplateHandler(c *gin.Context) {

	raw, err := c.GetRawData()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	if res, err := r.template.Create(context.TODO(), string(raw)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, res)
	}

}
