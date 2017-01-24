package lynxtemplate

// import "html/template"

// const {
// 	Index = iota
// 	Post
// }

// type TemplateType int

// type TemplateInfo {
// 	Type TemplateType
// 	Template *template.Template
// }

// type LynxTemplate {
// 	Templates []TemplateInfo
// }

const (
	IndexTemplate = `<!DOCTYPE html>
<html>
	<head profile="http://www.w3.org/2005/10/profile">
		<meta charset="utf-8">
		<link rel="stylesheet" href="./css/style.css">
		<title>{{.Title}}</title>

		<link rel="icon" 
		      type="image/png" 
		      href="https://i.imgur.com/BePVFjT.png">
		<link rel="apple-touch-icon"
		      href="https://i.imgur.com/BePVFjT.png">
	</head>
	<body>
		<main>
		<details>
			<summary>{{.Title}}</summary>
			<p>{{.Description}}</p>
		</details>
		<ul>
			{{range $index, $page := .Pages}}
			<li>
				<time>{{ $page.BirthTime.Format "January 2" }}</time>
				<a href="{{$page.RelativeLink}}">{{ $page.Title }}</a>
			</li>
			{{else}}<li>No posts.</li>{{end}}
		</ul>
		</main>
	</body>
</html>
`
	PostTemplate = `<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<link rel="stylesheet" href="../css/style.css">
		<title>{{.Title}}</title>
	</head>
	<body>
		<main>
		<a href="../">Index</a>
		<h1 class="title">{{.Title}}</h1>
		<section>
			<ul>
				<li><time>{{ .BirthTime.Format "January 2, 2006" }}</time></li>
				<li><a href="https://en.wikipedia.org/wiki/Flesch–Kincaid_readability_tests#Flesch.E2.80.93Kincaid_grade_level">Readability</a>: ~{{ .FleschKinkaid }}</li>
			</ul>
		</section>
		<article>
			{{ .Content }}
		</article>
		<footer>
		{{if .Next}}
			<a href="../{{.Next.RelativeLink}}">Next Post</a>
		{{end}}
		</footer>
		</main>
	</body>
</html>
`
// 	Css = `/* bolo style for lynx site-generator by mtso */

// main {
// 	overflow: hidden;
// 	max-width: 500px;
// 	margin-left: 8em;
// 	margin-bottom: 4em;
// }

// body {
// 	font-size: 1em;
// }

// details > p, 
// main > ul > li > time,
// section > ul > li > a {	
// 	font-size: 1em;
// }

// .title, 
// main > ul > li:first-child a {
// 	margin-top: 0.2em;
// 	margin-bottom: 0.2em;
// 	font-size: 6em;
// }

// {
// 	font-size: 2em;
// }

// summary,
// main > a,
// ul > li > a {
// 	font-size: 1.6em;
// }

// article,
// footer {
// 	font-size: 1.2em;
// }

// img, 
// video {
// 	width: 100%;
// }

// article {
// 	word-wrap: break-word;
// }

// time, 
// summary, 
// main > a, 
// main > ul > li:first-child a, 
// ul > li > a,
// footer > a {
// 	font-weight: bold;
// }

// main > ul > li > time {
// 	font-weight: normal;
// }

// ul time {
// 	display: block;
// }

// main > ul > li:first-child time {
// 	display: none;
// }

// ul {
// 	list-style: none;
// 	padding-left: 0;
// }

// section > ul > li,
// section > ul > li > time {
// 	display: inline;
// }

// section > ul > li {
// 	margin-right: 1em;
// }

// main > ul > li,
// article {
// 	margin-bottom: 1.6em;
// }

// /* small-screen layout */

// @media (max-width: 760px) {

// 	body {
// 		font-size: 1.4em;
// 	}

// 	.title, 
// 	main > ul > li:first-child a {
// 		font-size: 5em;
// 	}

// 	main {
// 		max-width: 100%;
// 		margin-left: 0.5em;
// 		padding-right: 0.5em;
// 		margin-bottom: 4em;
// 	}
// }

// @media only screen 
// 	and (min-device-width: 375px) 
// 	and (max-device-width: 667px) 
// 	and (-webkit-min-device-pixel-ratio: 2),
// only screen 
// 	and (min-device-width: 414px) 
// 	and (max-device-width: 736px) 
// 	and (-webkit-min-device-pixel-ratio: 3) {

// 	body {
// 		font-size: 2em;
// 	}

// 	.title, 
// 	main > ul > li:first-child a {
// 		font-size: 6em;
// 	}
	
// 	article {
// 		font-size: 1.6em;
// 	}
	
// 	summary,
// 	main > a,
// 	ul > li > a,
// 	footer {
// 		font-size: 2em;
// 	}

// 	main {
// 		max-width: 100%;
// 		margin-left: 0.5em;
// 		padding-right: 0.5em;
// 		margin-bottom: 2em;
// 	}
// }
// `
)

