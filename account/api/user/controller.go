package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smartbot/account/pkg/validator"
)

type UserController struct {
	service UserService
}

func (uc UserController) OnboardUser(c *gin.Context) {

	log.Print("under onboard.......")
	var onboardRequest OnboardRequest
	err := validator.ValidateBody(c, &onboardRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	userId, _ := c.Get("user_id")
	userName, _ := c.Get("username")

	log.Println("onboard valid success")

	response, onboardError := uc.service.OnboardUser(userId.(string), userName.(string), onboardRequest)
	if onboardError != nil {
		c.JSON(onboardError.Code, onboardError)
		return
	}

	c.JSON(http.StatusCreated, response)

}

func (uc UserController) GetUser(c *gin.Context) {

	value, _ := c.Get("user_id")

	res, err := uc.service.GetUser(value.(string))

	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, res)
}
