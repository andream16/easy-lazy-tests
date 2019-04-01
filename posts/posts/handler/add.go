package handler

import (
	"net/http"

	"github.com/andream16/personal-go-projects/posts/posts"
)

func (h *Handler) add(r *http.Request, w http.ResponseWriter) {

	var post posts.Post

	if error := h.serializer.Deserialize(r.Body, &post); error != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !post.Valid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newPost, err := h.service.Add(post)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, error := h.serializer.Serialize(newPost)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(b)

}