var HtmlTemplates = [...]string{
	`<!DOCTYPE html>
<html>
	<head profile="http://www.w3.org/2005/10/profile">
		<meta charset="utf-8">
		<link rel="stylesheet" href="./css/style.css">
		<title>{{.Title}}</title>

		<link rel="icon" 
		      type="image/png" 
		      href="https://i.imgur.com/BePVFjT.png">
		<link rel="apple-touch-icon"
		      href="https://i.imgur.com/BePVFjT.png">
	</head>
	<body>
		<main>
		<details>
			<summary>{{.Title}}</summary>
			<p>{{.Description}}</p>
		</details>
		<ul>
			{{range $index, $page := .Pages}}
			<li>
				<time>{{ $page.BirthTime.Format "January 2" }}</time>
				<a href="{{$page.RelativeLink}}">{{ $page.Title }}</a>
			</li>
			{{else}}<li>No posts.</li>{{end}}
		</ul>
		</main>
	</body>
</html>
`,
	`<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<link rel="stylesheet" href="../css/style.css">
		<title>{{.Title}}</title>
	</head>
	<body>
		<main>
		<a href="../">Index</a>
		<h1 class="title">{{.Title}}</h1>
		<section>
			<ul>
				<li><time>{{ .BirthTime.Format "January 2, 2006" }}</time></li>
				<li><a href="https://en.wikipedia.org/wiki/Flesch–Kincaid_readability_tests#Flesch.E2.80.93Kincaid_grade_level">Readability</a>: ~{{ .FleschKinkaid }}</li>
			</ul>
		</section>
		<article>
			{{ .Content }}
		</article>
		<footer>
		{{if .Next}}
			<a href="../{{.Next.RelativeLink}}">Next Post</a>
		{{end}}
		</footer>
		</main>
	</body>
</html>
`,
}

var Css = [...]string{
	`/* bolo style for lynx site-generator by mtso */

main {
	overflow: hidden;
	max-width: 500px;
	margin-left: 8em;
	margin-bottom: 4em;
}

body {
	font-size: 1em;
}

details > p, 
main > ul > li > time,
section > ul > li > a {	
	font-size: 1em;
}

.title, 
main > ul > li:first-child a {
	margin-top: 0.2em;
	margin-bottom: 0.2em;
	font-size: 6em;
}

{
	font-size: 2em;
}

summary,
main > a,
ul > li > a {
	font-size: 1.6em;
}

article,
footer {
	font-size: 1.2em;
}

img, 
video {
	width: 100%;
}

article {
	word-wrap: break-word;
}

time, 
summary, 
main > a, 
main > ul > li:first-child a, 
ul > li > a,
footer > a {
	font-weight: bold;
}

main > ul > li > time {
	font-weight: normal;
}

ul time {
	display: block;
}

main > ul > li:first-child time {
	display: none;
}

ul {
	list-style: none;
	padding-left: 0;
}

section > ul > li,
section > ul > li > time {
	display: inline;
}

section > ul > li {
	margin-right: 1em;
}

main > ul > li,
article {
	margin-bottom: 1.6em;
}

/* small-screen layout */

@media (max-width: 760px) {

	body {
		font-size: 1.4em;
	}

	.title, 
	main > ul > li:first-child a {
		font-size: 5em;
	}

	main {
		max-width: 100%;
		margin-left: 0.5em;
		padding-right: 0.5em;
		margin-bottom: 4em;
	}
}

@media only screen 
	and (min-device-width: 375px) 
	and (max-device-width: 667px) 
	and (-webkit-min-device-pixel-ratio: 2),
only screen 
	and (min-device-width: 414px) 
	and (max-device-width: 736px) 
	and (-webkit-min-device-pixel-ratio: 3) {

	body {
		font-size: 2em;
	}

	.title, 
	main > ul > li:first-child a {
		font-size: 6em;
	}
	
	article {
		font-size: 1.6em;
	}
	
	summary,
	main > a,
	ul > li > a,
	footer {
		font-size: 2em;
	}

	main {
		max-width: 100%;
		margin-left: 0.5em;
		padding-right: 0.5em;
		margin-bottom: 2em;
	}
}
`,
}
