# Expert Systems

Expert System Research (diagnosis of Tuberculosis, focuses on TB for lungs, with Forward Chaining and Certainty Factor algorithm) in the form of an Application Programming Interface (API) with a touch of Software Engineering with Go and React.

This research has been completed and is now in progress to be submitted in an academic journal. This documentation will be updated with the bits and pieces of the research after it has been submitted properly and published.

## Abstract

Despite being a curable disease, Tuberculosis has become the leading cause of death of infectious disease prior to COVID-19. It has asymptomatic infections that are hard to detect for weeks or years. Although there have been many studies on Tuberculosis disease detection and prevention, very few of them discuss the creation of an expert system based on API. Hence, in this study we propose an Expert API that implements Forward Chaining and Certainty Factor algorithms for the task of Tuberculosis early detection. The evaluation of the proposed system was carried out using several testing methods and in-depth interviews with medical experts. We got a satisfactory result for this study.

**Keywords**: Certainty Factor, Expert System, Forward Chaining, REST API, Tuberculosis.

## Features

- API-based architecture with Go as the implementation of the inference engine.
- Usage of Certainty Factor and Forward Chaining algorithms in order to infer results.
- Responsive web design with complete `a11y` support with React and TypeScript.
- Built with performance in mind and deployed natively on a Cloud Platform.
- Codebase is fully formatted, linted, and documented with either JSDoc or Godoc.
- Fully unit-tested API with 100% code coverage.
- Supports `i18n`, with `en` and `id` as the internationalized languages.
- Simple, intuitive UI for a good user experience with ChakraUI and Ant Design.
- App includes 404 page and dark mode support.

## Requirements

- [Docker](https://www.docker.com/) (if you want to containerize)
- [Go 1.18+](https://golang.org/)
- [Node.js 16+](https://nodejs.org/en/)
- [Yarn 1.22+](https://yarnpkg.com/)
- Shell that supports `curl`, `make`, and `sh`. WSL / Ubuntu / OS X should be able to do this without much configuration (Bash version 3.0+).

Note: The tech stack may or may not be upgraded as time progresses (for example, from Node.js 16 to Node.js 18, or from React 18 to React 19 in the future).

## GitHub Codespaces

This repository supports GitHub Codespaces with a dedicated `.devcontainer`. You can create a Codespace based on the provided template if you want to get this application up and running as soon as possible. If you are going with this approach, after summoning the Codespace, you need to spawn two terminals: one to run `make start`, and the other to run `cd web` and then `yarn start`. You can see the result in the port-forwarded URL. All dependencies and setup are done in the container creation.

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

## Deployment

GitHub Actions has been set up in this repository to automatically test and deploy the production version of the application to [Heroku](https://www.heroku.com/). The Go application will serve the React frontend for simplicity, and the whole application is wrapped in a Docker container to ensure consistent builds and deployments.

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
