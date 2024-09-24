# trail-browser

## Introduction
This app uses the trails data in the following website (https://opendata-bouldercounty.hub.arcgis.com/datasets/3a950053bbef46c6a3c2abe3aceee3de_0/explore) and processes it from from a csv file (part of this project) and list the trails in it.

There are two filters available:

`park_spaces`: Filters results with that exact amount of parking spaces
`access_name`: Filters results where the access_name field contains the string sent

## Setup
run `docker compose up -d`
send a http request to `localhost` using the defined endpoint

## Endpoints available

`GET /trails` this will return a list of trails

