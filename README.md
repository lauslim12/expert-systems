# Expert Systems

Expert Systems Research with a touch of Software Engineering with Go and React.

Will be used for a research paper like my previous projects: [building management system](https://github.com/lauslim12/Asuna), and [NDFA methodology to build cereal vending machines (Bahasa Indonesia)](http://proceeding.unindra.ac.id/index.php/simponi/article/view/375).

## Requirements

- Docker (if you want to containerize)
- Go 1.16+
- Node.js 16+
- Yarn 1.22+

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
yarn
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

- Clone this repository.

```bash
git clone git@github.com:lauslim12/expert-systems.git
```

- Switch to this repository.

```bash
cd expert-systems
```

- If running with Docker, do the following commands and after that, please open `localhost:8080` in your browser or run the provided integration tests with `make e2e`.

```bash
docker build . -t expert-systems:latest
docker run -d -p 8080:8080 expert-systems:latest
```

- If running with manual way, run tests first if you want.

```bash
make test
```

- Build React application.

```bash
cd web && yarn build
```

- Start application. Don't forget to return to the main repository by using `cd ..`, assuming you are in `web` folder.

```bash
make start
```

- Look at e2e or integration tests while the application is running. You might need to use another terminal.

```bash
make e2e
```
