package user

import (
	"github.com/stretchr/testify/require"
	"go-base/fixture/database"
	"go-base/model"
	"go-base/repository"
	"go-base/usecase"
	"net/http"
	"testing"
)

func TestRoute_Login(t *testing.T) {
	db := database.InitDatabse()
	defer db.TruncateTables()

	repo := repository.New(db.GetClient())
	r := Route{useCase: usecase.New(repo)}
	t.Run("200", func(t *testing.T) {
		require.NoError(t, db.ExecFixture("../../../fixture/user/login.sql"))

		req := model.LoginReq{
			Email:    "quang@gmail.com",
			Password: "1",
		}
		rec, c := setUpTestLogin(req)

		require.NoError(t, r.Login(c))
		require.Equal(t, http.StatusOK, rec.Code)

		// remove data for the next test case
		db.TruncateTables()
	})
	t.Run("400", func(t *testing.T) {
		test := []struct {
			Name     string
			Req      model.LoginReq
			WantCode int
		}{
			{
				"Missing password",
				model.LoginReq{
					Email: "quang@gmail.com",
				},
				http.StatusBadRequest,
			},
			{
				"Missing email",
				model.LoginReq{
					Password: "1",
				},
				http.StatusBadRequest,
			},
			{
				"Invalid email",
				model.LoginReq{
					Password: "1",
					Email:    "123123",
				},
				http.StatusBadRequest,
			},
			{
				"Password incorrect",
				model.LoginReq{
					Password: "2",
					Email:    "phanngocquang2000@gmail.com",
				},
				http.StatusBadRequest,
			},
			{
				"Email incorrect",
				model.LoginReq{
					Password: "2",
					Email:    "phanngocquang@gmail.com",
				},
				http.StatusBadRequest,
			},
		}
		for _, tc := range test {
			t.Run(tc.Name, func(t *testing.T) {
				rec, c := setUpTestLogin(tc.Req)

				require.NoError(t, r.Create(c))
				require.Equal(t, tc.WantCode, rec.Code)

				// remove data for the next test case
				db.TruncateTables()
			})
		}
	})

	//t.Run("500", func(t *testing.T) {
	//	require.NoError(t, db.ExecFixture("../../../fixture/user/login_500.sql"))
	//
	//	req := model.LoginReq{
	//		Email:    "phanngocquang2000@gmail.com",
	//		Password: "1",
	//	}
	//	rec, c := setUpTestLogin(req)
	//	require.NoError(t, r.Login(c))
	//	require.Equal(t, http.StatusInternalServerError, rec.Code)
	//
	//	db.TruncateTables()
	//})
}
