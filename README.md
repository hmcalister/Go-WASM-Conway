# Go WASM Conway

A simple implementation of Conway's Game of Life using [Ebitengine](https://ebitengine.org/) and WebAssembly.

[See the game in action.](https://conway.hmcalister.nz)

## Compiling and Running 

To compile and run the project locally, simply clone this repository and`cd` into the project root.

Note that Ebitengine has several requirements to build the project which you must install before compiling. Ensure you have installed the [requirements for your operating system](https://ebitengine.org/en/documents/install.html) and [read through the information on Ebitengine and WASM](https://ebitengine.org/en/documents/webassembly.html).

This project uses WebAssembly (WASM), which Go supports with the provided `wasm_exec.js` file. While a copy of this file is present in this repo, it may have been updated since the last push. Get a new copy by either manually copying the file in to the `assets` directory from `$(GOROOT_PATH)/misc/wasm/wasm_exec.js`, or simply run `make init` to do this for you.

Run `make build` to compile the package at `cmd/wasm` to a WebAssembly file, as well as the package at `cmd/server` to a simple webserver serving the `assets` directory. Then, run the server with `./ConwayWASM` to serve the project on `http://localhost:8080`.

## Explanation

Ebitengine, a simple game engine for Go, can be compiled to WebAssembly, which can then be run natively in the browser. This means you can write relatively complex and powerful web applications in languages other than Javascript.

This process is not without its difficulties. The default behavior for Go's `wasm_exec.js` is to effectively overwrite any other styling you place on the page. Therefore, it is suggested you either:
- Write the entire application to interface through the WebAssembly, or
- Embed the WebAssembly page in an `<iframe>` tag.

The first approach requires an immense amount of work as the application scales up, as we cannot rely on any of the preexisting HTML or CSS architecture. Theoretically, such an app may have excellent performance, but would require significant development time.

The second approach allows for other web app features to surround the `<iframe>`, but then requires the user to focus into the frame to interact with the WASM app. This is also not ideal, as the user will naturally want to interact with both the traditional web application and the WASM application interchangeably.

Aside from this, Ebitengine is a *game* engine, and is therefore not suited for making industry scale web applications. It is, however, relatively easy to make visualizations or other simple interactive apps (e.g. Conway's Game of Life). So not all is lost!