FROM golang:alpine AS build-env
WORKDIR /tmp/app
COPY src/github.com/kambui/write-to-files/write-to-files.go .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o write-to-file

FROM registry.access.redhat.com/ubi7/ubi AS myapp
WORKDIR /app
RUN yum install glibc -y
COPY --from=build-env /tmp/app/write-to-file /app/
ENTRYPOINT ./write-to-file
CMD [""]
