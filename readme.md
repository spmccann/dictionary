# Dictionary

![dictionary](https://github.com/spmccann/dictionary/blob/main/screenshots/Screenshot2.png)

## What It Does 
The webpage swiftly (and without refreshing the page) returns word definitions from a user search

## Technologies

- [Go](https://go.dev/) (Backend)
- [Templ](https://templ.guide/) (Templating engine)
- [HTMX](https://htmx.org/) (Front end framework)

## Data

The dictionary data is a JSON file sourced from eddydn's [DictionaryDatabase](https://github.com/eddydn/DictionaryDatabase) repo which was pulled from The Online Text Plain English Dictionary (OPTED)

## Project Directory Structure
```
.
├── components
│   ├── components.templ
│   └── components_templ.go
├── data
│   ├── EDMTDictionary.json
│   └── EDMTDictionary.json:Zone.Identifier
├── go.mod
├── go.sum
├── handlers
│   └── default.go
├── main.go
├── readme.md
└── services
    └── definition.go
```
## How To Run and View

The web page has only been implemented locally. It can be accessed by cloning/forking the repo, `go run .` and viewed in the browser on localhost:3000 

## Project Recap
I made this web page for the first personal project on [Boot.dev](http://boot.dev). It doesn't use any Javascript or CSS and has a minimal amount of HTML

### Things I Found Challenging

- Trying to understand how HTTP works. I have a fuzzy idea what a handler is trying to do with a GET or POST but there's a lot more going on than I have time to digest for the timeframe of this project
- Using Templ for the first time, I don't like how `templ generate` has to be run with every change to the code base. For my next web project, I'll likely try something else or just use Go's html/templ

### Things I Found Fun

- Writing a binary search function
- HTMX making frontend easy by just working out of the box
