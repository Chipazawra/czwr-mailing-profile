package profile

import (
	"net/http"
	"strconv"

	"github.com/Chipazawra/czwr-mailing-auth/pkg/jwtmng"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type IDBctx interface {
	Create(usr, receiver string) (uint, error)
	Read(usr string) ([]string, error)
	Update(usr string, id uint, receiver string) error
	Delete(usr string, id uint) error
}

type Profile struct {
	dbctx IDBctx
}

func New(dbctx IDBctx) *Profile {
	return &Profile{dbctx: dbctx}
}

func (r *Profile) Create(usr, receiver string) (uint, error) {

	id, err := r.dbctx.Create(usr, receiver)
	if err != nil {
		return 0, err
	}
	return id, nil

}

func (r *Profile) Read(usr string) ([]string, error) {

	lst, err := r.dbctx.Read(usr)
	if err != nil {
		return nil, err
	}
	return lst, nil

}

func (r *Profile) Update(usr string, id uint, receiver string) error {

	err := r.dbctx.Update(usr, id, receiver)
	if err != nil {
		return err
	}
	return err

}

func (r *Profile) Delete(usr string, id uint) error {

	err := r.dbctx.Delete(usr, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *Profile) Register(g *gin.Engine) *gin.RouterGroup {
	gr := g.Group("/profile")
	gr.GET("/i", p.iHandler)
	gr.GET("/me", p.meHandler)
	gr.POST("/reciviers/:usr/:receiver", p.CreateHandler)
	gr.GET("/reciviers/:usr", p.ReadHandler)
	gr.PATCH("/reciviers/:usr/:id/:receiver", p.UpdateHandler)
	gr.DELETE("/reciviers/:usr/:id", p.DeleteHandler)

	return gr
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

// profile receivers godoc
// @Summary create recievier in recieviers list
// @Tags profile
// @Description get auth data
// @Accept  json
// @Produce  json
// @Success 200
// @Router /reciviers/:usr/:receiver [post]
func (r *Profile) CreateHandler(c *gin.Context) {

	usr := c.Param("usr")
	receiver := c.Param("receiver")
	id, err := r.Create(usr, receiver)

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
// @Summary get recieviers list
// @Tags profile
// @Description get auth data
// @Accept  json
// @Produce  json
// @Success 200
// @Router /reciviers/:usr [get]
func (r *Profile) ReadHandler(c *gin.Context) {

	usr := c.Param("usr")

	receivers, err := r.Read(usr)

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
// @Summary update recievier in recievier list
// @Tags profile
// @Description get auth data
// @Accept  json
// @Produce  json
// @Success 200
// @Router /reciviers/:usr/:id/:receiver [patch]
func (r *Profile) UpdateHandler(c *gin.Context) {

	usr := c.Param("usr")
	receiver := c.Param("receiver")
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	err = r.Update(usr, uint(id), receiver)

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
// @Summary delete recievier from recievier list
// @Tags profile
// @Description get auth data
// @Accept  json
// @Produce  json
// @Success 200
// @Router /reciviers/:usr/:id [delete]
func (r *Profile) DeleteHandler(c *gin.Context) {

	usr := c.Param("usr")

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	err = r.Delete(usr, uint(id))

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
