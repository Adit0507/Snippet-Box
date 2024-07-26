package main

import "github.com/Adit0507/Snippet-Box/internal/models"

type templateData struct {
	Snippet *models.Snippet
	// field for holding a slice of snippets
	Snippets []*models.Snippet	
}