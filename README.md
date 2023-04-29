# Pawsitively Purrfect

## Table of Contents

1. [Introduction](#introduction)
2. [Architecture](#architecture)
    1. [System Components](#system-components)
    2. [System Design](#system-design)
    3. [Data Flow](#data-flow)
3. [Data Model](#data-model)
    1. [Entities and Relationships](#entities-and-relationships)
    2. [Schema and Sample Data](#schema-and-sample-data)
4. [API Documentation](#api-documentation)
    1. [Endpoints](#endpoints)
    2. [Rest APIs](#rest-apis)
    3. [GraphQL API](#graphql-api)
5. [Codebase](#codebase)
6. [Deployment / Installation](#deployment--installation)
   1. [Local Development](#local-development)
   2. [Production Environment](#production-environment)
7. [Conclusion](#conclusion)
   1. [Summary](#summary)
   2. [Future Plans and Enhancements](#future-plans-and-enhancements)


## Introduction
**_Pawsitively Purrfect_** is a project aimed at providing an online platform for pet adoption. This platform is designed to help pet lovers find their perfect furry companions while also providing a way for animal shelters and rescue organizations to showcase their available pets.

The "Pawsitively Purrfect" platform provides a ~~user-friendly~~ interface for browsing pets based on different search criteria such as pet type, breed, location, and more. Users can view detailed information about each pet and contact the pet's owner or the organization responsible for its care.

In addition to helping pets find their forever homes, "Pawsitively Purrfect" also provides a way for animal shelters to manage their available pets and adoption applications. With this platform, they can easily add new pets, update pet information, and track adoption applications.

## Architecture
### System Components
The System consists of three major Components mainly.

1. **_Backend Server_**: The backend server is built using Golang and GraphQL, with the Flamego web framework. It consists of multiple layers, including:
    - **_Resolver_**: Handles incoming GraphQL queries and mutations, and maps them to specific service functions.
    - **_Service_**: Implements business logic and interacts with repositories.
    - **_Repository_**: Provides an abstraction layer between the service layer and the database layer.
2. **_Databases_**: The system uses both NoSQL and SQL databases. ArangoDB is used as a NoSQL database, and Postgres is used as a SQL database, implemented using gRPC.
3. **_Frontend_**: The frontend is built using Tailwind CSS and includes basic pages such as login, register, and profile pages.

### System Design

The system uses a layered architecture that separates concerns and ensures loose coupling between the different components. It includes the following layers:

- **_Frontend_**: The frontend layer provides the user interface for the system and the main interface is yet to be built.
- **_GraphQL API_**: The GraphQL API layer handles incoming requests from the frontend and translates them into queries and mutations that can be executed by the backend.
- **_Resolver_**: The resolver layer maps the incoming GraphQL requests to their corresponding service methods.
- **_Service_**: The service layer contains the business logic of the system and performs the necessary operations on the data.
- **_Repository_**: The repository layer provides an abstraction layer over the database and handles the storage and retrieval of data.
- **_Database_**: The database layer stores the data used by the system and is responsible for ensuring its consistency and integrity.

This layered architecture enables the system to be easily extensible and maintainable, as changes to one layer do not affect the others.

### Data Flow
Data flows through our system as follows:

1. The frontend client sends a request to the backend server using GraphQL.
2. The request is received by the GraphQL API layer, which validates and parses the request.
3. The request is then passed to the resolver layer, which maps the request to the corresponding service methods.
4. The service layer contains the business logic of the system and performs the necessary operations on the data.
5. The repository layer provides an abstraction layer over the database and handles the storage and retrieval of data.
6. The database layer stores the data used by the system and is responsible for ensuring its consistency and integrity.
7. The requested data is retrieved from the database and returned to the service layer.
8. The service layer processes the data and returns the response to the resolver layer.
9. The resolver layer maps the response to the GraphQL schema and returns it to the GraphQL API layer.
10. The GraphQL API layer sends the response back to the frontend client.

## Data Model
The "Pawsitively Purrfect" project uses a relational data model to store information about users, pets, shelters, and pet adoptions. There's also option to use NoSQL type database instead of SQL type database.

### Entities and Relationships
There are four main entities in the data model: User, Pet, Shelter, and Pet Adoption. The entities have various attributes that define them, and have different relationships between them.

Here's a breakdown of the different entities and their relationships:

- User Entity:<br/>
    The User entity represents the users of the platform. User has various attributes such as their full name, email, username, location etc. Each User can adopt multiple Pets from a Shelter which is also owned by a User.

- Shelter Entity:<br/>
    The Shelter entity represents the shelters on the platform. Shelter has various attributes such as its name, description, contact information etc. Each shelter can have only one owner, represented by the OwnerID attribute. Each shelter can own multiple pets.

- Pet Entity:<br/>
    The Pet entity represents the pets on the platform. Pet has various attributes such as its name, type, breed, adoption status etc. The AdoptionStatus attribute represents whether the pet is available for adoption or has already been adopted. Each pet belongs to one shelter, represented by the ShelterID attribute. However, a pet can be adopted by a user, which is represented by a PetAdoption record.

- PetAdoption Entity:<br/>
    The PetAdoption entity represents the adoption of pets by users. Each PetAdoption record has a unique ID and is associated with a specific pet and a specific user.

### Schema and Sample Data
JSON Schemas for the entities used in the "Pawsitively Purrfect" project:

User:
```json
{
  "type": "object",
  "properties": {
    "id": {"type": "string"},
    "firstName": {"type": "string"},
    "lastName": {"type": "string"},
    "bio": {"type": "string"},
    "location": {"type": "string"},
    "avatar": {"type": "string"},
    "username": {"type": "string"},
    "email": {"type": "string"},
    "passwordHash": {"type": "string"},
    "isActive": {"type": "boolean"},
    "isAdmin": {"type": "boolean"},
    "createdUnix": {"type": "integer"},
    "updatedUnix": {"type": "integer"},
    "lastLoginUnix": {"type": "integer"}
  }
}
```

Shelter:
```json
{
  "type": "object",
  "properties": {
    "id": {"type": "string"},
    "name": {"type": "string"},
    "description": {"type": "string"},
    "website": {"type": "string"},
    "location": {"type": "string"},
    "contactInformation": {"type": "string"},
    "logo": {"type": "string"},
    "numberOfPets": {"type": "integer"},
    "ownerID": {"type": "string"}
  }
}
```

Pet:
```json
{
  "type": "object",
  "properties": {
    "id": {"type": "string"},
    "name": {"type": "string"},
    "type": {"type": "string"},
    "breed": {"type": "string"},
    "gender": {"type": "string"},
    "photo": {"type": "string"},
    "adoptionStatus": {"type": "string"},
    "shelterID": {"type": "string"}
  }
}
```

PetAdoption:
```json
{
  "type": "object",
  "properties": {
    "id": {"type": "string"},
    "petID": {"type": "string"},
    "userID": {"type": "string"}
  }
}
```

And here's an example of some sample data that could be inserted into these tables:

User:
| ID | FirstName | LastName | Bio | Location      | Avatar | Username | Email                | PasswordHash | IsActive | IsAdmin | CreatedUnix | UpdatedUnix | LastLoginUnix |
|----|-----------|----------|-----|---------------|--------|----------|----------------------|--------------|----------|---------|-------------|-------------|---------------|
| 1  | John      | Doe      | I love dogs! | New York, NY | john.jpg | johnd  | john@pawsshelter.com | 1234567890   | true     | false   | 1620187800  | 1620187800  | 1620187800    |
| 2  | Jane      | Smith    | I am a cat lover! | Los Angeles, CA | jane.jpg | janes  | jane@furryfriends.org | 0987654321   | true     | false   | 1620187800  | 1620187800  | 1620187800    |
| 3  | Bob       | Johnson  | I love all animals! | Seattle, WA | bob.jpg  | bobj   | bob@happytailsrescue.org | 1357908642 | true     | true    | 1620187800  | 1620187800  | 1620187800    |


Shelter:
| ID | Name  | Description | Website | Location | ContactInformation | Logo | NumberOfPets | OwnerID |
|----|-------|-------------|---------|----------|--------------------|------|--------------|---------|
| 1  | PAWS  | Animal rescue organization | www.pawsshelter.com | New York, NY | contact@pawsshelter.com | pawsshelterlogo.jpg | 100 | 1 |
| 2  | Furry Friends | Purrfect home for your furry friends | www.furryfriends.org | Los Angeles, CA | contact@furryfriends.org | furryfriendslogo.jpg | 75 | 2 |
| 3  | Happy Tails  | Helping pets find their forever homes | www.happytailsrescue.org | Seattle, WA | contact@happytailsrescue.org | happytailslogo.jpg | 50 | 3 |


Pet:
| ID | Name | Type | Breed | Gender | Photo | AdoptionStatus | ShelterID |
|----|------|------|-------|--------|-------|----------------|----------|
| 1  | Max  | Dog  | Labrador Retriever | Male   | maxphoto.jpg  | Available    | 1 |
| 2  | Bella | Cat  | Siamese           | Female | bellaphoto.jpg | Adopted        | 1 |
| 3  | Lucy | Dog   | Bulldog          | Female | lucyphoto.jpg  | Available    | 2 |
| 4  | Simba | Cat  | Maine Coon       | Male   | simbaphoto.jpg | Available    | 2 |
| 5  | Charlie | Dog  | Golden Retriever | Male  | charliephoto.jpg | Adopted     | 3 |


PetAdoption:
| ID | PetID | UserID |
|----|-------|--------|
| 1  | 2     | 1      |
| 2  | 5     | 2      |
| 3  | 3     | 1      |


## API Documentation

The Pawsitively Purrfect is a GraphQL server that provides a set of endpoints for managing users, pets, shelters, and pet adoptions. The API allows clients to query and mutate data using a single GraphQL endpoint.

### Endpoints
The Pawsitively Purrfect API has the following endpoints:

#### Rest APIs
The  server some Rest APIs to view the website documentation, login, register and view profile page from UI.

- `/docs` - Serves the GraphQL documentation
- `/user/login` - Serves Login Page
- `/user/register` - Serves Registration Page

- `/logout` - Logout page
- `/{name}` - User profile page

#### GraphQL API
- `/graphql` - The main endpoint for the GraphQL server.

- **_You can find the full GraphQL documentation in_** <a href="https://pawsitively-purrfect.tiiny.site" target="_blank">**_pawsitively-purrfect.tiiny.site_**</a>.

## Codebase
<!--
- Overview of the codebase
- Description of the different modules and packages used in the project
- Explanation of the different functions and their roles in the system
- Explanation of any notable design patterns or frameworks used in the project
-->

<details>
<summary>
Project Structures:
</summary>

```
- .github/
    - workflows/
        - ci.yml
        - release.yml
- api/
    - graphql/
        - resolvers/
            - resolver.go
            - pet.go
            - shelter.go
            - user.go
        - schema/
            - params.go
            - schema.go
    - http/
        - handlers/
            - graphql.go
            - user.go
        - middlewares/
            - auth.go
            - context.go
            - rate_limit.go
            - session_csrf.go
        - routes.go
- cmd/
    - config.go
    - root.go
    - serve.go
    - grpc.go
- configs/
    - .pawsitively-purrfect.yaml
    - config.go
- infra/
    - database/
        - nosql/
            - arangodb/
                - arangodb.go
            - mongodb/
            - mock/
                - mock.go
            - database.go
        - sql/
            - postgres/
                - pb/
                    - postgres_grpc.pb.go
                    - posgres.pb.go
                - server/
                    - postgres.go
                    - health.go
                    - table_sync.go
            - database.go
    - logr/
        - logger.go
- models/
    - gqtypes/
        - pet.go
        - shelter.go
        - types.go
        - user.go
        - pet_adoption.go
    - errors.go
    - pet.go
    - shelter.go
    - user.go
    - pet_adoption.go
- pkg/
    - decode.go
    - hash.go
    - path.go
- proto/
    - database/
        - postgres.proto
- repos/
    - pet/
        - pet.go
        - pet_mock.go
    - shelter/
        - shelter.go
        - shelter_mock.go
    - user/
        - user.go
        - user_mock.go
    - pet_adoption/
        - pet_adoption.go
        - pet_adoption_mock.go
    - pet.go
    - shelter.go
    - user.go
    - pet_adoption.go
- services/
    - all/
        - all.go
    - pet/
        - pet.go
    - shelter/
        - shelter.go
    - user/
        - user.go
    - pet.go
    - shelter.go
    - user.go
- templates/
    - images/
    - logo/

    - docs.tmpl
    - login.tmpl
    - register.tmpl
    - profile.tmpl
    - templates.go
- main.go

- docker-compose.yml
- Dockerfile
- Makefile
- README.md
```
</details>


## Deployment / Installation
We can use either ArangoDB (NoSQL) or Postgres (SQL) as database backend for the `Pawsitively Purrfect` Application.

The postgres database layer is served with gRPC server. So, if we wish to use Postgres as our database, we will have to run an extra gRPC server to serve the Postgres db.

### Local Development
For local development mode, we first have to clone the repository from Github.
- Clone the repository
    ```bash
    $ mkdir -p $HOME/go/src/github.com/masudur-rahman
    $ git clone git@github.com:masudur-rahman/pawsitively-purrfect.git

    $ cd pawsitively-purrfect
    ```
- To get the best experience with the login consistency, Run the following command
    ```bash
    $ sudo echo '127.0.0.1 pawsitively.purrfect' >> /etc/hosts
    ```
- Run the `Pawsitively Purrfect` Application <br/>
    We can start the server following either of the following processes.
    - Running without gRPC server
        ```bash
        $ make run
        $ # It actually runs `docker compose up` command
        $ # ArangoDB is used as the database
        ```
    - Running with gRPC server
        ```bash
        $ make run-with-grpc
        $ # It actually runs `docker compose up --file docker-compose-grpc.yml
        $ # Postgres is used as thye database but the postgres is served through a gRPC server
        ```
    The `Pawsitively Purrfect` application should be up and running. To access it head out to http://pawsitively.purrfect:62783

### Production Environment
To deploy `Pawsitively Purrfect` applicaiton in production environment, the preferred way is through Helm Chart.

First you need to add the repo for the helm chart.
```bash
$ helm repo add masud https://masudur-rahman.github.io/helm-charts/stable
$ helm repo update

$ helm search repo masud/pawsitively-purrfect
```

Just like running application in local environment, we have two installation procedures here too.
- Installing without gRPC server
    ```bash
    $ helm upgrade --install pawsitively-purrfect masud/pawsitively-purrfect -n purrfect --create-namespace
    ```

- Installing with gRPC server
    ```bash
    $ helm upgrade --install pawsitively-purrfect masud/pawsitively-purrfect -n purrfect \
        --create-namespace  --set grpc.enabled=true
    ```

- Verify Installation
To check if `Pawsitively Purrfect` is installed, run the following command:
    ```bash
    $ kubectl get pods -n purrfect -l "app.kubernetes.io/instance=pawsitively-purrfect"

    NAME                                             READY   STATUS    RESTARTS   AGE
    pawsitively-purrfect-698f968b44-d5mt6            1/1     Running   0          15s
    pawsitively-purrfect-arangodb-78db5b45bf-2wklw   1/1     Running   0          10s
    ```

To see the detailed configuration options, visit [here](https://github.com/masudur-rahman/helm-charts/tree/main/charts/pawsitively-purrfect).

<!--
- Description of the hosting and infrastructure requirements
-->

<!--
## Maintenance and Support
- Explanation of how to maintain and support the project in a production environment
- Information on how to debug and troubleshoot common issues
- Contact information for support and feedback
-->

## Conclusion
### Summary
The `Pawsitively Purrfect` is a GraphQL API written in Go that provides pet adoption functionalities such as creating and updating pet profiles, managing pet adoption applications, and enabling communication between adopters and shelter staff.

The API uses either ArangoDB or Postgres as its database backend and can be run with or without a gRPC server.

### Future Plans and Enhancements:
In the future, the application can be further improved by adding support for must columns in SQL database implementations, which will allow updating columns with zero values. Additionally, more queries and mutations can be added to the API to enhance its functionalities further. Finally, a frontend can be developed to provide a user-friendly interface for interacting with the API.
