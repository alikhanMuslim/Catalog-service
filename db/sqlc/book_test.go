package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/alikhanMuslim/Catalog-service/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomBook(t *testing.T, author Author, genre Genre) Book {
	arg := CreateBookParams{
		Title:     utils.RandomTitle(),
		AuthorID:  author.ID,
		GenreID:   genre.ID,
		Price:     int64(utils.RandomPrice(700, 1000)),
		Available: utils.RandomBool(),
	}

	book, err := testqueries.CreateBook(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, book)

	require.NotZero(t, book.ID)
	require.NotZero(t, book.CreatedAt)

	require.Equal(t, arg.AuthorID, book.AuthorID)
	require.Equal(t, arg.GenreID, book.GenreID)
	require.Equal(t, arg.Price, book.Price)
	require.Equal(t, arg.Available, book.Available)

	return book
}

func TestCreateBook(t *testing.T) {
	author := CreateRandomAuthor(t)
	genre := CreateRandomGenre(t)

	CreateRandomBook(t, author, genre)
}

func TestDeleteBook(t *testing.T) {
	author := CreateRandomAuthor(t)
	genre := CreateRandomGenre(t)

	book := CreateRandomBook(t, author, genre)

	err := testqueries.DeleteBook(context.Background(), book.ID)

	require.NoError(t, err)

	book2, err := testqueries.GetBook(context.Background(), book.ID)

	require.Error(t, err)
	require.Empty(t, book2)

	require.Equal(t, sql.ErrNoRows, err)
}

func TestGetBook(t *testing.T) {
	author := CreateRandomAuthor(t)
	genre := CreateRandomGenre(t)

	book := CreateRandomBook(t, author, genre)

	book2, err := testqueries.GetBook(context.Background(), book.ID)

	require.NoError(t, err)
	require.NotEmpty(t, book2)
	require.Equal(t, book.ID, book2.ID)
	require.Equal(t, book.AuthorID, book2.AuthorID)
	require.Equal(t, book.Available, book2.Available)
	require.WithinDuration(t, book.CreatedAt, book2.CreatedAt, time.Second)
	require.Equal(t, book.GenreID, book2.GenreID)
	require.Equal(t, book.Title, book2.Title)
	require.Equal(t, book.Price, book2.Price)

}

func TestListBooks(t *testing.T) {
	author := CreateRandomAuthor(t)
	genre := CreateRandomGenre(t)

	for i := 0; i < 10; i++ {
		CreateRandomBook(t, author, genre)
	}

	arg := ListBooksParams{
		Limit:  5,
		Offset: 5,
	}

	books, err := testqueries.ListBooks(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, books)
	require.Len(t, books, 5)
}

func TestUpdateBook(t *testing.T) {
	author := CreateRandomAuthor(t)
	genre := CreateRandomGenre(t)

	book := CreateRandomBook(t, author, genre)

	arg := UpdateBookParams{
		ID:        book.ID,
		Price:     int64(utils.RandomPrice(400, 500)),
		Available: utils.RandomBool(),
	}

	bookUpdated, err := testqueries.UpdateBook(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, bookUpdated)

	require.Equal(t, book.ID, bookUpdated.ID)
	require.Equal(t, book.AuthorID, bookUpdated.AuthorID)
	require.Equal(t, book.GenreID, bookUpdated.GenreID)
	require.Equal(t, book.Title, bookUpdated.Title)
	require.WithinDuration(t, book.CreatedAt, bookUpdated.CreatedAt, time.Second)
	require.Equal(t, arg.Available, bookUpdated.Available)
	require.Equal(t, arg.Price, bookUpdated.Price)

}
