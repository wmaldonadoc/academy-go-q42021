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

 ## ABOUT

 This project it's a REST API that works with [PokeAPI](https://pokeapi.co).

 Has 4 endpoints:

 **Error response format**
 *Example*
 ```json
 {
    "message": "Error searching pokemon by id",
    "httpStatus": 404,
    "code": 1000
 }
```
 
 Base URL: `{HOST}/api/v1`
### Get a pokemon by ID:

 PATH: `/pokemons/id/:id` 

**RESPONSE**
 Success response, HTTP Status `200`
 ```json
 {
    "ID": 220,
    "Name": "ditto",
    "Ability": "limber"
}
 ```
 ### Get a pokemon by name:

 PATH: `/pokemons/name/:name` 

**RESPONSE**
 Success response, HTTP Status `200`
 ```json
 {
    "ID": 220,
    "Name": "ditto",
    "Ability": "limber"
}
 ```
 ### Batch filter pokemons:

 PATH: `/pokemons/filter?type=STRING&items=INTEGER&items_per_worker=INTEGER` 

- **type:** Only support "odd" or "even"
- **items:** Is an Int and is the amount of valid items you need to display as a response
- **items_per_workers:** Is an Int and is the amount of valid items the worker should append to the response

**RESPONSE**
 Success response, HTTP Status `200`
 ```json
[
        {
        "ID": 220,
        "Name": "ditto",
        "Ability": "limber"
        },
        {
            "ID": 220,
            "Name": "ditto",
            "Ability": "limber"
        },
        {
            "ID": 220,
            "Name": "ditto",
            "Ability": "limber"
        }
]
 ```
  ### Get service health

 PATH: `/health` 

**RESPONSE**
 Success response, HTTP Status `200`
 ```json
{
    "Uptime": 137,
    "StatusCode": 200
}
 ```