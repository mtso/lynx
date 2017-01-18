package lynx

// Configuration contains the Title and Description of the website.
// It is passed into GenerateWith to build a website that uses
// the configuration properties.
type Configuration struct {

	// Title is a string of the Index page.
	Title string

	// Description is a string describing the generated website.
	// HTML tags included in here will be parsed during site generation.
	Description string

	ExportDir *string
}
