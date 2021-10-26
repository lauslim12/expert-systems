# Expert Systems

Expert Systems Research (diagnosis of Tuberculosis with Forward Chaining and Certainty Factor) with a touch of Software Engineering with Go and React. Licensed under MIT License, and in progress to be submitted in an academic journal.

This research is still in progress and not completed yet.

## Features

- API-based architecture with Go as the implementation of the inference engine.
- Usage of Certainty Factor and Forward Chaining algorithms in order to infer results.
- Responsive web design with complete `a11y` support with React and TypeScript.
- Built with performance in mind and deployed natively on a Cloud Platform.
- Fully unit-tested with 100% code coverage.
- Simple, intuitive UI for a good user experience (includes 404 page!).

## Requirements

- [Docker](https://www.docker.com/) (if you want to containerize)
- [Go 1.17+](https://golang.org/)
- [Node.js 16+](https://nodejs.org/en/)
- [Yarn 1.22+](https://yarnpkg.com/)
- Shell that supports `curl`, `make`, and `sh`. WSL / Ubuntu / OS X should be able to do this without much configuration (Bash version 3.0+).

## Installation (Development)

To develop further, you just need to do the following steps.

- Clone this repository.

```bash
git clone git@github.com:lauslim12/expert-systems.git
```

- Switch to this repository.

```bash
cd expert-systems
```

- Spawn another terminal, as you need two terminals: one to run Go, one to run the React application. Make sure to run the Go application first to prevent 'fail to proxy' error in React (development only).

```bash
# terminal 1
make start

# terminal 2
cd web
yarn --frozen-lockfile
yarn start
```

- You're good to go! Remember that the Go application does not hot-reload, while the React application does!

- Keep in mind that CI exists in order to make the code in proper standards. However, it's always nice to run the following before committing:

```bash
# for Go application
make format

# for React application
yarn lint
```

## Installation (Production)

There are two ways to run this application **for production**: first is to run with Docker, second is to run this manually. The recommended way is to run this with Docker.

### Manual

- Clone this repository.

```bash
git clone git@github.com:lauslim12/expert-systems.git
```

- Switch to this repository.

```bash
cd expert-systems
```

- Run unit tests first if you want.

```bash
make test
```

- If you want to run this manually, please build React application first.

```bash
cd web && yarn --frozen-lockfile && yarn build
```

- Build the Go application. Don't forget to return to the main repository by using `cd ..`, assuming you are in `web` folder.

```bash
make build
```

- Start the application!

```bash
./expert-systems
```

- Look at e2e or integration tests while the application is running. You might need to use another terminal.

```bash
make e2e
```

### Docker

- If running with Docker, do the following commands and after that, please open `localhost:8080` in your browser or run the provided integration tests with `make e2e`.

```bash
# Build and start
docker build . -t expert-systems:latest # choose either this or the below one
docker build . -t expert-systems:latest --build-arg GO_ENV=production # if you want HTTPS with 'X-Forwarded-Proto' header, some services like Heroku use this for HTTPS
docker run --name expert-systems -d -p 8080:8080 expert-systems:latest

# Remove image for cleanup
docker stop expert-systems
docker rm expert-systems
docker image rm expert-systems:latest
```

## License

This application is open sourced under MIT License.

## References

We would like to acknowledge these sources of helpful information for they have helped us during the research process.

- [Victor Caelina for the TB Symptoms Dataset](https://www.kaggle.com/victorcaelina/tuberculosis-symptoms)
