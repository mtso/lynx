# lynx [![Build Status](https://travis-ci.org/mtso/lynx.svg?branch=master)](https://travis-ci.org/mtso/lynx) [![Coverage Status][codecov-badge]][codecov]

Static blog site generator.

## Roadmap

0. [x] Load templates into memory
0. [x] Load post content into memory
0. [x] Execute content on HTML templates
0. [x] Save generated HTML to public folder
0. [x] Get content file creation time
0. [x] Parse YAML front-matter in `content/`
0. [ ] Copy static assets into public folder
0. [ ] Add default template to lynx package (probably as a subdirectory)
0. [ ] Conform templates into go packages (only one location to update)
0. [ ] Add documentation on how to get started.
0. [x] Add reading ease test

#### Backlog
0. [ ] Preprocess SCSS and save to public folder
0. [ ] Add unit tests to hit higher coverage %

[codecov-badge]: https://img.shields.io/codecov/c/github/mtso/lynx.svg
[codecov]: https://codecov.io/github/mtso/lynx
