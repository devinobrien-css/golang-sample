# Golang Sample
A sample golang server performing crud on a database

## Table of Contents

- [Prerequisites](#Prerequisites)
- [Installation](#Installation)
- [Testing](#Testing)


## Prerequisites
Basic knowledge of Golang is required. You can install [GOlang here](https://go.dev/doc/install)

This server is dependent on a running database instance. You can use one of the following:

### PostgreSQL
1. Clone the `docker-postgres` repository to your local machine:
    ```bash
    $ git clone https://github.com/devinobrien-css/docker-postgres.git
    ```

2. Follow the instructions provided in that repository's README to launch the instance.

## Installation
To get started with this Golang app, follow these steps:

1. Clone this repository to your local machine:
    ```bash
    $ git clone https://github.com/devinobrien-css/golang-sample.git
    $ cd your-repo
    ```

2. Build and run the app:
    ```bash
    $ go run .
    ```

4. Open your web browser and visit `http://localhost:8080` to access the app.

5. You're all set! Feel free to explore and modify the code to suit your needs.


### PostgreSQL

1. Run the following script to install the PG driver for go
    ```bash
    go get -u github.com/lib/pq
    ```

2. Clone `docker-postgres` repository to your local machine:
    ```bash
    $ git clone https://github.com/devinobrien-css/docker-postgres.git
    ```

3. Follow the instructions provided in that repository's readMe to launch the instance

## Testing

A [postman collection](postman-testing.json) has been included for convenient testing

