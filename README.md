# Expert Systems

Expert Systems API Research (diagnosis of Tuberculosis, focuses on TB for lungs, with Forward Chaining and Certainty Factor) with a touch of Software Engineering with Go and React.

This research has been completed and is now in progress to be submitted in an academic journal.

## Abstract

Tuberculosis is a disease that is often overlooked by medical personnel and is difficult to diagnose when it is at its early stage. The disease itself is treatable and there are indeed early detection tools, but one of the challenges for the early detection tools is its price. As with all other diseases, it is important to make early detection tools that are easy to use, accurate, and can be used by people from various backgrounds and environments. Early detection tools with said characteristics will be able to help contribute in eradicating Tuberculosis from the world. This research will focus on development of an Expert System with a REST API based architecture to diagnose Tuberculosis based on the available symptoms. The API is built so the system could scale, achieve more separation of concerns, and be portable. The Expert System will implement Forward Chaining and Certainty Factor algorithms in order to provide a more in-depth diagnosis. The user interface of the system is developed as a web application with several features, one of which is internationalization, and the languages that are used are English and Bahasa Indonesia. The evaluation of this research is carried out with unit tests, integration tests, interviews with general users, and in-depth interviews with medical experts. We received a satisfactory result, where most respondents claimed that the system is easy to use, informative, accurate, helpful, and they consider using this system to diagnose themselves in the future.

**Keywords**: API, Web Application, Internationalization, Expert System, Tuberculosis

## Features

- API-based architecture with Go as the implementation of the inference engine.
- Usage of Certainty Factor and Forward Chaining algorithms in order to infer results.
- Responsive web design with complete `a11y` support with React and TypeScript.
- Built with performance in mind and deployed natively on a Cloud Platform.
- Codebase is fully formatted, linted, and documented with either JSDoc or Godoc.
- Fully unit-tested with 100% code coverage.
- Supports `i18n`, with `en` and `id` as the internationalized languages.
- Simple, intuitive UI for a good user experience with ChakraUI and Ant Design.
- App includes 404 page and dark mode support.

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

This application is open sourced under MIT License. The research paper itself will follow the license of the publisher after it has been published.

## Acknowledgements

We would like to acknowledge these sources of helpful information for they have helped us during the research process.

- Doctor Maria from Halodoc who had reviewed this application.
- Doctor Devi from Halodoc who had reviewed this application.
- Segun Adebayo for ChakraUI.
- Peter Kieltyka for Chi web framework.
- Kamijin Fanta for React Icons.
- i18next for the internationalization library.
- All testers of the system, who have spent their time evaluating, giving feedback, and helping us to improve this system to be a better one.

## References

The bibliography of our research can be seen in our research paper after it has been published.
