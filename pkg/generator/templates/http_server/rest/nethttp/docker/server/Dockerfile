# ============
# Builder
# ============
FROM golang:1.18.3 as Builder

ENV TZ=Asia/Tokyo
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV ROOTPATH=/go/src

WORKDIR ${ROOTPATH}

COPY go.mod go.sum ./
RUN go mod download

COPY app/ ${ROOTPATH}/app/

RUN go build \
    -o ${ROOTPATH}/bin/main \
    -ldflags '-s -w' \
    ./app/cmd/server/main.go


# ============
# Deploy
# ============
FROM gcr.io/distroless/static-debian11
ENV TZ=Asia/Tokyo

COPY --from=Builder /go/src/bin/main /src/main

EXPOSE 8000
USER nonroot:nonroot
CMD ["/src/main"]