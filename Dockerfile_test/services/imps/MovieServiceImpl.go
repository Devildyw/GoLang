package imps

import (

	//"Go_Project/GO_LEARN/Day08/iris/myMVCapp/datasource"

	"Dockerfile_test/datamodels"
	"Dockerfile_test/datasource"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MovieServiceImpl struct{

}
var db *sql.DB
func init(){
	db = datasource.DataSourceInit()
}


func (s *MovieServiceImpl) GetAll() []datamodels.Movie {
	result, err2 := db.Query("select * from movie")
	if err2 != nil {
		log.Println("DBConn.QueryContext err2 = ", err2)
	}
	movies := make([]datamodels.Movie, 0, 10)
	for result.Next() {
		var movie datamodels.Movie
		result.Scan(&movie.ID, &movie.Name, &movie.Genre, &movie.Poster, &movie.Year)
		movies = append(movies, movie)
	}

	return movies
}

func (s *MovieServiceImpl) GetById(id int64) (datamodels.Movie, bool) {
	result := db.QueryRow("select * from movie where id = ?", id)
	var movie datamodels.Movie
	if result != nil {
		result.Scan(&movie.ID, &movie.Name, &movie.Genre, &movie.Poster, &movie.Year)
		return movie, true
	}
	return movie, false

}

func (s *MovieServiceImpl) DeleteById(id int64) bool {
	result, err := db.Exec("delete from movie where id = ?", id)
	if err != nil {
		log.Println(err)
		return false
	}
	RowsAffected, err2 := result.RowsAffected()
	if err2 != nil {
		log.Println(err2)
	}
	if RowsAffected < 0 {
		return false
	}
	return true
}

func (s *MovieServiceImpl) UpdatePosterAndGenreById(id int64, poster string, genre string) (datamodels.Movie, error) {
	result, err := db.Exec("update movie set poster = ?, genre = ? where id = ?", id)
	if err != nil {
		log.Println(err)
	}
	var movie datamodels.Movie
	if rowa, _ := result.RowsAffected(); rowa > 0 {
		result2 := db.QueryRow("select * from movie where id = ?", id)
		if result2 != nil {
			result2.Scan(&movie.ID, &movie.Name, &movie.Genre, &movie.Poster, &movie.Year)
		}
	}

	return movie, err
}
