package main

import "github.com/zerobl21/letsgo/internal/models"

type templateData struct {
  Snippet *models.Snippet
  Snippets []*models.Snippet
}
