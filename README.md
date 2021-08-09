# ts

Tool for adding timestamps to a stream of lines from stdin.

## Install

## Usage

### Example

A good example for using `ts` is to annotate a `docker build` output with timestamps or step durations:

```
$ docker build . | docker run -i goeven/ts -d
   0s Sending build context to Docker daemon  79.87kB
   0s Step 1/7 : FROM golang:1.15-alpine as build
2.301s 1.15-alpine: Pulling from library/golang
691ms 801bfaa63ef2: Pulling fs layer
   0s ee0a1ba97153: Pulling fs layer
   0s 1db7f31c0ee6: Pulling fs layer
   0s ecebeec079cf: Pulling fs layer
   0s 63b48972323a: Pulling fs layer
  5ms ecebeec079cf: Waiting
   0s 63b48972323a: Waiting
699ms ee0a1ba97153: Verifying Checksum
   0s ee0a1ba97153: Download complete
  4ms 1db7f31c0ee6: Verifying Checksum
  1ms 1db7f31c0ee6: Download complete
 57ms 801bfaa63ef2: Verifying Checksum
   0s 801bfaa63ef2: Download complete
200ms 801bfaa63ef2: Pull complete
149ms ee0a1ba97153: Pull complete
 75ms 1db7f31c0ee6: Pull complete
176ms 63b48972323a: Verifying Checksum
   0s 63b48972323a: Download complete
2.836s ecebeec079cf: Verifying Checksum
   0s ecebeec079cf: Download complete
5.202s ecebeec079cf: Pull complete
 73ms 63b48972323a: Pull complete
 12ms Digest: sha256:49b4eac11640066bc72c74b70202478b7d431c7d8918e0973d6e4aeb8b3129d2
  5ms Status: Downloaded newer image for golang:1.15-alpine
  4ms  ---> 1463476d8605
   0s Step 2/7 : WORKDIR /src
665ms  ---> Running in af27aec81c47
109ms Removing intermediate container af27aec81c47
   0s  ---> 1d9f17228361
   0s Step 3/7 : ADD . .
166ms  ---> e6747312362f
   0s Step 4/7 : RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o /usr/local/bin/ts .
 39ms  ---> Running in c1bdab7de3ac
8.439s Removing intermediate container c1bdab7de3ac
   0s  ---> fcdd96453f10
   0s Step 5/7 : FROM scratch
   0s  --->
   0s Step 6/7 : COPY --from=build /usr/local/bin/ts /usr/local/bin/ts
209ms  ---> 557e2a332c0d
   0s Step 7/7 : ENTRYPOINT [ "/usr/local/bin/ts" ]
 38ms  ---> Running in 8e71999a8297
 79ms Removing intermediate container 8e71999a8297
   0s  ---> 82d370adecfb
  2ms Successfully built 82d370adecfb
```
