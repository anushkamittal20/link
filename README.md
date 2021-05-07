 #create a package that makes it easy to parse an HTML file and extract all of the links
 you can add the required html in the examples folder and  run " go run examples/ex1/main.go  " this will execute all the links with the text inside it stripped off any html and formatting
 This is basic parsing of an html document and can be worked on to do more
 package used https://pkg.go.dev/golang.org/x/net/html .
 Package html implements an HTML5-compliant tokenizer and parser

parse.go creates a slice of Link struct which contains the 'href' and text inside a tag.
To execute this program as it is ```go run examples/ex1/main.go``` for any pre written example that you'd like.
```
output for ex1 :-
 [{Href:/other-page Text:A link to another page some text}]
 ```
To execute the functionality for customized html program, in main.go of examples/ex1
```
var exampleHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page
  <span>some text</span></a>
</body>
</html>
`
```
replace the content inside the backticks of ```var exampleHtml``` with your html.
To execute this, run the following in the terminal
```
go run examples/ex1/main.go
```
