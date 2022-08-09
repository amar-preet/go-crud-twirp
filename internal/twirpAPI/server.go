package twirpapi

import (
	"context"
	"database/sql"
	"go-crud-twirp/db/db"
	"go-crud-twirp/rpc/twirpAPI"

	"github.com/twitchtv/twirp"
)

type Server struct {
	DB *sql.DB
}

func (s *Server) GetAlbums(ctx context.Context, req *twirpAPI.GetAlbumsReq) (*twirpAPI.GetAlbumsResp, error) {
	var albums []*twirpAPI.Album
	queries := db.New(s.DB)
	rows, err := queries.GetAlbums(ctx)
	if err != nil {
		return nil, twirp.WrapError(twirp.NewError(twirp.Internal, "something went wrong"), err)
	}
	for _, v := range rows {
		album := &twirpAPI.Album{
			Id:     v.ID,
			Title:  v.Title.String,
			Artist: v.Artist.String,
			Price:  v.Price.Int32,
		}

		albums = append(albums, album)
	}

	return &twirpAPI.GetAlbumsResp{
		Albums: albums,
	}, nil
}

func (s *Server) DeleteAlbumByID(ctx context.Context, req *twirpAPI.DeleteAlbumByIDReq) (*twirpAPI.DeleteAlbumByIDResp, error) {
	queries := db.New(s.DB)
	err := queries.DeleteAlbumByID(ctx, int32(req.Id))

	if err != nil {
		return nil, twirp.WrapError(twirp.NewError(twirp.Internal, "Error deleting an album"), err)
	}

	return &twirpAPI.DeleteAlbumByIDResp{
		Success: true,
	}, nil
}

func (s *Server) GetAlbumByID(ctx context.Context, req *twirpAPI.GetAlbumByIDReq) (*twirpAPI.GetAlbumByIDResp, error) {
	queries := db.New(s.DB)
	a, err := getAlbumByID(queries, ctx, int32(req.Id))

	if err != nil {
		return nil, twirp.WrapError(twirp.NewError(twirp.NotFound, "Error getting an album by ID"), err)
	}

	return &twirpAPI.GetAlbumByIDResp{
		Album: &twirpAPI.Album{Id: a.ID, Artist: a.Artist.String, Title: a.Title.String, Price: a.Price.Int32},
	}, nil
}

func (s *Server) PostAlbums(ctx context.Context, req *twirpAPI.PostAlbumsReq) (*twirpAPI.PostAlbumsResp, error) {
	a := db.PostAlbumsParams{
		Title:  sql.NullString{String: req.Title, Valid: true},
		Artist: sql.NullString{String: req.Artist, Valid: true},
		Price:  sql.NullInt32{Int32: req.Price, Valid: true},
	}
	queries := db.New(s.DB)
	id, err := queries.PostAlbums(ctx, a)
	if err != nil {
		return nil, twirp.WrapError(twirp.NewError(twirp.Internal, "Error posting an album"), err)
	}

	return &twirpAPI.PostAlbumsResp{
		Album: &twirpAPI.Album{Id: int32(id), Artist: req.Artist, Title: req.Title, Price: req.Price},
	}, nil
}

func (s *Server) UpdateAlbumByID(ctx context.Context, req *twirpAPI.UpdateAlbumByIDReq) (*twirpAPI.UpdateAlbumByIDResp, error) {
	a := db.UpdateAlbumByIDParams{
		ID:     req.Id,
		Title:  sql.NullString{String: req.Title, Valid: true},
		Artist: sql.NullString{String: req.Artist, Valid: true},
		Price:  sql.NullInt32{Int32: req.Price, Valid: true},
	}
	queries := db.New(s.DB)
	_, err := getAlbumByID(queries, ctx, int32(req.Id))

	if err != nil {
		return nil, twirp.WrapError(twirp.NewError(twirp.NotFound, "Error getting an album by ID"), err)
	}

	album, err := queries.UpdateAlbumByID(ctx, a)

	if err != nil {
		return nil, twirp.WrapError(twirp.NewError(twirp.Internal, "Error updating an album"), err)
	}

	return &twirpAPI.UpdateAlbumByIDResp{
		Album: &twirpAPI.Album{Id: int32(album.ID), Artist: album.Artist.String, Title: album.Title.String, Price: album.Price.Int32},
	}, nil
}

func getAlbumByID(queries *db.Queries, ctx context.Context, id int32) (db.Album, error) {
	return queries.GetAlbumByID(ctx, id)
}
