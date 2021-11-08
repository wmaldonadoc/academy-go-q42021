# API

## Setup

Pre requisites:
 - Docker >= 20.10.x
 - Docker-compose >= 1.29.x
 - Postman

 ## Docs

 You can import the **collection** and **environment** located at `docs/postman`.
 > **Note:** The **URL** settled in the environment file is mapped to **localhost:300**

 ## Running the project

> You need to create a `.env` file based on `.env.example` and fill the necessary values.

 Execute the docker-compose command

 ```bash
 docker-compose up --build
 ```