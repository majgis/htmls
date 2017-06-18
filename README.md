# htmls
A golang experiment to mashup HTML from multiple microservices serving html snippets

## Start server

To start the server, call https with one argument of a template path

    https template.html

## Template tokens

Tokens in the HTML file take the following form:

    {{uri="http://localhost:3000"}}

If there is an error loading the content from the URI, the following will be placed in its stead:

    <!-- ERROR -->

