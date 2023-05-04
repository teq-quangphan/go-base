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

func TestRoute_Create(t *testing.T) {
	db := database.InitDatabse()
	defer db.TruncateTables()

	repo := repository.New(db.GetClient())
	r := Route{useCase: usecase.New(repo)}

	//test
	t.Run("201_Success", func(t *testing.T) {
		req := model.User{
			UserName: "quangpn",
			Password: "1",
			Email:    "quang@gmail.com",
		}
		rec, c := setUpTestCreateUser(req)

		require.NoError(t, r.Create(c))
		require.Equal(t, http.StatusCreated, rec.Code)

		// remove data for the next test case
		db.TruncateTables()
	})

	t.Run("400", func(t *testing.T) {
		test := []struct {
			Name     string
			Req      model.User
			WantCode int
		}{
			{
				"Missing password",
				model.User{
					UserName: "quangpn",
					Email:    "quang@gmail.com",
				},
				http.StatusBadRequest,
			},
			{
				"Missing email",
				model.User{
					UserName: "quangpn",
					Password: "1",
				},
				http.StatusBadRequest,
			},
			{
				"Invalid email",
				model.User{
					UserName: "quangpn",
					Password: "1",
					Email:    "123123",
				},
				http.StatusBadRequest,
			},
		}
		for _, tc := range test {
			t.Run(tc.Name, func(t *testing.T) {
				rec, c := setUpTestCreateUser(tc.Req)

				require.NoError(t, r.Create(c))
				require.Equal(t, tc.WantCode, rec.Code)

				// remove data for the next test case
				db.TruncateTables()
			})
		}
	})
}
