# GO HTMX BOOTSTRAP

A simple starter app that uses **GOTH** stack. 

## Notable Dependencies 
- [Templ](https://templ.guide/server-side-rendering/creating-an-http-server-with-templ) - use as the main templating engine 
- [Go Chi](https://go-chi.io/) - HTTP routing 
- [Temp UI](https://templui.io/docs/introduction) - UI tool kit
- [HTMX](https://htmx.org/) - already included but you can replace it if you want in `assets/js`and download your desired version
- [Alpine JS ](https://alpinejs.dev/) - Currently used via CDN will change in the future hehe 😇

## Running the app

### rename.sh
The repository includes a script that renames all imports to your liking

```
$ ./rename.sh github.com/rhenerp/go-htmx-bootstrap myapp.com/your-much-better-app
```

If permission denied, in project dir

```
$ chmod +x rename.sh
```

Then run `go mod tidy` after 

### Developm
ent 

Before running make dev make sure you installed the following 
- Tailwind on your local machine
- templ compiler
- [air](https://github.com/air-verse/air)- for hot reloading  

```
$ make dev
```

This will listen to project changes and reload it, definitions of the watched files can me seen in the Makefile

During development static files are being served from disk and is not being cached by the browser this decision can be seen in the `main.go` file `w.Header().Set("Cache-Control", "no-store")`.

Development config can be seen in `config/dev.go`

### Building and Running 

```
$ GOOS= GOARCH= CGO_ENABLED= go build -ldflags="-s -w" -o yourapp
```

then run`ENV=staging PORT=3000 ./yourapp`, running the app in **non development** ENV will use the embedded static files instead of disk read. 

## Folder structure 
```
.
├── assets                      # Embedded static files
│   ├── assets.go              # Embed definition
│   ├── css
│   │   └── input.css
│   └── js
│       ├── htmx.min.js        # HTMX from CDN (can be replaced)
│       ├── input.min.js       # From templUI
│       └── tabs.min.js        # From templUI

├── config                     # Environment-based config
│   ├── config.go              # Config type defs and helpers
│   ├── dev.go                 # Development config
│   ├── production.go
│   └── staging.go

├── handlers                   # Route controllers
│   ├── home.go
│   └── items.go

├── models                     # Type definitions
│   └── Item.go

├── services                   # Business logic
│   └── items.go

├── ui                         # Templ UI components
│   ├── components             # Generated components
│   ├── layouts                # Layouts (restructure as needed)
│   ├── modules                # Modules (restructure as needed)
│   └── pages                  # Page-level templates

├── utils
│   └── templui.go             # Needed for templUI

├── .templui.json              # Metadata for templUI
├── go.mod
├── go.sum
├── items.json                 # Sample data
├── main.go                    # App entry point
├── Makefile
├── rename.sh                  # Module rename script
```
