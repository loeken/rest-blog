package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/loeken/rest-blog/api/auth"
	"github.com/loeken/rest-blog/api/models"
	"github.com/loeken/rest-blog/api/responses"
	"github.com/loeken/rest-blog/api/utils/formaterror"
)
// CreatePost godoc
// @Summary Creates a new post
// @Description Creates a new post
// @Accept json
// @Param body body models.Post true "Json body containing post data"
// @Success 200 {array} models.Post
// @Router /post [post]
// @Security ApiKeyAuth
// @param Authorization header string true "Format: 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MDQ1MzExMDIsInVzZXJfaWQiOjF9.O63ZS_Poy29dDdcZqHDN0XeMPbYPX-Vfyl_FPfsMTvQ'"
// @tags posts
func (server *Server) CreatePost(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	post := models.Post{}
	err = json.Unmarshal(body, &post)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	post.Prepare()
	err = post.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if uid != post.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	postCreated, err := post.SavePost(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, postCreated.ID))
	responses.JSON(w, http.StatusCreated, postCreated)
}
// GetPosts godoc
// @Summary lists all posts
// @Description lists all posts
// @Accept  json
// @Success 200 {array} models.Post
// @Produces string
// @Router /post [get]
// @tags posts
func (server *Server) GetPosts(w http.ResponseWriter, r *http.Request) {

	post := models.Post{}

	posts, err := post.FindAllPosts(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, posts)
}
// GetPost godoc
// @Summary list post by id
// @Description list post by id
// @Accept  json
// @Success 200 {object} models.Post
// @Param id path int true "Post ID"
// @Produces string
// @Router /post/{id} [get]
// @tags posts
func (server *Server) GetPost(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	post := models.Post{}

	postReceived, err := post.FindPostByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, postReceived)
}
// UpdatePost godoc
// @Summary Updates a post
// @Description Updates a post
// @Accept json
// @Param body body models.Post true "Json body containing post data"
// @Success 200 {object} models.Post
// @Router /post/{id} [put]
// @Security ApiKeyAuth
// @param Authorization header string true "Format: 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MDQ1MzExMDIsInVzZXJfaWQiOjF9.O63ZS_Poy29dDdcZqHDN0XeMPbYPX-Vfyl_FPfsMTvQ'"
// @Param id path int true "Post ID"
// @tags posts
func (server *Server) UpdatePost(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	post := models.Post{}
	err = json.Unmarshal(body, &post)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	post.Prepare()
	err = post.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if uid != post.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	postUpdated, err := post.UpdateAPost(server.DB, pid)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, postUpdated)
}
// DeletePost godoc
// @Summary Deletes a post
// @Description Deletes a post
// @Accept json
// @Success 200 {string} string	"Entity Count"
// @Router /post/{id} [delete]
// @Param id path int true "Post ID"
// @Security ApiKeyAuth
// @param Authorization header string true "Format: 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MDQ1MzExMDIsInVzZXJfaWQiOjF9.O63ZS_Poy29dDdcZqHDN0XeMPbYPX-Vfyl_FPfsMTvQ'"
// @tags posts
func (server *Server) DeletePost(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	post := models.Post{}

	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	_, err = post.DeleteAPost(server.DB, pid, uid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent, "")
}
