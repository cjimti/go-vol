[![Docker Container Image Size](https://shields.beevelop.com/docker/image/image-size/cjimti/go-vol/latest.svg)](https://hub.docker.com/r/cjimti/go-vol/)
[![Docker Container Layers](https://shields.beevelop.com/docker/image/layers/cjimti/go-vol/latest.svg)](https://hub.docker.com/r/cjimti/go-vol/)
[![Docker Container Pulls](https://img.shields.io/docker/pulls/cjimti/go-vol.svg)](https://hub.docker.com/r/cjimti/go-vol/)

# GO Vol

Go Vol is for testing [Kubernetes] volumes. Pods will write a file to the specified path with a file consisting the environment variable vales for NODE_NAME and POD_NAME. This helps test and debug issues where Pods are not correctly connecting to a volume. The list of files on the volume is available from an http endpoint on a specified port.

## Test with [Docker]

Run the container from a terminal:
```bash
docker run --rm -it -p 8080:80 -v "$(pwd)"/files:/files \
    -e VOL_PATH="/files" -e PORT=80 \
    -e NODE_NAME=n1.imti.cloud \
    -e POD_NAME=go-vol-xxx2 \
    cjimti/go-vol
```

Browse to http://localhost:8080 and you should see JSON output similar to the one below:

```json
{
  "client_ip": "172.17.0.1",
  "count": 2,
  "file": "/files/gv_n1.imti.cloud_go-vol-xxx2.txt",
  "files": [
    ".gitignore",
    "gv_n1.imti.cloud_go-vol-xxx2.txt",
    "gv_n1.imti.cloud_go-vol-xxxx.txt",
    "hello-world.txt"
  ],
  "message": "go-vol",
  "node_name": "n1.imti.cloud",
  "pod_ip": "",
  "pod_name": "go-vol-xxx2",
  "pod_namepage": "",
  "service_account": "",
  "status_message": "OK",
  "time": "2018-03-29T20:56:00.3109565Z",
  "uuid_call": "0fe81ee2-1e5f-40e5-79ac-ae219af26a3a",
  "uuid_instance": "f118620f-3194-46fc-76c5-d32b5d7bfe01",
  "version": 1,
  "version_msg": "version 1"
}
```


## Resources
- [Docker]
- [Kubernetes]


[Docker]: https://www.docker.com/
[Kubernetes]: https://kubernetes.io/