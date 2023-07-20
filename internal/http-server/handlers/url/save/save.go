package save

import (
	resp "url-shortener/internal/lib/api/response"
)

type URLSaver interface {
	Save(urlToSave, alias string) (int64, error)
}

type Request struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	resp.Response
	Alias string `json:"alias"`
}
