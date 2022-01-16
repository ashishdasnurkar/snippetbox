package main

import "ashishdasnurkar.com/snippetbox/pkg/models"

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}