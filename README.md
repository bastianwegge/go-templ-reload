## templ live-reload without watch-proxy

I noticed that I'd have to switch around between `templ generate --watch` a lot if I wanted to commit my work. This is a simple example on how to avoid switching back and forth, thus logging into a nice developer experience when working with templ.

### How to run

```bash
# install https://taskfile.dev/installation/
brew install go-task/tap/go-task

# run the install task
task install

# run the dev task
task dev
```

### How it works

The `task dev` command will start a development server on `localhost:3000` and watch for changes in the files to generate the templ files and then recompile. A small script that's inside [page.templ](./pkg/feature-x/page.templ) will create an [EventSource](https://developer.mozilla.org/en-US/docs/Web/API/EventSource) that is trying to connect to the server, the server is sending a keepalive ping. As soon as the server restarts, the browser also refreshes using `window.location.reload();`.