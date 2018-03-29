
# GO Vol

Go Vol is for testing [Kubernetes] volumes. Pods will write a file to the specified path and filename. The list of files on the volume is available from an http endpoint on a specified port.

## Test with [Docker]

Run the container from a terminal:
```bash
docker run --rm -it -e TCP_PORT=2701 -e NODE_NAME="EchoNode" -p 2701:2701 cjimti/go-echo
```



## Resources
- [Docker]
- [Kubernetes]


[Docker]: https://www.docker.com/
[Kubernetes]: https://kubernetes.io/