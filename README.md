# Go Market

The go-market built in Golang facilitates interactions between customer and product owner. It provides essential functionalities such as user authentication, products browsing, and purchasing.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Documentations](#documentations)

## Installation

1. Make sure you have Golang installed. If not, you can download it from [Golang Official Website](https://go.dev/doc/install).

2. Install 'make' if not already installed. 

    * On Debian/Ubuntu, you can use

    ```bash
    sudo apt-get update
    sudo apt-get install make
    ```

   * On macOS, you can use [Homebrew](https://brew.sh/)

    ```bash
    brew install make
    ```

   * On Windows, you can use [Chocolatey](https://chocolatey.org/)

    ```bash
    choco install make
    ```

3. Clone the repository

    ```bash
    git clone https://github.com/wildanfaz/go-market.git
    ```

4. Change to the project directory

    ```bash
    cd go-market
    ```

5. Copy example-config.json to config.json

    ```bash
    cp example-config.json config.json
    ```

## Usage

1. Start the application using docker

    ```bash
    docker-compose up -d
    ```

## Commands

1. Install all dependencies
    ```bash
    make install
    ```

2. Start the application without docker
    ```bash
    make start
    ```

3. Add user balance (change host.docker.internal in config.json->database.mysql_dsn to localhost)
    ```bash
    make add-balance email=${email}
    ```

## Documentations

1. [Postman](https://documenter.getpostman.com/view/22978251/2sA3QwbUs1)

2. [ERD](https://dbdiagram.io/d/Go-Market-665b2823b65d9338793e99c2)

3. [Docker](https://hub.docker.com/repository/docker/muhamadwildanfaz/go-market)