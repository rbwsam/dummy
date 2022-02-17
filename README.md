# dummy
HTTP server for testing cluster deployments.

- Very small container image for fast deployments.
- Responds with JSON including the server's hostname.

# Run
Run the container:
```bash
$ docker run -p 127.0.0.1:80:80 run rbwsam/dummy
```

In another terminal, send a request to the server:
```bash
$ curl -v 127.0.0.1
*   Trying 127.0.0.1:80...
* Connected to 127.0.0.1 (127.0.0.1) port 80 (#0)
> GET / HTTP/1.1
> Host: 127.0.0.1
> User-Agent: curl/7.74.0
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=UTF-8
< Date: Thu, 17 Feb 2022 19:43:53 GMT
< Content-Length: 33
< 
{
  "hostname": "27c04ba3c3ba"
}
* Connection #0 to host 127.0.0.1 left intact

```