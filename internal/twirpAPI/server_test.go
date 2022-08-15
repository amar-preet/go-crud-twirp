package twirpapi

import (
	"context"
	"errors"
	"go-crud-twirp/rpc/twirpAPI"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetAlbums(t *testing.T) {
	Convey("Given a call is made to GetAlbums", t, func() {
		ctx := context.TODO()
		req := &twirpAPI.GetAlbumsReq{}
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		rows := sqlmock.NewRows([]string{"id", "title", "artist", "price"}).
			AddRow(1, "test", "john", 23).
			AddRow(2, "test 2", "paul", 12)
		Convey("When a DB call is made", func() {

			mock.ExpectQuery("SELECT id, title, artist, price FROM albums").WillReturnRows(rows)

			s := &Server{DB: db}
			res, _ := s.GetAlbums(ctx, req)

			Convey("The number of rows returned should be 2", func() {
				So(len(res.Albums), ShouldEqual, 2)
				So(res, ShouldNotBeNil)
			})
		})
	})
}

func TestGetAlbumsWithError(t *testing.T) {
	Convey("Given a call is made to GetAlbums", t, func() {
		ctx := context.TODO()
		req := &twirpAPI.GetAlbumsReq{}
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		Convey("When an error is returned", func() {
			e := errors.New("No row returned")
			mock.ExpectQuery("SELECT id, title, artist, price FROM albums").WillReturnError(e)

			s := &Server{DB: db}
			_, erre := s.GetAlbums(ctx, req)

			Convey("The error is of twirp error internal", func() {
				So(erre.Error(), ShouldEqual, "twirp error internal: something went wrong")
			})
		})
	})
}
