# server-side-session-golang

Example for user authentication using server-side-session in golang using _fiber_ and _bcrpyt_

## Configuring Docker

```bash
docker run --name auth-psql -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=test -d posgres:14

docker ps

docker inspect CONTAINER_ID

get IP address (default 172.17.0.2)

docker exec -it CONTAINER_ID /bin/bash

```
