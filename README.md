 #create a package that makes it easy to parse an HTML file and extract all of the links
 you can add the required html in the examples folder and  run " go run examples/ex1/main.go  " this will execute all the links with the text inside it stripped off any html and formatting
 This is basic parsing of an html document and can be worked on to do more
 package used https://pkg.go.dev/golang.org/x/net/html



 output for ex1 :-
 [{Href:/other-page Text:A link to another page some text}]