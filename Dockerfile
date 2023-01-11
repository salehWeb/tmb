FROM golang:1.19
FROM node:16-alpine

WORKDIR ./

COPY . .

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]
RUN npm ci
# Build the app
RUN npm run build
# ==== RUN =======
# Set the env to "production"
ENV NODE_ENV production
# Expose the port on which the app will be running (3000 is the default that `serve` uses)
EXPOSE 3000
# Start the app
CMD [ "npx", "serve", "build" ]
