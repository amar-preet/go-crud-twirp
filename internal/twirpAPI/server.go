package twirpapi

import (
	"context"
	"database/sql"
	"go-crud-twirp/rpc/twirpAPI"

	"github.com/twitchtv/twirp"
)

type Server struct {
	DB *sql.DB
}

func (s *Server) GetAlbums(ctx context.Context, req *twirpAPI.GetAlbumsReq) (*twirpAPI.GetAlbumsResp, error) {
	var albums []*twirpAPI.Album

	rows, err := s.DB.Query("SELECT id, title, artist, price FROM albums;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c twirpAPI.Album
		err = rows.Scan(&c.Id, &c.Artist, &c.Title, &c.Price)
		if err != nil {
			return nil, twirp.WrapError(twirp.NewError(twirp.Internal, "something went wrong"), err)
		}
		albums = append(albums, &c)
	}

	return &twirpAPI.GetAlbumsResp{
		Albums: albums,
	}, nil
}

func (s *Server) DeleteAlbumByID(ctx context.Context, req *twirpAPI.DeleteAlbumByIDReq) (*twirpAPI.DeleteAlbumByIDResp, error) {
	sqlStatement := `DELETE FROM albums
	WHERE id=$1`
	_, err := s.DB.Exec(sqlStatement, req.Id)
	if err != nil {
		return nil, twirp.WrapError(twirp.NewError(twirp.Internal, "Error deleting an album"), err)
	}

	return &twirpAPI.DeleteAlbumByIDResp{
		Success: true,
	}, nil
}

func (s *Server) GetAlbumByID(ctx context.Context, req *twirpAPI.GetAlbumByIDReq) (*twirpAPI.GetAlbumByIDResp, error) {
	var c twirpAPI.Album

	sqlStatement := `SELECT id, title, artist, price 
	FROM albums 
	WHERE id=$1`
	row := s.DB.QueryRow(sqlStatement, req.Id)
	err := row.Scan(&c.Id, &c.Title, &c.Artist, &c.Price)
	if err != nil {
		return nil, twirp.WrapError(twirp.NewError(twirp.Internal, "Error getting an album by ID"), err)
	}

	return &twirpAPI.GetAlbumByIDResp{
		Album: &twirpAPI.Album{Id: c.Id, Artist: c.Artist, Title: c.Title, Price: c.Price},
	}, nil
}

func (s *Server) PostAlbums(ctx context.Context, req *twirpAPI.PostAlbumsReq) (*twirpAPI.PostAlbumsResp, error) {
	id := 0
	sqlStatement := `INSERT INTO albums (title, artist, price)
	VALUES ($1, $2, $3)
	RETURNING id`
	err := s.DB.QueryRow(sqlStatement, req.Title, req.Artist, req.Price).Scan(&id)
	if err != nil {
		return nil, twirp.WrapError(twirp.NewError(twirp.Internal, "Error posting an album"), err)
	}

	return &twirpAPI.PostAlbumsResp{
		Album: &twirpAPI.Album{Id: int32(id), Artist: req.Artist, Title: req.Title, Price: req.Price},
	}, nil
}

func (s *Server) UpdateAlbumByID(ctx context.Context, req *twirpAPI.UpdateAlbumByIDReq) (*twirpAPI.UpdateAlbumByIDResp, error) {
	sqlStatement := `UPDATE albums
	SET title = $2, artist = $3, price = $4 
	WHERE id=$1`
	_, err := s.DB.Exec(sqlStatement, req.Id, req.Title, req.Artist, req.Price)
	if err != nil {
		return nil, twirp.WrapError(twirp.NewError(twirp.Internal, "Error updating an album"), err)
	}

	return &twirpAPI.UpdateAlbumByIDResp{
		Album: &twirpAPI.Album{Id: int32(req.Id), Artist: req.Artist, Title: req.Title, Price: req.Price},
	}, nil
}
