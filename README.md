## How to start

Start API server (conclude migration) - local environment

```bash
./scripts/bin.sh api start
```

Start worker update site state with interval = 10s - local environment

```bash
./scripts/bin.sh api worker_start
```
API to get info of a site

```bash
localhost:8090/v1/sites?name=weibo.com
```

API to get site with maximum access time
```bash
localhost:8090/v1/sites?is_maximum_access_time=true
```
API to get site with minimum access time
```bash
localhost:8090/v1/sites?is_minimum_access_time=true
```
