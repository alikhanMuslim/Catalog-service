package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/alikhanMuslim/Catalog-service/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomAuthor(t *testing.T) Author {

	arg := CreateAuthorParams{
		Name: utils.RandomName(),
		Bio:  utils.RandomBio(),
	}

	author, err := testqueries.CreateAuthor(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, author)

	require.NotZero(t, author.ID)
	require.Equal(t, arg.Name, author.Name)
	require.Equal(t, arg.Bio, author.Bio)

	return author
}

func TestCreateAuthor(t *testing.T) {
	CreateRandomAuthor(t)
}

func TestDeleteAuthor(t *testing.T) {
	author := CreateRandomAuthor(t)

	err := testqueries.DeleteAuthor(context.Background(), author.ID)

	require.NoError(t, err)

	author2, err := testqueries.GetAuthor(context.Background(), author.ID)

	require.Error(t, err)
	require.Equal(t, sql.ErrNoRows, err)
	require.Empty(t, author2)
}

func TestGetAuthor(t *testing.T) {
	author := CreateRandomAuthor(t)

	author2, err := testqueries.GetAuthor(context.Background(), author.ID)

	require.NoError(t, err)
	require.NotEmpty(t, author2)

	require.Equal(t, author, author2)
	require.Equal(t, author.ID, author2.ID)
	require.Equal(t, author.Name, author2.Name)
	require.Equal(t, author.Bio, author2.Bio)

}

func TestListAuthors(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomAuthor(t)
	}

	arg := ListAuthorsParams{
		Limit:  5,
		Offset: 5,
	}

	authors, err := testqueries.ListAuthors(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, authors)
	require.Len(t, authors, 5)
}

func TestUpdateAuthor(t *testing.T) {
	author := CreateRandomAuthor(t)

	arg := UpdateAuthorParams{
		ID:  author.ID,
		Bio: utils.RandomBio(),
	}

	author2, err := testqueries.UpdateAuthor(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, author2)

	require.Equal(t, author.ID, author2.ID)
	require.Equal(t, author.Name, author2.Name)
	require.Equal(t, arg.Bio, author2.Bio)
}
