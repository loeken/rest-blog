package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/loeken/rest-blog/api/auth"
	"github.com/loeken/rest-blog/api/models"
	"github.com/loeken/rest-blog/api/responses"
	"github.com/loeken/rest-blog/api/utils/formaterror"
)

// Login godoc
// @Summary Login to the application, generates a JWT
// @Description Allows login to the application by generating a JWT
// @Accept json
// @Param body body models.User true "Login"
// @Success 200 {string} string	"Token"
// @tags auth
// @Router /auth [post]
func (server *Server) Login(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}

	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()

	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, "Bearer "+token)
}

func (server *Server) SignIn(email, password string) (string, error) {

	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)

	if err != nil {
		return "", err
	}
	return auth.CreateToken(user.ID)
}
