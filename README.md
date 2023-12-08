# Movie App
Movie App is a simple web application which enables user to store movie data in database. 
It consists of simple UI and API server written in go. 
Data is persisted in the MySQL database, optionally Redis cache can be configured to store recent API responses.

## Configuring Application
Backend server requires passing configuration file as a `-c` flag and database password as `DB_PASSWORD` environment variable. Application is always running of `8080` port.

```yaml
database:
  name: movies
  user: root
  address: localhost
  port: 3306
cache:
  enabled: false
  address: localhost
  port: 6379
```

- database.movies - name of the database to connect to
- database.user - name of database user
- database.address - URL of the database
- database.port - database port
- cache.enabled - if true, app should use Redis cache
- cache.address - URL of Redis cache
- cache.port - Redis port

Frontend application requires passing `REACT_APP_BACKEND_URL` env set to the backend URL.

## API

```bash
    GET /movies         # return a complete list of movies
    GET /movies/:id     # return a movie with given id
    POST /movies        # add a movie
    GET /readyz         # get application readiness
    GET /healthz        # get application health
    GET /server-info    # get some server-info
```


You can post data using e.g. `curl`.
```bash
curl localhost:8080/movie -XPOST -d'{"Title":"Pulp Fiction", "Director":"Quentin Tarantino"}'
```

# Building application

Backend:
```bash
cd backend

docker build -t localhost:5001/movieapp:0.1 -f Dockerfile .
# Or use multistage build
docker build -t localhost:5001/movieapp:0.1 -f Dockerfile.multistage .
```

Frontend:
```bash
cd ui

docker build -t localhost:5001/movieapp-ui:0.1 -f Dockerfile .
```

# Running application locally
-   API server
    ```bash
    DB_PASSWORD=root_password go run cmd/main.go -c ./sample-config.yaml
    ```

-   MySQL
    ```bash
    docker run --name mysqldb -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root_password -d mysql

    docker exec -ti mysqldb mysql -uroot -proot_password

    CREATE DATABASE movies;
    use movies;
    CREATE TABLE `movies` (`id` int(5) NOT NULL AUTO_INCREMENT, `title` varchar(255) DEFAULT NULL, `director` varchar(255) DEFAULT NULL, PRIMARY KEY (id)) ENGINE=InnoDB
    DEFAULT CHARSET=utf8;
    ```

-   Redis
    ```bash
    docker run --name redis -p 6379:6379 -d redis
    ```

-   UI
    ```bash
    REACT_APP_BACKEND_URL=http://<ip_addr> npm start
    ```
# Deploying on k8s
1. Prepare Deployment for the backend server
    - Set up liveness and readiness probe
    - Create ConfigMap with configuration and mount it
    - Create Secret with database password and use it as env
    - Expose it as a Service

2. Deploy MySQL database
    - Create ConfigMap with init script and mount it at `/docker-entrypoint-initdb.d` (or use a Job to initialize data)
    - Configure env variables
    - Expose it as a Service

3. Create Ingress for backend

4. Configure UI
    - Pass backend address as env
    - Expose it as a Service

5. Add Ingress for UI

6. Deploy Redis cache
    - Create Deployment
    - Expose it as a Service

7. Create HorizontalPodAutoscaler for backend

8. Optional load test: `ab -n 10000 -c 50 http://<backend_ingress>/movies`
