# Golang Sample

A sample golang server performing crud on a database

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Testing](#testing)

## Prerequisites

Basic knowledge of Golang is required. You can install [GOlang here](https://go.dev/doc/install)

This server is dependent on a running database instance. You can use one of the following:

### PostgreSQL

1. Clone the `docker-postgres-sample` repository to your local machine:

    ```bash
    $ git clone https://github.com/devinobrien-css/docker-postgres-sample.git
    ...
    ```

2. Follow the instructions provided in that repository's README to launch the instance.

## Installation

To get started with this Golang app, follow these steps:

1. Clone this repository to your local machine:

    ```bash
    $ git clone https://github.com/devinobrien-css/golang-sample.git
    $ cd your-repo
    $
    ```

2. Build and run the app:

    ```bash
    $ go run .
    Server is running on http://localhost:8080
    ```

3. Open your web browser and visit `http://localhost:8080` to access the app.
4. You're all set! Feel free to explore and modify the code to suit your needs.

### Launch Container (optional)

1. Clone `docker-postgres` repository to your local machine:

    ```bash
    $ git clone https://github.com/devinobrien-css/docker-postgres.git
    $
    ```

2. Follow the instructions in the repository's README for next steps.

## Testing

A [postman collection](postman-testing.json) has been included for convenient testing
