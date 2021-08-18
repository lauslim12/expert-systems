# First stage: Go process.
FROM golang:1.16.6 AS api

# Set working directory.
WORKDIR /backend-compile

# Copy dependency locks so we can cache.
COPY go.mod go.sum .

# Get all of our dependencies.
RUN go mod download

# Copy all of our remaining application.
COPY . .

# Build our application.
RUN CGO_ENABLED=0 GOOS=linux go build -o expert-systems ./cmd/expert-systems/main.go

# Get Node image from DockerHub.
FROM node:16.6.2 AS web

# Set working directory.
WORKDIR /frontend-compile

# Copy dependency locks.
COPY ./web/package.json ./web/yarn.lock .

# Get all of our dependencies.
RUN yarn

# Copy all of our remaining application.
COPY ./web .

# Build our application.
RUN yarn build

# Use 'scratch' image for super-mini build.
FROM scratch AS prod

# Set working directory for this stage.
WORKDIR /production

# Copy our compiled executable from the last stage.
COPY --from=api /backend-compile/expert-systems .
COPY --from=web /frontend-compile/build ./web/build

# Run application and expose port 8080.
EXPOSE 8080
CMD ["./expert-systems"]
