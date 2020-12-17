# setup builder
FROM 916622979474.dkr.ecr.ap-southeast-1.amazonaws.com/base/golang:1.15.6-alpine as builder
# set tag builder to be removed later
LABEL tag="builder"

# add git to fetch private repo
RUN apk add git

# send GITHUB_USERNAME and PAT
ARG GITHUB_USERNAME
ARG GITHUB_PAT

# use GOMODULE
ENV GO111MODULE=on

# supply remote creds
RUN printf "machine github.com\n\
    login ${GITHUB_USERNAME}\n\
    password ${GITHUB_PAT}\n\
    \n\
    machine api.github.com\n\
    login ${GITHUB_USERNAME}\n\
    password ${GITHUB_PAT}\n"\
    >> /root/.netrc
RUN chmod 600 /root/.netrc

WORKDIR /app
COPY ./go.mod .
COPY ./go.sum .
COPY ./config /builder/config
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build \
    -o /builder/warden-api ./cmd/warden-api/main.go

FROM 916622979474.dkr.ecr.ap-southeast-1.amazonaws.com/base/ubuntu:18.04
WORKDIR /app
COPY --from=builder /builder .
ARG APPLICATION_ENVIRONMENT
ENV APPLICATION_ENVIRONMENT=$APPLICATION_ENVIRONMENT
EXPOSE 8000
ENTRYPOINT ["./warden-api"]

