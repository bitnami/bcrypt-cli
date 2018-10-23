# bcrypt in a container
#
# echo -n mysecret | docker run --rm -i bitnami/bcrypt
#
FROM golang:1.11-stretch as build

RUN apt-get update && apt-get install -y --no-install-recommends \
    git make upx \
    && rm -rf /var/lib/apt/lists/*

RUN wget -q -O dep https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 && \
    echo '287b08291e14f1fae8ba44374b26a2b12eb941af3497ed0ca649253e21ba2f83 dep' | sha256sum -c - && \
    mv dep /usr/bin/ && chmod +x /usr/bin/dep

RUN go get -u \
        github.com/golang/lint/golint \
        golang.org/x/tools/cmd/goimports \
        github.com/golang/dep/cmd/dep \
        && rm -rf $GOPATH/src/* && rm -rf $GOPATH/pkg/*


WORKDIR /go/src/app
COPY . .

RUN rm -rf out

RUN make

RUN upx --ultra-brute out/bcrypt

FROM bitnami/minideb:stretch

COPY --from=build /go/src/app/out/bcrypt /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/bcrypt"]
