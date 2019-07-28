# API Documentation

## Note About Setup

The API assumes that there is a psql database called yawoen, with username and
password set to postgres running on port 5432. You can change these settings in
the .env file or create a Database via the psql shell. It is important to note
that the password has to be manually set since the default authentication 
method is set to ident.

## Installation

The API imports the external packages mux and godotenv, as well as the psql
driver. Other than that, all imports are from the Go Standard Library.

## Usage

The API has 6 endpoints which satisfy the requirements for the challenge.

### "/" Endpoint

Simply redirects to the "/company" endpoint.

### "/company" Endpoint

Returns all the entries in the database.

### "/company/{id}" Endpoint

Returns the information for a single company, which is specified in the URL via
ID.

### "/company/match" Endpoint

Returns the information for a single company. This company is specified by 
Name and Zip information which are expected in the request body as a JSON.

### "/populate" Endpoint

Populates the psql database with information loaded from a csv file. This 
endpoint expects a path in the request body (to the csv file).

### "/integrate_website" Endpoint

Adds website information to the entries where applicable from the database.

