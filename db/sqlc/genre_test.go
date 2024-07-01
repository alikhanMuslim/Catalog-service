package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/alikhanMuslim/Catalog-service/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomGenre(t *testing.T) Genre {
	name := utils.RandomGenre()
	genre, err := testqueries.CreateGenre(context.Background(), name)

	require.NoError(t, err)
	require.NotEmpty(t, genre)

	require.NotZero(t, genre.ID)
	require.Equal(t, name, genre.Name)

	return genre
}

func TestCreateGenre(t *testing.T) {
	CreateRandomGenre(t)
}

func TestGetGenre(t *testing.T) {
	genre1 := CreateRandomGenre(t)
	genre2, err := testqueries.GetGenre(context.Background(), genre1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, genre2)

	require.Equal(t, genre1, genre2)
	require.Equal(t, genre1.ID, genre2.ID)
	require.Equal(t, genre1.Name, genre2.Name)

}

func TestListGenres(t *testing.T) {
	for i := 0; i < 5; i++ {
		CreateRandomGenre(t)
	}

	genres, err := testqueries.ListGenres(context.Background(), 5)

	require.NoError(t, err)
	require.NotEmpty(t, genres)
	require.Len(t, genres, 5)
}

func TestDeleteGenre(t *testing.T) {
	genre := CreateRandomGenre(t)
	err := testqueries.DeleteGenre(context.Background(), genre.ID)

	require.NoError(t, err)

	genre2, err := testqueries.GetGenre(context.Background(), genre.ID)

	require.Error(t, err)
	require.Equal(t, sql.ErrNoRows, err)
	require.Empty(t, genre2)
}

func TestUpdateGenre(t *testing.T) {
	genre := CreateRandomGenre(t)

	arg := UpdateGenreParams{
		ID:   genre.ID,
		Name: utils.RandomGenre(),
	}
	genre2, err := testqueries.UpdateGenre(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, genre2)
	require.Equal(t, genre.ID, genre2.ID)
	require.Equal(t, arg.Name, genre2.Name)
}
