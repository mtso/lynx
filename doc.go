// Package lynx is a static blog site generator.
// It reads markdown files from the `content` folder
// to produce the HTML files that comprise the exported website.
// The exported files will be placed in a directory named `public`.
//
// In its current state, a template folder must be included
// that contains valid Go template files named `index.html` and `post.html`
// and a subdirectory named `css` that contains a stylesheet named `default.css`.
package lynx
