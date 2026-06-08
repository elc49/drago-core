# run templ generation in watch mode to detect all .templ files and 
# re-create _templ.txt files on change, then send reload event to browser. 
# Default url: http://localhost:7331
live/templ:
	go tool templ generate --watch --proxy="http://localhost:8080" --cmd="go run ." --open-browser=false

# run tailwindcss to generate the styles.css bundle in watch mode.
live/tailwind:
	npx --yes tailwindcss -i ./assets/css/input.css -o ./assets/css/styles.css --minify --watch

# run esbuild to generate the index.js bundle in watch mode.
live/esbuild:
	npx --yes esbuild ./assets/js/index.ts --bundle --outdir=./assets/js/ --watch

# watch for any js or css change in th# start all 5 watch processes in parallel.
live: 
	make -j5 live/tailwind live/templ live/esbuild
