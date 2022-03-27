package controlleres

import (
	"Dockerfile_test/services"
	"Dockerfile_test/services/imps"
	"Dockerfile_test/web/response"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type MovieController struct {
	ctx iris.Context
	Service services.MovieService
}

func (c *MovieController) GetAllMovie() (result response.Result){
	return result.Success(c.Service.GetAll())
}

func (c *MovieController) GetMovieById(id string) (result response.Result) {
	i, _ := strconv.ParseInt(id, 10, 64)
	m, b := c.Service.GetById(i)
	if b {
		return result.Success(m)
	} else {
		return result.Fail(900, "查询失败")
	}
}
func (c *MovieController) DeleteMovieById(id int64) (result response.Result) {
	b := c.Service.DeleteById(id)
	if b {
		return result.Success(nil)
	} else {
		return result.Fail(901, "删除失败")
	}
}

func (c *MovieController) UpdateMoviePosterAndGenreById(id int64, poster string, genre string) (result response.Result) {
	m, _ := c.Service.UpdatePosterAndGenreById(id, poster, genre)
	return result.Success(m)
}
func NewMovieController() *MovieController{
	return &MovieController{
		Service: &imps.MovieServiceImpl{},
	}
}
func (c *MovieController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(
		"GET",
		"movie",
		"GetAllMovie",
	)
	b.Handle(
		"GET",
		"/movie/{id}",
		"GetMovieById",
	)
	b.Handle(
		"DELETE",
		"/movie/delete/{id}",
		"DeleteMovieById",
	)
	b.Handle(
		"PUT",
		"/movie/delete/{id}",
		"UpdateMoviePosterAndGenreById",
	)
	b.Handle(
		"GET",
		"/id",
		"GetID",
	)
}
func (c *MovieController) GetID() (result response.Result) {
	return result.Success("dfas")
}
