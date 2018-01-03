# nested

This prototype demonstrates how we can embed one gRPC server (`common`) in another (`extended`) so as to prevent any `extended` concepts from leaking into `common`.

## Motivation

A quick proof of concept of the design suggested in [FAB-7069](https://jira.hyperledger.org/browse/FAB-7069?focusedCommentId=36034&page=com.atlassian.jira.plugin.system.issuetabpanels:comment-tabpanel#comment-36034). Quick sketch of that idea:

![](https://user-images.githubusercontent.com/14876848/33777738-f75a2ee2-dc13-11e7-88ba-09ed6d200be4.png)

## Installation

```bash
$ go get -u github.com/kchristidis/nested/cmd/...
```

This will download, build, and install the `extended-server` and `extended-client` programs into the `$GOPATH/bin/` directory.

## Usage

Launch the server:

```bash
$ extended-server # binds extended server to localhost:5052
2018-01-03T18:30:48.574-0500    INFO    extended-server/main.go:37    Setting up common server at localhost:5051
2018-01-03T18:30:49.079-0500    INFO    extended-server/main.go:53    Common client connected to localhost:5051
2018-01-03T18:30:49.080-0500    INFO    extended-server/main.go:75    Setting up extended server at localhost:5052
2018-01-03T18:31:00.986-0500    DEBUG   extended-server/main.go:91    Received GET request with 'common.value' set to 42 and 'extra' set to 'hello world'
```

Then run the client:

```bash
$ extended-client localhost:5052 extended
2018-01-03T18:31:00.983-0500    INFO    extended-client/main.go:49    Invoking extended GET with 'common.value' set to '42' and 'extra' set to 'hello world'
2018-01-03T18:31:00.986-0500    INFO    extended-client/main.go:54    GET response: extra:"42 and hello world"
```

(Replace the `extended` argument with `common` for an example of a common request.)