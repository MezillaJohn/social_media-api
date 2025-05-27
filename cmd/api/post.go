package main

import (
	"fmt"
	"net/http"

	"errors"

	"github.com/MezillaJohn/social_media-api/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type createPostType struct {
	Content string   `json:"content"`
	Title   string   `json:"title"`
	Tags    []string `json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse("077e7f50-9bf4-4969-b896-d4429de312d8")
	if err != nil {
		app.badRequestError(w, r, err)
		return
	}

	var payload createPostType

	if err := readJson(w, r, &payload); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	post := &store.Post{
		Content: payload.Content,
		Title:   payload.Title,
		UserID:  id,
		Tags:    payload.Tags,
	}

	ctx := r.Context()
	if err := app.store.Posts.Create(ctx, post); err != nil {
		fmt.Println("DB Error:", err) // This prints the actual DB error
		app.internalServerError(w, r, err)
		return
	}

	if err := writeJson(w, http.StatusCreated, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}

}

func (app *application) getPostById(w http.ResponseWriter, r *http.Request) {
	postIdStr := chi.URLParam(r, "postId")

	postId, err := uuid.Parse(postIdStr)

	if err != nil {
		app.badRequestError(w, r, err)
	}

	ctx := r.Context()

	post, err := app.store.Posts.GetPostById(ctx, postId)

	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundError(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	if err := writeJson(w, http.StatusOK, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
