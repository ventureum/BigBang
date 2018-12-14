
[![Build Status](https://travis-ci.com/ventureum/BigBang.svg?branch=master)](https://travis-ci.com/ventureum/BigBang)

# BigBang

BigBang is a server-less backend system that is responsible to provide OFF-CHAIN data backups and APIs for Milestone
platform operations. It contains AWS API Gateway endpoints powered by lambda functions written in golang and built
using Bazel. It has three tiers: the bottom tier Postgres Database,  the middle tier AWS Lambda, and the top
tier AWS API Gateway.

1. Postgres Database
    - It backups OFF-CHAIN data in postgres database, and serves requests for data CRUD (create, read, update and delete)
      operations from the middle tier AWS Lambda.
2. AWS Lambda
    - It acts the middle communicator that converts the business logic requests from API ending points into
      database-level operations.
3. AWS API Gateway
   - It provides API ending points for business logic level operations, and send business logic requests or retrieves
     response from the middle tier  AWS Lambda.
     
![](assets/images/BigBang.png)    


# Set Up BigBang Components


## Create ENV File

Copy a new ENV file from ENV_TEMPLATE to set up ENV variables, and name your file like .env.test or any style you like

It contains the following variables

```js
AWS_ACCESS_KEY_ID=<Add Value>
AWS_SECRET_ACCESS_KEY=<Add Value>
AWS_SECRET_KEY=<Add Value>
AWS_REGION=<Add Value>

STREAM_API_KEY=<Add Value>
STREAM_API_SECRET=<Add Value>

DB_USER=<Add Value>
DB_PASSWORD=<Add Value>
DB_NAME=<Add Value>
DB_HOST=<Add Value>

DB_NAME_PREFIX=<Add Value>
DB_HOST_POSTFIX=<Add Value>

MuMaxFuel=<Add Value>
MuMinFuel=<Add Value>
PostFuelCost=<Add Value>
ReplyFuelCost=<Add Value>
AuditFuelCost=<Add Value>
BetaMax=<Add Value>
REFUEL_INTERVAL=<Add Value>
FUEL_REPLENISHMENT_HOURLY=<Add Value>
TIME_INTERVAL_FOR_FUEL_UPDATE=<Add Value>
MAX_FUEL_FOR_FUEL_UPDATE_INTERVAL=<Add Value>
```

## Set Up local Postgres Database

### Get Postgres Image 

```js 
docker pull  postgres
```

### Run Postgres Container 

```js
docker run --rm --name feed-sys-postgres  -e POSTGRES_PASSWORD=<postgres_password> -e POSTGRES_USER=<postgres_user> -e POSTGRES_DB=<db_name>  -P --publish 127.0.0.1:5432:5432 postgres
```

It will show a log console. Leave it open for debugging

### Get Network IP for postgres 

```js
docker inspect feed-sys-postgres
```

Get to the end, and pick up the  "IPAddress": "172.17.0.3" (maybe different from  the given example)

```js
            "MacAddress": "02:42:ac:11:00:03",
            "Networks": {
                "bridge": {
                    "IPAMConfig": null,
                    "Links": null,
                    "Aliases": null,
                    "NetworkID": "bd2654aa417b2573149d80ca443177cd3ad37656bb4421374bdd1b4a4ceb1187",
                    "EndpointID": "da2365799cdaba7bf7aa05510277b16b3a5e64b7087b128f187d4894b916832b",
                    "Gateway": "172.17.0.1",
                    "IPAddress": "172.17.0.3",
                    "IPPrefixLen": 16,
                    "IPv6Gateway": "",
                    "GlobalIPv6Address": "",
                    "GlobalIPv6PrefixLen": 0,
                    "MacAddress": "02:42:ac:11:00:03",
                    "DriverOpts": null
                }
            }
```


### Set local ENV

Fill the following ENV values in the ENV file you created for your own setting in BigBang root folder

```
DB_USER=<postgres_user>
DB_PASSWORD=<postgres_password>
DB_NAME=<db_name>
DB_HOST=172.17.0.3 
```

## Set Up local BigBang Docker


### Create BigBang Docker Image

Run the following command in the root of BigBang

```js
docker build -t bigbang . 
```

### Create BigBang Container

Run the following command in the root of BigBang with ENV file <env_file> loaded 

```js
 sudo docker run --name bigbangimage -ti  --net=host -e DISPLAY —env-file <env_file> -v ~/ventureum_projects/go_projects/src/BigBang:/go/src/BigBang bigbang:latest
```


### Run BigBang Container 

Run the following command to get the CONTAINER ID (<container_id>) for the BigBang image 

```js
docker ps --filter "name=bigbangimage"
```

Then Run the container with the CONTAINER ID <container_id>

```js
sudo docker exec -it <container_id> bash 
```

Then you will be directed to /go root folder.  `cd /go/src/`
