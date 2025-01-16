FROM --platform=${BUILDPLATFORM} golang:alpine3.21 AS builder

# enable go modules
ENV GO111MODULE=on
ENV PORT=4242

WORKDIR /app

COPY ../go.mod .
COPY ../go.sum .

# run only if go.mod or go.sum will be changed (cache)
RUN go mod download

COPY ../ ./

# Set target platform
ARG TARGETOS
ARG TARGETARCH

# build binary for corresponding platform without debug info
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags='-s -w' -o /app/greeter_api ./cmd/greeter

# generate clean, final image
FROM scratch

# copy golang binary into container
COPY --from=builder /app/greeter_api /app/

# executable
CMD ["/app/greeter_api"]