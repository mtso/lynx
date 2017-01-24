package lynx

import (
	"os"
	"path/filepath"

	"github.com/mtso/lynx/bolo"
)

const (
	IndexTemplateName = "index.html"
	defaultExportDir  = "public"

	// For use later:
	defaultContentDir  = "content"
	defaultTemplateDir = "template"
)

// GenerateWith reads files from content and template directories to export
// a public directory containing compiled static site files.
// Takes a Configuration containing Title and Description as an argument.
func GenerateWith(config Configuration) (err error) {

	var exportDir string
	if configDir := config.ExportDir; configDir != nil {
		exportDir = *configDir
	} else {
		exportDir = defaultExportDir
	}

	// Delete to refresh export folder
	err = os.RemoveAll(exportDir)

	// Make public directory to store files
	if err = mkdirIfNone(exportDir); err != nil {
		return
	}

	//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!//
	// TESTING OUT TEMPLATE PACKAGE //
	//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!//
	mkdirIfNone("./public/css")
	writeStringToFile(lynxtemplate.Css[0], "./public/css/tester.css")

	// Copy static files
	err = copyFromTo("./template/css/default.css", "./public/css/style.css")

	// Load content pages
	pages, err := loadPagesIn("content")
	if err != nil {
		return
	}

	path := filepath.Join("template", "post.html")
	if err = pages.loadTemplate(path); err != nil {
		return
	}

	// Sort pages in reverse chronological order
	pages = pages.reverseChronological()
	pages.relinkNext()

	// Execute template
	pages.executeTemplate()

	// Save content pages
	err = pages.exportTo(exportDir)

	// Generate Index Page
	Index := newIndex(config.Title, config.Description, pages)
	IndexTemplatePath := filepath.Join("template", IndexTemplateName)
	if err = Index.loadTemplate(IndexTemplatePath); err != nil {
		return
	}

	if err = Index.executeTemplate(); err != nil {
		return
	}

	if err = Index.writeTo(exportDir); err != nil {
		return
	}

	return
}

// Generate calls GenerateWith using a default Configuration object of
// {Title: "Blog", Description: "Blog description."}
func Generate() (err error) {
	return GenerateWith(Configuration{
		Title:       "Blog",
		Description: "Blog description.",
	})
}

// Makes a directory if none exists
func mkdirIfNone(dirname string) error {
	err := os.MkdirAll(dirname, os.ModePerm)
	if err == os.ErrInvalid || err == os.ErrPermission {
		return err
	}
	return nil
}
