# Docker

## 1. Install Docker

1. Install Docker on your machine. You can find the installation instructions [here](https://docs.docker.com/get-docker/).
   1. *EXTRA*: Attach a screenshot of the output of `docker --version` and `docker-compose --version`.

## 2. Create a Dockerfile

1. Create a Dockerfile image in the root of the repository. The image should compile the Golang program, and expose port 8080.
   1. *EXTRA*: Use a multi-stage build to reduce the size of the final image.
   2. *EXTRA*: Run the Docker image locally and test that it works. Attach a screenshot of the output.

## 3. Create a docker-compose.yml

1. Create a docker-compose.yml file in the root of the repository. The file should contain two services:
   1. A service that runs the Golang program.
   2. A service that runs a MySQL database (this service is not used yet, but will be used in the next task).
   3. *EXTRA*: Run the docker-compose locally and test that it works. Attach a screenshot of the output.

## 4. Connect the application to the database

1. Modify the docker-compose.yml file to connect the application to the database.

To enable the `/db` endpoint, you need to set the `ENABLE_DB` in the docker-compose file.

You will also need to set the following environment variables:
- `DB_HOST` - the host of the database. This should be the name of the database service in the docker-compose file.
- `DB_USER` - the username of the database. This should be the username of the database service in the docker-compose file.
- `DB_PASSWORD` - the password of the database. This should be the password of the database service in the docker-compose file.
- `DB_DATABASE` - the name of the database. This should be set to `app_db`.

To initialize the database, you can use the `app_db.sql` file in the `db/init` folder of the repository, mounting it as a volume in the docker-compose file.

```yaml
volumes:
  - ./db/init:/docker-entrypoint-initdb.d
```

## 5. Push the image to Docker Hub

1. Create a Docker Hub account if you don't have one already. It's free, and you can find the instructions [here](https://docs.docker.com/docker-hub/).
2. Push the image to Docker Hub. You can find the instructions [here](https://docs.docker.com/docker-hub/repos/).
3. Provide the link to the image in the README.md file.

LINK : https://hub.docker.com/repository/docker/skurub2/my-public-repo
