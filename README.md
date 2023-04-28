# Pawsitively Purrfect

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

The Pawsitively Purrfect is a GraphQL server that provides a set of endpoints for managing pets, shelters, and pet adoptions. The API allows clients to query and mutate data using a single endpoint.

### Endpoints
The Pawsitively Purrfect API has the following endpoints:

GraphQL API
- /graphql - the main endpoint for the GraphQL server.

Rest APIs
- /docs - serves the documentation
- /user/login - Serves Login Page
- /user/register - Serves Registration Page

- /logout - Logout page
- /{name} - user profile page


<html class="no-js" lang="en">
  <body id="spectaql">
    <div id="page" class="drawer-layout">
      <div id="sidebar">
        <div class="sidebar-top-container">
          <div id="logo">
            <img src="https://lh5.googleusercontent.com/9ZQCJ7yj0nccSqTTk-euc5Q7qzc5uKrsoNBD0zJ6trV-GSs7t68f-ZlxqEeKyglihTA=w2400" title="Pawsitively Purrfect GraphQL API Reference" />
          </div>
        </div>
        <nav id="nav" role="navigation">
          <div class="nav-group">
            <h4 class="nav-group-title">Introduction</h4>
            <ul class="nav-group-items">
              <li><a href="#introduction">Welcome</a></li>
            </ul>
          </div>
          <div class="nav-group">
            <h4 class="nav-group-title">Operations</h4>
            <ul class="nav-group-items">
              <li class="nav-group-section">
                <h5 class="nav-group-section-title">
                  <a href="#group-Operations-Queries">Queries</a>
                </h5>
                <ul class="nav-group-section-items">
                  <li><a href="#query-findPets">findPets</a></li>
                  <li><a href="#query-listPets">listPets</a></li>
                  <li><a href="#query-listShelterPets">listShelterPets</a></li>
                  <li><a href="#query-listShelters">listShelters</a></li>
                  <li><a href="#query-pet">pet</a></li>
                  <li><a href="#query-profile">profile</a></li>
                  <li><a href="#query-shelter">shelter</a></li>
                  <li><a href="#query-user">user</a></li>
                </ul>
              </li>
              <li class="nav-group-section">
                <h5 class="nav-group-section-title">
                  <a href="#group-Operations-Mutations">Mutations</a>
                </h5>
                <ul class="nav-group-section-items">
                  <li><a href="#mutation-addPet">addPet</a></li>
                  <li><a href="#mutation-addShelter">addShelter</a></li>
                  <li><a href="#mutation-adoptPet">adoptPet</a></li>
                  <li><a href="#mutation-deleteShelter">deleteShelter</a></li>
                  <li><a href="#mutation-login">login</a></li>
                  <li><a href="#mutation-register">register</a></li>
                  <li><a href="#mutation-updatePet">updatePet</a></li>
                  <li><a href="#mutation-updateProfile">updateProfile</a></li>
                  <li><a href="#mutation-updateShelter">updateShelter</a></li>
                </ul>
              </li>
              <li><a href="#"></a></li>
            </ul>
          </div>
          <div class="nav-group">
            <h4 class="nav-group-title">Types</h4>
            <ul class="nav-group-items">
              <li><a href="#definition-AdoptionStatus">AdoptionStatus</a></li>
              <li><a href="#definition-Boolean">Boolean</a></li>
              <li><a href="#definition-Gender">Gender</a></li>
              <li><a href="#definition-ID">ID</a></li>
              <li><a href="#definition-Int">Int</a></li>
              <li><a href="#definition-Pet">Pet</a></li>
              <li><a href="#definition-PetType">PetType</a></li>
              <li><a href="#definition-Shelter">Shelter</a></li>
              <li><a href="#definition-String">String</a></li>
              <li><a href="#definition-User">User</a></li>
            </ul>
          </div>
        </nav>
      </div>
      <div id="docs">
        <div id="mobile-navbar">
          <button class="sidebar-open-button" type="button">
            <span class="hamburger"></span>
            <span class="sr-only">All topics</span>
          </button>
        </div>
        <article id="content">
          <h1 class="doc-heading">Pawsitively Purrfect GraphQL API Reference</h1>
          <div id="introduction" data-traverse-target="introduction">
            <div id="welcome" class="doc-row">
              <div class="doc-copy">
                <p>Welcome to Pawsitively Purrfect..!</p>
              </div>
              <div class="doc-examples">
                <div class="example-section welcome-contact-section">
                  <h5>Contact</h5>
                  <p class="contact-name">API Support</p>
                  <p class="contact-email"><a href="mailto:masudjuly02@gmail.com">masudjuly02@gmail.com</a></p>
                </div>
                <div class="example-section welcome-license-section">
                  <h5>License</h5>
                  <p class="license-name">Apache 2.0</p>
                  <p class="license-url"><a href="https://www.apache.org/licenses/LICENSE-2.0.html">https://www.apache.org/licenses/LICENSE-2.0.html</a></p>
                </div>
                <div class="example-section welcome-api-endpoints-section example-section-is-code">
                  <h5>API Endpoints</h5>
                  <pre><code># Development:
http://pawsitively.purrfect:62783/graphql
# Production:
http://pawsitively.purrfect:62783/graphql
</code></pre>
                </div>
              </div>
            </div>
          </div>
          <h1 id="group-Operations-Queries" class="group-heading" data-traverse-target="group-Operations-Queries">Queries</h1>
          <section id="query-findPets" class="operation operation-query" data-traverse-target="query-findPets">
            <h2 class="operation-heading ">
              <code>findPets</code>
            </h2>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-description doc-copy-section">
                  <h5>Description</h5>
                  <p>Find pets accross platform based on search criteria</p>
                </div>
              </div>
            </div>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-response doc-copy-section">
                  <h5>Response</h5>
                  <p> Returns <a href="#definition-Pet"><code>[Pet]</code></a>
                  </p>
                </div>
                <div class="operation-arguments doc-copy-section">
                  <h5>Arguments</h5>
                  <table>
                    <thead>
                      <tr>
                        <th>Name</th>
                        <th>Description</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr>
                        <td>
                          <span class="property-name"><code>breed</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>gender</code></span> - <span class="property-type"><a href="#definition-Gender"><code>Gender</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>shelterID</code></span> - <span class="property-type"><a href="#definition-ID"><code>ID</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>adoptionStatus</code></span> - <span class="property-type"><a href="#definition-AdoptionStatus"><code>AdoptionStatus</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>name</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>type</code></span> - <span class="property-type"><a href="#definition-PetType"><code>PetType</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
              <div class="doc-examples">
                <h4 class="example-heading">Example</h4>
                <div class="example-section example-section-is-code operation-query-example">
                  <h5>Query</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-gql"><span class="hljs-symbol"><span class="hljs-keyword">query</span> FindPets<span class="hljs-tag">(
  <span class="hljs-code">$breed</span>:<span class="hljs-type"> String,</span>
  <span class="hljs-code">$gender</span>:<span class="hljs-type"> Gender,</span>
  <span class="hljs-code">$shelterID</span>: ID,
  <span class="hljs-code">$adoptionStatus</span>:<span class="hljs-type"> AdoptionStatus,</span>
  <span class="hljs-code">$name</span>:<span class="hljs-type"> String,</span>
  <span class="hljs-code">$type</span>:<span class="hljs-type"> PetType
</span>)</span> <span class="hljs-tag">{
  <span class="hljs-symbol">findPets<span class="hljs-tag">(
    breed: <span class="hljs-code">$breed</span>,
    gender: <span class="hljs-code">$gender</span>,
    shelterID: <span class="hljs-code">$shelterID</span>,
    adoptionStatus: <span class="hljs-code">$adoptionStatus</span>,
    name: <span class="hljs-code">$name</span>,
    type: <span class="hljs-code">$type</span>
  )</span> <span class="hljs-tag">{
    <span class="hljs-symbol">adoptionStatus</span>
    <span class="hljs-symbol">breed</span>
    <span class="hljs-symbol">currentOwnerID</span>
    <span class="hljs-symbol">gender</span>
    <span class="hljs-symbol">id</span>
    <span class="hljs-symbol">name</span>
    <span class="hljs-symbol">photo</span>
    <span class="hljs-symbol">shelterID</span>
    <span class="hljs-symbol">type</span>
  }</span></span>
}</span></span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-variables-example">
                  <h5>Variables</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span>
  <span class="hljs-attr">"breed"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"gender"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Male"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"shelterID"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"4"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"adoptionStatus"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Available"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"name"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"type"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Cat"</span>
<span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-response-example">
                  <h5>Response</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span>
  <span class="hljs-attr">"data"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">{</span>
    <span class="hljs-attr">"findPets"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">[</span>
      <span class="hljs-punctuation">{</span>
        <span class="hljs-attr">"adoptionStatus"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Available"</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"breed"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"currentOwnerID"</span><span class="hljs-punctuation">:</span> <span class="hljs-number">4</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"gender"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Male"</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"id"</span><span class="hljs-punctuation">:</span> <span class="hljs-number">4</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"name"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"photo"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"shelterID"</span><span class="hljs-punctuation">:</span> <span class="hljs-number">4</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"type"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Cat"</span>
      <span class="hljs-punctuation">}</span>
    <span class="hljs-punctuation">]</span>
  <span class="hljs-punctuation">}</span>
<span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
              </div>
            </div>
          </section>
          <section id="query-listPets" class="operation operation-query" data-traverse-target="query-listPets">
            <div class="operation-group-name">
              <a href="#group-Operations-Queries">Queries</a>
            </div>
            <h2 class="operation-heading ">
              <code>listPets</code>
            </h2>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-description doc-copy-section">
                  <h5>Description</h5>
                  <p>List all pets owned by logged-in user</p>
                </div>
              </div>
            </div>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-response doc-copy-section">
                  <h5>Response</h5>
                  <p> Returns <a href="#definition-Pet"><code>[Pet]</code></a>
                  </p>
                </div>
              </div>
              <div class="doc-examples">
                <h4 class="example-heading">Example</h4>
                <div class="example-section example-section-is-code operation-query-example">
                  <h5>Query</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-gql"><span class="hljs-symbol"><span class="hljs-keyword">query</span> ListPets <span class="hljs-tag">{
  <span class="hljs-symbol">listPets <span class="hljs-tag">{
    <span class="hljs-symbol">adoptionStatus</span>
    <span class="hljs-symbol">breed</span>
    <span class="hljs-symbol">currentOwnerID</span>
    <span class="hljs-symbol">gender</span>
    <span class="hljs-symbol">id</span>
    <span class="hljs-symbol">name</span>
    <span class="hljs-symbol">photo</span>
    <span class="hljs-symbol">shelterID</span>
    <span class="hljs-symbol">type</span>
  }</span></span>
}</span></span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-response-example">
                  <h5>Response</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span>
  <span class="hljs-attr">"data"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">{</span>
    <span class="hljs-attr">"listPets"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">[</span>
      <span class="hljs-punctuation">{</span>
        <span class="hljs-attr">"adoptionStatus"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Available"</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"breed"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"currentOwnerID"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"4"</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"gender"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Male"</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"id"</span><span class="hljs-punctuation">:</span> <span class="hljs-number">4</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"name"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"photo"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"shelterID"</span><span class="hljs-punctuation">:</span> <span class="hljs-number">4</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"type"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Cat"</span>
      <span class="hljs-punctuation">}</span>
    <span class="hljs-punctuation">]</span>
  <span class="hljs-punctuation">}</span>
<span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
              </div>
            </div>
          </section>
          <section id="query-listShelterPets" class="operation operation-query" data-traverse-target="query-listShelterPets">
            <div class="operation-group-name">
              <a href="#group-Operations-Queries">Queries</a>
            </div>
            <h2 class="operation-heading ">
              <code>listShelterPets</code>
            </h2>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-description doc-copy-section">
                  <h5>Description</h5>
                  <p>List pets by shelter</p>
                </div>
              </div>
            </div>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-response doc-copy-section">
                  <h5>Response</h5>
                  <p> Returns <a href="#definition-Pet"><code>[Pet]</code></a>
                  </p>
                </div>
                <div class="operation-arguments doc-copy-section">
                  <h5>Arguments</h5>
                  <table>
                    <thead>
                      <tr>
                        <th>Name</th>
                        <th>Description</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr>
                        <td>
                          <span class="property-name"><code>shelterID</code></span> - <span class="property-type"><a href="#definition-ID"><code>ID!</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>adoptionStatus</code></span> - <span class="property-type"><a href="#definition-AdoptionStatus"><code>AdoptionStatus</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
              <div class="doc-examples">
                <h4 class="example-heading">Example</h4>
                <div class="example-section example-section-is-code operation-query-example">
                  <h5>Query</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-gql"><span class="hljs-symbol"><span class="hljs-keyword">query</span> ListShelterPets<span class="hljs-tag">(
  <span class="hljs-code">$shelterID</span>: ID!,
  <span class="hljs-code">$adoptionStatus</span>:<span class="hljs-type"> AdoptionStatus
</span>)</span> <span class="hljs-tag">{
  <span class="hljs-symbol">listShelterPets<span class="hljs-tag">(
    shelterID: <span class="hljs-code">$shelterID</span>,
    adoptionStatus: <span class="hljs-code">$adoptionStatus</span>
  )</span> <span class="hljs-tag">{
    <span class="hljs-symbol">adoptionStatus</span>
    <span class="hljs-symbol">breed</span>
    <span class="hljs-symbol">currentOwnerID</span>
    <span class="hljs-symbol">gender</span>
    <span class="hljs-symbol">id</span>
    <span class="hljs-symbol">name</span>
    <span class="hljs-symbol">photo</span>
    <span class="hljs-symbol">shelterID</span>
    <span class="hljs-symbol">type</span>
  }</span></span>
}</span></span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-variables-example">
                  <h5>Variables</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span>
  <span class="hljs-attr">"shelterID"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"4"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"adoptionStatus"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Available"</span>
<span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-response-example">
                  <h5>Response</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span>
  <span class="hljs-attr">"data"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">{</span>
    <span class="hljs-attr">"listShelterPets"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">[</span>
      <span class="hljs-punctuation">{</span>
        <span class="hljs-attr">"adoptionStatus"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Available"</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"breed"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"currentOwnerID"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"4"</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"gender"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Male"</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"id"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"4"</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"name"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"photo"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"shelterID"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"4"</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"type"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Cat"</span>
      <span class="hljs-punctuation">}</span>
    <span class="hljs-punctuation">]</span>
  <span class="hljs-punctuation">}</span>
<span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
              </div>
            </div>
          </section>
          <section id="query-listShelters" class="operation operation-query" data-traverse-target="query-listShelters">
            <div class="operation-group-name">
              <a href="#group-Operations-Queries">Queries</a>
            </div>
            <h2 class="operation-heading ">
              <code>listShelters</code>
            </h2>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-description doc-copy-section">
                  <h5>Description</h5>
                  <p>List shelters by filter</p>
                </div>
              </div>
            </div>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-response doc-copy-section">
                  <h5>Response</h5>
                  <p> Returns <a href="#definition-Shelter"><code>[Shelter]</code></a>
                  </p>
                </div>
                <div class="operation-arguments doc-copy-section">
                  <h5>Arguments</h5>
                  <table>
                    <thead>
                      <tr>
                        <th>Name</th>
                        <th>Description</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr>
                        <td>
                          <span class="property-name"><code>description</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>website</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>location</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>ownerID</code></span> - <span class="property-type"><a href="#definition-ID"><code>ID</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>name</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
              <div class="doc-examples">
                <h4 class="example-heading">Example</h4>
                <div class="example-section example-section-is-code operation-query-example">
                  <h5>Query</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-gql"><span class="hljs-symbol"><span class="hljs-keyword">query</span> ListShelters<span class="hljs-tag">(
  <span class="hljs-code">$description</span>:<span class="hljs-type"> String,</span>
  <span class="hljs-code">$website</span>:<span class="hljs-type"> String,</span>
  <span class="hljs-code">$location</span>:<span class="hljs-type"> String,</span>
  <span class="hljs-code">$ownerID</span>: ID,
  <span class="hljs-code">$name</span>:<span class="hljs-type"> String
</span>)</span> <span class="hljs-tag">{
  <span class="hljs-symbol">listShelters<span class="hljs-tag">(
    description: <span class="hljs-code">$description</span>,
    website: <span class="hljs-code">$website</span>,
    location: <span class="hljs-code">$location</span>,
    ownerID: <span class="hljs-code">$ownerID</span>,
    name: <span class="hljs-code">$name</span>
  )</span> <span class="hljs-tag">{
    <span class="hljs-symbol">contactInformation</span>
    <span class="hljs-symbol">description</span>
    <span class="hljs-symbol">id</span>
    <span class="hljs-symbol">location</span>
    <span class="hljs-symbol">logo</span>
    <span class="hljs-symbol">name</span>
    <span class="hljs-symbol">numberOfPets</span>
    <span class="hljs-symbol">ownerID</span>
    <span class="hljs-symbol">website</span>
  }</span></span>
}</span></span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-variables-example">
                  <h5>Variables</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span>
  <span class="hljs-attr">"description"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"website"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"location"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"ownerID"</span><span class="hljs-punctuation">:</span> <span class="hljs-number">4</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"name"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span>
<span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-response-example">
                  <h5>Response</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span>
  <span class="hljs-attr">"data"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">{</span>
    <span class="hljs-attr">"listShelters"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">[</span>
      <span class="hljs-punctuation">{</span>
        <span class="hljs-attr">"contactInformation"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"description"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"id"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"4"</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"location"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"logo"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"name"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"numberOfPets"</span><span class="hljs-punctuation">:</span> <span class="hljs-number">123</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"ownerID"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"4"</span><span class="hljs-punctuation">,</span>
        <span class="hljs-attr">"website"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span>
      <span class="hljs-punctuation">}</span>
    <span class="hljs-punctuation">]</span>
  <span class="hljs-punctuation">}</span>
<span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
              </div>
            </div>
          </section>
          <section id="query-pet" class="operation operation-query" data-traverse-target="query-pet">
            <div class="operation-group-name">
              <a href="#group-Operations-Queries">Queries</a>
            </div>
            <h2 class="operation-heading ">
              <code>pet</code>
            </h2>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-description doc-copy-section">
                  <h5>Description</h5>
                  <p>Get a pet by ID.</p>
                </div>
              </div>
            </div>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-response doc-copy-section">
                  <h5>Response</h5>
                  <p> Returns a <a href="#definition-Pet"><code>Pet</code></a>
                  </p>
                </div>
                <div class="operation-arguments doc-copy-section">
                  <h5>Arguments</h5>
                  <table>
                    <thead>
                      <tr>
                        <th>Name</th>
                        <th>Description</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr>
                        <td>
                          <span class="property-name"><code>id</code></span> - <span class="property-type"><a href="#definition-ID"><code>ID!</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
              <div class="doc-examples">
                <h4 class="example-heading">Example</h4>
                <div class="example-section example-section-is-code operation-query-example">
                  <h5>Query</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-gql"><span class="hljs-symbol"><span class="hljs-keyword">query</span> Pet<span class="hljs-tag">(<span class="hljs-code">$id</span>: ID!)</span> <span class="hljs-tag">{
  <span class="hljs-symbol">pet<span class="hljs-tag">(id: <span class="hljs-code">$id</span>)</span> <span class="hljs-tag">{
    <span class="hljs-symbol">adoptionStatus</span>
    <span class="hljs-symbol">breed</span>
    <span class="hljs-symbol">currentOwnerID</span>
    <span class="hljs-symbol">gender</span>
    <span class="hljs-symbol">id</span>
    <span class="hljs-symbol">name</span>
    <span class="hljs-symbol">photo</span>
    <span class="hljs-symbol">shelterID</span>
    <span class="hljs-symbol">type</span>
  }</span></span>
}</span></span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-variables-example">
                  <h5>Variables</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span><span class="hljs-attr">"id"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"4"</span><span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-response-example">
                  <h5>Response</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span>
  <span class="hljs-attr">"data"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">{</span>
    <span class="hljs-attr">"pet"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">{</span>
      <span class="hljs-attr">"adoptionStatus"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Available"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"breed"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"currentOwnerID"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"4"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"gender"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Male"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"id"</span><span class="hljs-punctuation">:</span> <span class="hljs-number">4</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"name"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"photo"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"shelterID"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"4"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"type"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Cat"</span>
    <span class="hljs-punctuation">}</span>
  <span class="hljs-punctuation">}</span>
<span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
              </div>
            </div>
          </section>
          <section id="query-profile" class="operation operation-query" data-traverse-target="query-profile">
            <div class="operation-group-name">
              <a href="#group-Operations-Queries">Queries</a>
            </div>
            <h2 class="operation-heading ">
              <code>profile</code>
            </h2>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-description doc-copy-section">
                  <h5>Description</h5>
                  <p>Get logged-in user profile</p>
                </div>
              </div>
            </div>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-response doc-copy-section">
                  <h5>Response</h5>
                  <p> Returns a <a href="#definition-User"><code>User</code></a>
                  </p>
                </div>
              </div>
              <div class="doc-examples">
                <h4 class="example-heading">Example</h4>
                <div class="example-section example-section-is-code operation-query-example">
                  <h5>Query</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-gql"><span class="hljs-symbol"><span class="hljs-keyword">query</span> Profile <span class="hljs-tag">{
  <span class="hljs-symbol">profile <span class="hljs-tag">{
    <span class="hljs-symbol">avatar</span>
    <span class="hljs-symbol">bio</span>
    <span class="hljs-symbol">email</span>
    <span class="hljs-symbol">firstName</span>
    <span class="hljs-symbol">id</span>
    <span class="hljs-symbol">isActive</span>
    <span class="hljs-symbol">isAdmin</span>
    <span class="hljs-symbol">lastName</span>
    <span class="hljs-symbol">location</span>
    <span class="hljs-symbol">username</span>
  }</span></span>
}</span></span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-response-example">
                  <h5>Response</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span>
  <span class="hljs-attr">"data"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">{</span>
    <span class="hljs-attr">"profile"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">{</span>
      <span class="hljs-attr">"avatar"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"bio"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"email"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"firstName"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"id"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"4"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"isActive"</span><span class="hljs-punctuation">:</span> <span class="hljs-literal"><span class="hljs-keyword">false</span></span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"isAdmin"</span><span class="hljs-punctuation">:</span> <span class="hljs-literal"><span class="hljs-keyword">false</span></span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"lastName"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"location"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"username"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span>
    <span class="hljs-punctuation">}</span>
  <span class="hljs-punctuation">}</span>
<span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
              </div>
            </div>
          </section>
          <section id="query-shelter" class="operation operation-query" data-traverse-target="query-shelter">
            <div class="operation-group-name">
              <a href="#group-Operations-Queries">Queries</a>
            </div>
            <h2 class="operation-heading ">
              <code>shelter</code>
            </h2>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-description doc-copy-section">
                  <h5>Description</h5>
                  <p>Get a shelter by ID.</p>
                </div>
              </div>
            </div>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-response doc-copy-section">
                  <h5>Response</h5>
                  <p> Returns a <a href="#definition-Shelter"><code>Shelter</code></a>
                  </p>
                </div>
                <div class="operation-arguments doc-copy-section">
                  <h5>Arguments</h5>
                  <table>
                    <thead>
                      <tr>
                        <th>Name</th>
                        <th>Description</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr>
                        <td>
                          <span class="property-name"><code>id</code></span> - <span class="property-type"><a href="#definition-ID"><code>ID!</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
              <div class="doc-examples">
                <h4 class="example-heading">Example</h4>
                <div class="example-section example-section-is-code operation-query-example">
                  <h5>Query</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-gql"><span class="hljs-symbol"><span class="hljs-keyword">query</span> Shelter<span class="hljs-tag">(<span class="hljs-code">$id</span>: ID!)</span> <span class="hljs-tag">{
  <span class="hljs-symbol">shelter<span class="hljs-tag">(id: <span class="hljs-code">$id</span>)</span> <span class="hljs-tag">{
    <span class="hljs-symbol">contactInformation</span>
    <span class="hljs-symbol">description</span>
    <span class="hljs-symbol">id</span>
    <span class="hljs-symbol">location</span>
    <span class="hljs-symbol">logo</span>
    <span class="hljs-symbol">name</span>
    <span class="hljs-symbol">numberOfPets</span>
    <span class="hljs-symbol">ownerID</span>
    <span class="hljs-symbol">website</span>
  }</span></span>
}</span></span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-variables-example">
                  <h5>Variables</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span><span class="hljs-attr">"id"</span><span class="hljs-punctuation">:</span> <span class="hljs-number">4</span><span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-response-example">
                  <h5>Response</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span>
  <span class="hljs-attr">"data"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">{</span>
    <span class="hljs-attr">"shelter"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">{</span>
      <span class="hljs-attr">"contactInformation"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"description"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"id"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"4"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"location"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"logo"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"name"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"numberOfPets"</span><span class="hljs-punctuation">:</span> <span class="hljs-number">987</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"ownerID"</span><span class="hljs-punctuation">:</span> <span class="hljs-number">4</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"website"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span>
    <span class="hljs-punctuation">}</span>
  <span class="hljs-punctuation">}</span>
<span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
              </div>
            </div>
          </section>
          <section id="query-user" class="operation operation-query" data-traverse-target="query-user">
            <div class="operation-group-name">
              <a href="#group-Operations-Queries">Queries</a>
            </div>
            <h2 class="operation-heading ">
              <code>user</code>
            </h2>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-description doc-copy-section">
                  <h5>Description</h5>
                  <p>Get a user by ID.</p>
                </div>
              </div>
            </div>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-response doc-copy-section">
                  <h5>Response</h5>
                  <p> Returns a <a href="#definition-User"><code>User</code></a>
                  </p>
                </div>
                <div class="operation-arguments doc-copy-section">
                  <h5>Arguments</h5>
                  <table>
                    <thead>
                      <tr>
                        <th>Name</th>
                        <th>Description</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr>
                        <td>
                          <span class="property-name"><code>id</code></span> - <span class="property-type"><a href="#definition-ID"><code>ID</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>name</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
              <div class="doc-examples">
                <h4 class="example-heading">Example</h4>
                <div class="example-section example-section-is-code operation-query-example">
                  <h5>Query</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-gql"><span class="hljs-symbol"><span class="hljs-keyword">query</span> User<span class="hljs-tag">(
  <span class="hljs-code">$id</span>: ID,
  <span class="hljs-code">$name</span>:<span class="hljs-type"> String
</span>)</span> <span class="hljs-tag">{
  <span class="hljs-symbol">user<span class="hljs-tag">(
    id: <span class="hljs-code">$id</span>,
    name: <span class="hljs-code">$name</span>
  )</span> <span class="hljs-tag">{
    <span class="hljs-symbol">avatar</span>
    <span class="hljs-symbol">bio</span>
    <span class="hljs-symbol">email</span>
    <span class="hljs-symbol">firstName</span>
    <span class="hljs-symbol">id</span>
    <span class="hljs-symbol">isActive</span>
    <span class="hljs-symbol">isAdmin</span>
    <span class="hljs-symbol">lastName</span>
    <span class="hljs-symbol">location</span>
    <span class="hljs-symbol">username</span>
  }</span></span>
}</span></span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-variables-example">
                  <h5>Variables</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span>
  <span class="hljs-attr">"id"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"4"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"name"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span>
<span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-response-example">
                  <h5>Response</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span>
  <span class="hljs-attr">"data"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">{</span>
    <span class="hljs-attr">"user"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">{</span>
      <span class="hljs-attr">"avatar"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"bio"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"email"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"firstName"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"id"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"4"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"isActive"</span><span class="hljs-punctuation">:</span> <span class="hljs-literal"><span class="hljs-keyword">false</span></span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"isAdmin"</span><span class="hljs-punctuation">:</span> <span class="hljs-literal"><span class="hljs-keyword">false</span></span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"lastName"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"location"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"username"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span>
    <span class="hljs-punctuation">}</span>
  <span class="hljs-punctuation">}</span>
<span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
              </div>
            </div>
          </section>
          <h1 id="group-Operations-Mutations" class="group-heading" data-traverse-target="group-Operations-Mutations">Mutations</h1>
          <section id="mutation-addPet" class="operation operation-mutation" data-traverse-target="mutation-addPet">
            <h2 class="operation-heading ">
              <code>addPet</code>
            </h2>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-description doc-copy-section">
                  <h5>Description</h5>
                  <p>Add new pet to a shelter</p>
                </div>
              </div>
            </div>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-response doc-copy-section">
                  <h5>Response</h5>
                  <p> Returns a <a href="#definition-Pet"><code>Pet</code></a>
                  </p>
                </div>
                <div class="operation-arguments doc-copy-section">
                  <h5>Arguments</h5>
                  <table>
                    <thead>
                      <tr>
                        <th>Name</th>
                        <th>Description</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr>
                        <td>
                          <span class="property-name"><code>gender</code></span> - <span class="property-type"><a href="#definition-Gender"><code>Gender</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>shelterID</code></span> - <span class="property-type"><a href="#definition-ID"><code>ID!</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>name</code></span> - <span class="property-type"><a href="#definition-String"><code>String!</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>type</code></span> - <span class="property-type"><a href="#definition-PetType"><code>PetType!</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>breed</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
              <div class="doc-examples">
                <h4 class="example-heading">Example</h4>
                <div class="example-section example-section-is-code operation-query-example">
                  <h5>Query</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-gql"><span class="hljs-symbol"><span class="hljs-keyword">mutation</span> AddPet<span class="hljs-tag">(
  <span class="hljs-code">$gender</span>:<span class="hljs-type"> Gender,</span>
  <span class="hljs-code">$shelterID</span>: ID!,
  <span class="hljs-code">$name</span>:<span class="hljs-type"> String!</span>,
  <span class="hljs-code">$type</span>:<span class="hljs-type"> PetType!</span>,
  <span class="hljs-code">$breed</span>:<span class="hljs-type"> String
</span>)</span> <span class="hljs-tag">{
  <span class="hljs-symbol">addPet<span class="hljs-tag">(
    gender: <span class="hljs-code">$gender</span>,
    shelterID: <span class="hljs-code">$shelterID</span>,
    name: <span class="hljs-code">$name</span>,
    type: <span class="hljs-code">$type</span>,
    breed: <span class="hljs-code">$breed</span>
  )</span> <span class="hljs-tag">{
    <span class="hljs-symbol">adoptionStatus</span>
    <span class="hljs-symbol">breed</span>
    <span class="hljs-symbol">currentOwnerID</span>
    <span class="hljs-symbol">gender</span>
    <span class="hljs-symbol">id</span>
    <span class="hljs-symbol">name</span>
    <span class="hljs-symbol">photo</span>
    <span class="hljs-symbol">shelterID</span>
    <span class="hljs-symbol">type</span>
  }</span></span>
}</span></span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-variables-example">
                  <h5>Variables</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span>
  <span class="hljs-attr">"gender"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Male"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"shelterID"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"4"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"name"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"type"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Cat"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"breed"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span>
<span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-response-example">
                  <h5>Response</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span>
  <span class="hljs-attr">"data"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">{</span>
    <span class="hljs-attr">"addPet"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">{</span>
      <span class="hljs-attr">"adoptionStatus"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Available"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"breed"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"currentOwnerID"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"4"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"gender"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Male"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"id"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"4"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"name"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"photo"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"shelterID"</span><span class="hljs-punctuation">:</span> <span class="hljs-number">4</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"type"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Cat"</span>
    <span class="hljs-punctuation">}</span>
  <span class="hljs-punctuation">}</span>
<span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
              </div>
            </div>
          </section>
          <section id="mutation-addShelter" class="operation operation-mutation" data-traverse-target="mutation-addShelter">
            <div class="operation-group-name">
              <a href="#group-Operations-Mutations">Mutations</a>
            </div>
            <h2 class="operation-heading ">
              <code>addShelter</code>
            </h2>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-description doc-copy-section">
                  <h5>Description</h5>
                  <p>Add new shelter to the system</p>
                </div>
              </div>
            </div>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-response doc-copy-section">
                  <h5>Response</h5>
                  <p> Returns a <a href="#definition-Shelter"><code>Shelter</code></a>
                  </p>
                </div>
                <div class="operation-arguments doc-copy-section">
                  <h5>Arguments</h5>
                  <table>
                    <thead>
                      <tr>
                        <th>Name</th>
                        <th>Description</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr>
                        <td>
                          <span class="property-name"><code>contactInformation</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>name</code></span> - <span class="property-type"><a href="#definition-String"><code>String!</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>description</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>website</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>location</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
              <div class="doc-examples">
                <h4 class="example-heading">Example</h4>
                <div class="example-section example-section-is-code operation-query-example">
                  <h5>Query</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-gql"><span class="hljs-symbol"><span class="hljs-keyword">mutation</span> AddShelter<span class="hljs-tag">(
  <span class="hljs-code">$contactInformation</span>:<span class="hljs-type"> String,</span>
  <span class="hljs-code">$name</span>:<span class="hljs-type"> String!</span>,
  <span class="hljs-code">$description</span>:<span class="hljs-type"> String,</span>
  <span class="hljs-code">$website</span>:<span class="hljs-type"> String,</span>
  <span class="hljs-code">$location</span>:<span class="hljs-type"> String
</span>)</span> <span class="hljs-tag">{
  <span class="hljs-symbol">addShelter<span class="hljs-tag">(
    contactInformation: <span class="hljs-code">$contactInformation</span>,
    name: <span class="hljs-code">$name</span>,
    description: <span class="hljs-code">$description</span>,
    website: <span class="hljs-code">$website</span>,
    location: <span class="hljs-code">$location</span>
  )</span> <span class="hljs-tag">{
    <span class="hljs-symbol">contactInformation</span>
    <span class="hljs-symbol">description</span>
    <span class="hljs-symbol">id</span>
    <span class="hljs-symbol">location</span>
    <span class="hljs-symbol">logo</span>
    <span class="hljs-symbol">name</span>
    <span class="hljs-symbol">numberOfPets</span>
    <span class="hljs-symbol">ownerID</span>
    <span class="hljs-symbol">website</span>
  }</span></span>
}</span></span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-variables-example">
                  <h5>Variables</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span>
  <span class="hljs-attr">"contactInformation"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"name"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"description"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"website"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"location"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span>
<span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-response-example">
                  <h5>Response</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span>
  <span class="hljs-attr">"data"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">{</span>
    <span class="hljs-attr">"addShelter"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">{</span>
      <span class="hljs-attr">"contactInformation"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"description"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"id"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"4"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"location"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"logo"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"name"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"numberOfPets"</span><span class="hljs-punctuation">:</span> <span class="hljs-number">987</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"ownerID"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"4"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"website"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span>
    <span class="hljs-punctuation">}</span>
  <span class="hljs-punctuation">}</span>
<span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
              </div>
            </div>
          </section>
          <section id="mutation-adoptPet" class="operation operation-mutation" data-traverse-target="mutation-adoptPet">
            <div class="operation-group-name">
              <a href="#group-Operations-Mutations">Mutations</a>
            </div>
            <h2 class="operation-heading ">
              <code>adoptPet</code>
            </h2>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-description doc-copy-section">
                  <h5>Description</h5>
                  <p>Adopt pet from a shelter</p>
                </div>
              </div>
            </div>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-response doc-copy-section">
                  <h5>Response</h5>
                  <p> Returns a <a href="#definition-Pet"><code>Pet</code></a>
                  </p>
                </div>
                <div class="operation-arguments doc-copy-section">
                  <h5>Arguments</h5>
                  <table>
                    <thead>
                      <tr>
                        <th>Name</th>
                        <th>Description</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr>
                        <td>
                          <span class="property-name"><code>petID</code></span> - <span class="property-type"><a href="#definition-ID"><code>ID</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
              <div class="doc-examples">
                <h4 class="example-heading">Example</h4>
                <div class="example-section example-section-is-code operation-query-example">
                  <h5>Query</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-gql"><span class="hljs-symbol"><span class="hljs-keyword">mutation</span> AdoptPet<span class="hljs-tag">(<span class="hljs-code">$petID</span>: ID)</span> <span class="hljs-tag">{
  <span class="hljs-symbol">adoptPet<span class="hljs-tag">(petID: <span class="hljs-code">$petID</span>)</span> <span class="hljs-tag">{
    <span class="hljs-symbol">adoptionStatus</span>
    <span class="hljs-symbol">breed</span>
    <span class="hljs-symbol">currentOwnerID</span>
    <span class="hljs-symbol">gender</span>
    <span class="hljs-symbol">id</span>
    <span class="hljs-symbol">name</span>
    <span class="hljs-symbol">photo</span>
    <span class="hljs-symbol">shelterID</span>
    <span class="hljs-symbol">type</span>
  }</span></span>
}</span></span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-variables-example">
                  <h5>Variables</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span><span class="hljs-attr">"petID"</span><span class="hljs-punctuation">:</span> <span class="hljs-number">4</span><span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-response-example">
                  <h5>Response</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span>
  <span class="hljs-attr">"data"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">{</span>
    <span class="hljs-attr">"adoptPet"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">{</span>
      <span class="hljs-attr">"adoptionStatus"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Available"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"breed"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"currentOwnerID"</span><span class="hljs-punctuation">:</span> <span class="hljs-number">4</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"gender"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Male"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"id"</span><span class="hljs-punctuation">:</span> <span class="hljs-number">4</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"name"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"photo"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"shelterID"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"4"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"type"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Cat"</span>
    <span class="hljs-punctuation">}</span>
  <span class="hljs-punctuation">}</span>
<span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
              </div>
            </div>
          </section>
          <section id="mutation-deleteShelter" class="operation operation-mutation" data-traverse-target="mutation-deleteShelter">
            <div class="operation-group-name">
              <a href="#group-Operations-Mutations">Mutations</a>
            </div>
            <h2 class="operation-heading ">
              <code>deleteShelter</code>
            </h2>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-description doc-copy-section">
                  <h5>Description</h5>
                  <p>Delete a shelter</p>
                </div>
              </div>
            </div>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-response doc-copy-section">
                  <h5>Response</h5>
                  <p> Returns a <a href="#definition-Boolean"><code>Boolean</code></a>
                  </p>
                </div>
                <div class="operation-arguments doc-copy-section">
                  <h5>Arguments</h5>
                  <table>
                    <thead>
                      <tr>
                        <th>Name</th>
                        <th>Description</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr>
                        <td>
                          <span class="property-name"><code>id</code></span> - <span class="property-type"><a href="#definition-ID"><code>ID!</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
              <div class="doc-examples">
                <h4 class="example-heading">Example</h4>
                <div class="example-section example-section-is-code operation-query-example">
                  <h5>Query</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-gql"><span class="hljs-symbol"><span class="hljs-keyword">mutation</span> DeleteShelter<span class="hljs-tag">(<span class="hljs-code">$id</span>: ID!)</span> <span class="hljs-tag">{
  <span class="hljs-symbol">deleteShelter<span class="hljs-tag">(id: <span class="hljs-code">$id</span>)</span></span>
}</span></span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-variables-example">
                  <h5>Variables</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span><span class="hljs-attr">"id"</span><span class="hljs-punctuation">:</span> <span class="hljs-number">4</span><span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-response-example">
                  <h5>Response</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span><span class="hljs-attr">"data"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">{</span><span class="hljs-attr">"deleteShelter"</span><span class="hljs-punctuation">:</span> <span class="hljs-literal"><span class="hljs-keyword">false</span></span><span class="hljs-punctuation">}</span><span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
              </div>
            </div>
          </section>
          <section id="mutation-login" class="operation operation-mutation" data-traverse-target="mutation-login">
            <div class="operation-group-name">
              <a href="#group-Operations-Mutations">Mutations</a>
            </div>
            <h2 class="operation-heading ">
              <code>login</code>
            </h2>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-description doc-copy-section">
                  <h5>Description</h5>
                  <p>Login user to the system</p>
                </div>
              </div>
            </div>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-response doc-copy-section">
                  <h5>Response</h5>
                  <p> Returns a <a href="#definition-User"><code>User</code></a>
                  </p>
                </div>
                <div class="operation-arguments doc-copy-section">
                  <h5>Arguments</h5>
                  <table>
                    <thead>
                      <tr>
                        <th>Name</th>
                        <th>Description</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr>
                        <td>
                          <span class="property-name"><code>password</code></span> - <span class="property-type"><a href="#definition-String"><code>String!</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>username</code></span> - <span class="property-type"><a href="#definition-String"><code>String!</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
              <div class="doc-examples">
                <h4 class="example-heading">Example</h4>
                <div class="example-section example-section-is-code operation-query-example">
                  <h5>Query</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-gql"><span class="hljs-symbol"><span class="hljs-keyword">mutation</span> Login<span class="hljs-tag">(
  <span class="hljs-code">$password</span>:<span class="hljs-type"> String!</span>,
  <span class="hljs-code">$username</span>:<span class="hljs-type"> String!</span>
)</span> <span class="hljs-tag">{
  <span class="hljs-symbol">login<span class="hljs-tag">(
    password: <span class="hljs-code">$password</span>,
    username: <span class="hljs-code">$username</span>
  )</span> <span class="hljs-tag">{
    <span class="hljs-symbol">avatar</span>
    <span class="hljs-symbol">bio</span>
    <span class="hljs-symbol">email</span>
    <span class="hljs-symbol">firstName</span>
    <span class="hljs-symbol">id</span>
    <span class="hljs-symbol">isActive</span>
    <span class="hljs-symbol">isAdmin</span>
    <span class="hljs-symbol">lastName</span>
    <span class="hljs-symbol">location</span>
    <span class="hljs-symbol">username</span>
  }</span></span>
}</span></span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-variables-example">
                  <h5>Variables</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span>
  <span class="hljs-attr">"password"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"username"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span>
<span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-response-example">
                  <h5>Response</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span>
  <span class="hljs-attr">"data"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">{</span>
    <span class="hljs-attr">"login"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">{</span>
      <span class="hljs-attr">"avatar"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"bio"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"email"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"firstName"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"id"</span><span class="hljs-punctuation">:</span> <span class="hljs-number">4</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"isActive"</span><span class="hljs-punctuation">:</span> <span class="hljs-literal"><span class="hljs-keyword">false</span></span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"isAdmin"</span><span class="hljs-punctuation">:</span> <span class="hljs-literal"><span class="hljs-keyword">true</span></span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"lastName"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"location"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"username"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span>
    <span class="hljs-punctuation">}</span>
  <span class="hljs-punctuation">}</span>
<span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
              </div>
            </div>
          </section>
          <section id="mutation-register" class="operation operation-mutation" data-traverse-target="mutation-register">
            <div class="operation-group-name">
              <a href="#group-Operations-Mutations">Mutations</a>
            </div>
            <h2 class="operation-heading ">
              <code>register</code>
            </h2>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-description doc-copy-section">
                  <h5>Description</h5>
                  <p>Register a new user to the system</p>
                </div>
              </div>
            </div>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-response doc-copy-section">
                  <h5>Response</h5>
                  <p> Returns a <a href="#definition-User"><code>User</code></a>
                  </p>
                </div>
                <div class="operation-arguments doc-copy-section">
                  <h5>Arguments</h5>
                  <table>
                    <thead>
                      <tr>
                        <th>Name</th>
                        <th>Description</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr>
                        <td>
                          <span class="property-name"><code>username</code></span> - <span class="property-type"><a href="#definition-String"><code>String!</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>email</code></span> - <span class="property-type"><a href="#definition-String"><code>String!</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>password</code></span> - <span class="property-type"><a href="#definition-String"><code>String!</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
              <div class="doc-examples">
                <h4 class="example-heading">Example</h4>
                <div class="example-section example-section-is-code operation-query-example">
                  <h5>Query</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-gql"><span class="hljs-symbol"><span class="hljs-keyword">mutation</span> Register<span class="hljs-tag">(
  <span class="hljs-code">$username</span>:<span class="hljs-type"> String!</span>,
  <span class="hljs-code">$email</span>:<span class="hljs-type"> String!</span>,
  <span class="hljs-code">$password</span>:<span class="hljs-type"> String!</span>
)</span> <span class="hljs-tag">{
  <span class="hljs-symbol">register<span class="hljs-tag">(
    username: <span class="hljs-code">$username</span>,
    email: <span class="hljs-code">$email</span>,
    password: <span class="hljs-code">$password</span>
  )</span> <span class="hljs-tag">{
    <span class="hljs-symbol">avatar</span>
    <span class="hljs-symbol">bio</span>
    <span class="hljs-symbol">email</span>
    <span class="hljs-symbol">firstName</span>
    <span class="hljs-symbol">id</span>
    <span class="hljs-symbol">isActive</span>
    <span class="hljs-symbol">isAdmin</span>
    <span class="hljs-symbol">lastName</span>
    <span class="hljs-symbol">location</span>
    <span class="hljs-symbol">username</span>
  }</span></span>
}</span></span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-variables-example">
                  <h5>Variables</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span>
  <span class="hljs-attr">"username"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"email"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"password"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span>
<span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-response-example">
                  <h5>Response</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span>
  <span class="hljs-attr">"data"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">{</span>
    <span class="hljs-attr">"register"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">{</span>
      <span class="hljs-attr">"avatar"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"bio"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"email"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"firstName"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"id"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"4"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"isActive"</span><span class="hljs-punctuation">:</span> <span class="hljs-literal"><span class="hljs-keyword">false</span></span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"isAdmin"</span><span class="hljs-punctuation">:</span> <span class="hljs-literal"><span class="hljs-keyword">true</span></span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"lastName"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"location"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"username"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span>
    <span class="hljs-punctuation">}</span>
  <span class="hljs-punctuation">}</span>
<span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
              </div>
            </div>
          </section>
          <section id="mutation-updatePet" class="operation operation-mutation" data-traverse-target="mutation-updatePet">
            <div class="operation-group-name">
              <a href="#group-Operations-Mutations">Mutations</a>
            </div>
            <h2 class="operation-heading ">
              <code>updatePet</code>
            </h2>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-description doc-copy-section">
                  <h5>Description</h5>
                  <p>Update pet information</p>
                </div>
              </div>
            </div>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-response doc-copy-section">
                  <h5>Response</h5>
                  <p> Returns a <a href="#definition-Pet"><code>Pet</code></a>
                  </p>
                </div>
                <div class="operation-arguments doc-copy-section">
                  <h5>Arguments</h5>
                  <table>
                    <thead>
                      <tr>
                        <th>Name</th>
                        <th>Description</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr>
                        <td>
                          <span class="property-name"><code>name</code></span> - <span class="property-type"><a href="#definition-String"><code>String!</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>type</code></span> - <span class="property-type"><a href="#definition-PetType"><code>PetType!</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>breed</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>gender</code></span> - <span class="property-type"><a href="#definition-Gender"><code>Gender</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>id</code></span> - <span class="property-type"><a href="#definition-ID"><code>ID!</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
              <div class="doc-examples">
                <h4 class="example-heading">Example</h4>
                <div class="example-section example-section-is-code operation-query-example">
                  <h5>Query</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-gql"><span class="hljs-symbol"><span class="hljs-keyword">mutation</span> UpdatePet<span class="hljs-tag">(
  <span class="hljs-code">$name</span>:<span class="hljs-type"> String!</span>,
  <span class="hljs-code">$type</span>:<span class="hljs-type"> PetType!</span>,
  <span class="hljs-code">$breed</span>:<span class="hljs-type"> String,</span>
  <span class="hljs-code">$gender</span>:<span class="hljs-type"> Gender,</span>
  <span class="hljs-code">$id</span>: ID!
)</span> <span class="hljs-tag">{
  <span class="hljs-symbol">updatePet<span class="hljs-tag">(
    name: <span class="hljs-code">$name</span>,
    type: <span class="hljs-code">$type</span>,
    breed: <span class="hljs-code">$breed</span>,
    gender: <span class="hljs-code">$gender</span>,
    id: <span class="hljs-code">$id</span>
  )</span> <span class="hljs-tag">{
    <span class="hljs-symbol">adoptionStatus</span>
    <span class="hljs-symbol">breed</span>
    <span class="hljs-symbol">currentOwnerID</span>
    <span class="hljs-symbol">gender</span>
    <span class="hljs-symbol">id</span>
    <span class="hljs-symbol">name</span>
    <span class="hljs-symbol">photo</span>
    <span class="hljs-symbol">shelterID</span>
    <span class="hljs-symbol">type</span>
  }</span></span>
}</span></span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-variables-example">
                  <h5>Variables</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span>
  <span class="hljs-attr">"name"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"type"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Cat"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"breed"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"gender"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Male"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"id"</span><span class="hljs-punctuation">:</span> <span class="hljs-number">4</span>
<span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-response-example">
                  <h5>Response</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span>
  <span class="hljs-attr">"data"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">{</span>
    <span class="hljs-attr">"updatePet"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">{</span>
      <span class="hljs-attr">"adoptionStatus"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Available"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"breed"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"currentOwnerID"</span><span class="hljs-punctuation">:</span> <span class="hljs-number">4</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"gender"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Male"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"id"</span><span class="hljs-punctuation">:</span> <span class="hljs-number">4</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"name"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"photo"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"shelterID"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"4"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"type"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Cat"</span>
    <span class="hljs-punctuation">}</span>
  <span class="hljs-punctuation">}</span>
<span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
              </div>
            </div>
          </section>
          <section id="mutation-updateProfile" class="operation operation-mutation" data-traverse-target="mutation-updateProfile">
            <div class="operation-group-name">
              <a href="#group-Operations-Mutations">Mutations</a>
            </div>
            <h2 class="operation-heading ">
              <code>updateProfile</code>
            </h2>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-description doc-copy-section">
                  <h5>Description</h5>
                  <p>Update user profile</p>
                </div>
              </div>
            </div>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-response doc-copy-section">
                  <h5>Response</h5>
                  <p> Returns a <a href="#definition-User"><code>User</code></a>
                  </p>
                </div>
                <div class="operation-arguments doc-copy-section">
                  <h5>Arguments</h5>
                  <table>
                    <thead>
                      <tr>
                        <th>Name</th>
                        <th>Description</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr>
                        <td>
                          <span class="property-name"><code>email</code></span> - <span class="property-type"><a href="#definition-String"><code>String!</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>firstName</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>lastName</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>bio</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>location</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>avatar</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>id</code></span> - <span class="property-type"><a href="#definition-ID"><code>ID!</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>username</code></span> - <span class="property-type"><a href="#definition-String"><code>String!</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
              <div class="doc-examples">
                <h4 class="example-heading">Example</h4>
                <div class="example-section example-section-is-code operation-query-example">
                  <h5>Query</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-gql"><span class="hljs-symbol"><span class="hljs-keyword">mutation</span> UpdateProfile<span class="hljs-tag">(
  <span class="hljs-code">$email</span>:<span class="hljs-type"> String!</span>,
  <span class="hljs-code">$firstName</span>:<span class="hljs-type"> String,</span>
  <span class="hljs-code">$lastName</span>:<span class="hljs-type"> String,</span>
  <span class="hljs-code">$bio</span>:<span class="hljs-type"> String,</span>
  <span class="hljs-code">$location</span>:<span class="hljs-type"> String,</span>
  <span class="hljs-code">$avatar</span>:<span class="hljs-type"> String,</span>
  <span class="hljs-code">$id</span>: ID!,
  <span class="hljs-code">$username</span>:<span class="hljs-type"> String!</span>
)</span> <span class="hljs-tag">{
  <span class="hljs-symbol">updateProfile<span class="hljs-tag">(
    email: <span class="hljs-code">$email</span>,
    firstName: <span class="hljs-code">$firstName</span>,
    lastName: <span class="hljs-code">$lastName</span>,
    bio: <span class="hljs-code">$bio</span>,
    location: <span class="hljs-code">$location</span>,
    avatar: <span class="hljs-code">$avatar</span>,
    id: <span class="hljs-code">$id</span>,
    username: <span class="hljs-code">$username</span>
  )</span> <span class="hljs-tag">{
    <span class="hljs-symbol">avatar</span>
    <span class="hljs-symbol">bio</span>
    <span class="hljs-symbol">email</span>
    <span class="hljs-symbol">firstName</span>
    <span class="hljs-symbol">id</span>
    <span class="hljs-symbol">isActive</span>
    <span class="hljs-symbol">isAdmin</span>
    <span class="hljs-symbol">lastName</span>
    <span class="hljs-symbol">location</span>
    <span class="hljs-symbol">username</span>
  }</span></span>
}</span></span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-variables-example">
                  <h5>Variables</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span>
  <span class="hljs-attr">"email"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"firstName"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"lastName"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"bio"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"location"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"avatar"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"id"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"4"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"username"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span>
<span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-response-example">
                  <h5>Response</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span>
  <span class="hljs-attr">"data"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">{</span>
    <span class="hljs-attr">"updateProfile"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">{</span>
      <span class="hljs-attr">"avatar"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"bio"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"email"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"firstName"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"id"</span><span class="hljs-punctuation">:</span> <span class="hljs-number">4</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"isActive"</span><span class="hljs-punctuation">:</span> <span class="hljs-literal"><span class="hljs-keyword">true</span></span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"isAdmin"</span><span class="hljs-punctuation">:</span> <span class="hljs-literal"><span class="hljs-keyword">false</span></span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"lastName"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"location"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"username"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span>
    <span class="hljs-punctuation">}</span>
  <span class="hljs-punctuation">}</span>
<span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
              </div>
            </div>
          </section>
          <section id="mutation-updateShelter" class="operation operation-mutation" data-traverse-target="mutation-updateShelter">
            <div class="operation-group-name">
              <a href="#group-Operations-Mutations">Mutations</a>
            </div>
            <h2 class="operation-heading ">
              <code>updateShelter</code>
            </h2>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-description doc-copy-section">
                  <h5>Description</h5>
                  <p>Update shelter information</p>
                </div>
              </div>
            </div>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="operation-response doc-copy-section">
                  <h5>Response</h5>
                  <p> Returns a <a href="#definition-Shelter"><code>Shelter</code></a>
                  </p>
                </div>
                <div class="operation-arguments doc-copy-section">
                  <h5>Arguments</h5>
                  <table>
                    <thead>
                      <tr>
                        <th>Name</th>
                        <th>Description</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr>
                        <td>
                          <span class="property-name"><code>location</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>contactInformation</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>id</code></span> - <span class="property-type"><a href="#definition-ID"><code>ID!</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>name</code></span> - <span class="property-type"><a href="#definition-String"><code>String!</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>description</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <span class="property-name"><code>website</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
              <div class="doc-examples">
                <h4 class="example-heading">Example</h4>
                <div class="example-section example-section-is-code operation-query-example">
                  <h5>Query</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-gql"><span class="hljs-symbol"><span class="hljs-keyword">mutation</span> UpdateShelter<span class="hljs-tag">(
  <span class="hljs-code">$location</span>:<span class="hljs-type"> String,</span>
  <span class="hljs-code">$contactInformation</span>:<span class="hljs-type"> String,</span>
  <span class="hljs-code">$id</span>: ID!,
  <span class="hljs-code">$name</span>:<span class="hljs-type"> String!</span>,
  <span class="hljs-code">$description</span>:<span class="hljs-type"> String,</span>
  <span class="hljs-code">$website</span>:<span class="hljs-type"> String
</span>)</span> <span class="hljs-tag">{
  <span class="hljs-symbol">updateShelter<span class="hljs-tag">(
    location: <span class="hljs-code">$location</span>,
    contactInformation: <span class="hljs-code">$contactInformation</span>,
    id: <span class="hljs-code">$id</span>,
    name: <span class="hljs-code">$name</span>,
    description: <span class="hljs-code">$description</span>,
    website: <span class="hljs-code">$website</span>
  )</span> <span class="hljs-tag">{
    <span class="hljs-symbol">contactInformation</span>
    <span class="hljs-symbol">description</span>
    <span class="hljs-symbol">id</span>
    <span class="hljs-symbol">location</span>
    <span class="hljs-symbol">logo</span>
    <span class="hljs-symbol">name</span>
    <span class="hljs-symbol">numberOfPets</span>
    <span class="hljs-symbol">ownerID</span>
    <span class="hljs-symbol">website</span>
  }</span></span>
}</span></span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-variables-example">
                  <h5>Variables</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span>
  <span class="hljs-attr">"location"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"contactInformation"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"id"</span><span class="hljs-punctuation">:</span> <span class="hljs-number">4</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"name"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"description"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"website"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span>
<span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
                <div class="example-section example-section-is-code operation-response-example">
                  <h5>Response</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span>
  <span class="hljs-attr">"data"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">{</span>
    <span class="hljs-attr">"updateShelter"</span><span class="hljs-punctuation">:</span> <span class="hljs-punctuation">{</span>
      <span class="hljs-attr">"contactInformation"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"description"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"id"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"4"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"location"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"logo"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"name"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"numberOfPets"</span><span class="hljs-punctuation">:</span> <span class="hljs-number">123</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"ownerID"</span><span class="hljs-punctuation">:</span> <span class="hljs-number">4</span><span class="hljs-punctuation">,</span>
      <span class="hljs-attr">"website"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span>
    <span class="hljs-punctuation">}</span>
  <span class="hljs-punctuation">}</span>
<span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
              </div>
            </div>
          </section>
          <h1 id="group-Types" class="group-heading" data-traverse-target="group-Types">Types</h1>
          <section id="definition-AdoptionStatus" class="definition definition-enum" data-traverse-target="definition-AdoptionStatus">
            <h2 class="definition-heading">AdoptionStatus</h2>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="definition-properties doc-copy-section">
                  <h5>Values</h5>
                  <table>
                    <thead>
                      <tr>
                        <th>Enum Value</th>
                        <th>Description</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr>
                        <td>
                          <p><code>Available</code></p>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <p><code>Adopted</code></p>
                        </td>
                        <td>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
              <div class="doc-examples">
                <div class="example-section example-section-is-code">
                  <h5>Example</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-gql"><span class="hljs-symbol">"Available"</span>
</code></pre>
                    </body>
                  </html>
                </div>
              </div>
            </div>
          </section>
          <section id="definition-Boolean" class="definition definition-scalar" data-traverse-target="definition-Boolean">
            <div class="definition-group-name">
              <a href="#group-Types">Types</a>
            </div>
            <h2 class="definition-heading">Boolean</h2>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="definition-description doc-copy-section">
                  <h5>Description</h5>
                  <p>The <code>Boolean</code> scalar type represents <code>true</code> or <code>false</code>.</p>
                </div>
              </div>
              <div class="doc-examples">
                <div class="example-section example-section-is-code">
                  <h5>Example</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-literal"><span class="hljs-keyword">true</span></span>
</code></pre>
                    </body>
                  </html>
                </div>
              </div>
            </div>
          </section>
          <section id="definition-Gender" class="definition definition-enum" data-traverse-target="definition-Gender">
            <div class="definition-group-name">
              <a href="#group-Types">Types</a>
            </div>
            <h2 class="definition-heading">Gender</h2>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="definition-properties doc-copy-section">
                  <h5>Values</h5>
                  <table>
                    <thead>
                      <tr>
                        <th>Enum Value</th>
                        <th>Description</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr>
                        <td>
                          <p><code>Male</code></p>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <p><code>Female</code></p>
                        </td>
                        <td>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
              <div class="doc-examples">
                <div class="example-section example-section-is-code">
                  <h5>Example</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-gql"><span class="hljs-symbol">"Male"</span>
</code></pre>
                    </body>
                  </html>
                </div>
              </div>
            </div>
          </section>
          <section id="definition-ID" class="definition definition-scalar" data-traverse-target="definition-ID">
            <div class="definition-group-name">
              <a href="#group-Types">Types</a>
            </div>
            <h2 class="definition-heading">ID</h2>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="definition-description doc-copy-section">
                  <h5>Description</h5>
                  <p>The <code>ID</code> scalar type represents a unique identifier, often used to refetch an object or as key for a cache. The ID type appears in a JSON response as a String; however, it is not intended to be human-readable. When expected as an input type, any string (such as <code>&quot;4&quot;</code>) or integer (such as <code>4</code>) input value will be accepted as an ID.</p>
                </div>
              </div>
              <div class="doc-examples">
                <div class="example-section example-section-is-code">
                  <h5>Example</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-gql"><span class="hljs-symbol">"4"</span>
</code></pre>
                    </body>
                  </html>
                </div>
              </div>
            </div>
          </section>
          <section id="definition-Int" class="definition definition-scalar" data-traverse-target="definition-Int">
            <div class="definition-group-name">
              <a href="#group-Types">Types</a>
            </div>
            <h2 class="definition-heading">Int</h2>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="definition-description doc-copy-section">
                  <h5>Description</h5>
                  <p>The <code>Int</code> scalar type represents non-fractional signed whole numeric values. Int can represent values between -(2^31) and 2^31 - 1.</p>
                </div>
              </div>
              <div class="doc-examples">
                <div class="example-section example-section-is-code">
                  <h5>Example</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-number">123</span>
</code></pre>
                    </body>
                  </html>
                </div>
              </div>
            </div>
          </section>
          <section id="definition-Pet" class="definition definition-object" data-traverse-target="definition-Pet">
            <div class="definition-group-name">
              <a href="#group-Types">Types</a>
            </div>
            <h2 class="definition-heading">Pet</h2>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="definition-properties doc-copy-section">
                  <h5>Fields</h5>
                  <table>
                    <thead>
                      <tr>
                        <th>Field Name</th>
                        <th>Description</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr>
                        <td data-property-name=""><span class="property-name"><code>adoptionStatus</code></span> - <span class="property-type"><a href="#definition-AdoptionStatus"><code>AdoptionStatus</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td data-property-name=""><span class="property-name"><code>breed</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td data-property-name=""><span class="property-name"><code>currentOwnerID</code></span> - <span class="property-type"><a href="#definition-ID"><code>ID</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td data-property-name=""><span class="property-name"><code>gender</code></span> - <span class="property-type"><a href="#definition-Gender"><code>Gender</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td data-property-name=""><span class="property-name"><code>id</code></span> - <span class="property-type"><a href="#definition-ID"><code>ID</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td data-property-name=""><span class="property-name"><code>name</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td data-property-name=""><span class="property-name"><code>photo</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td data-property-name=""><span class="property-name"><code>shelterID</code></span> - <span class="property-type"><a href="#definition-ID"><code>ID</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td data-property-name=""><span class="property-name"><code>type</code></span> - <span class="property-type"><a href="#definition-PetType"><code>PetType</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
              <div class="doc-examples">
                <div class="example-section example-section-is-code">
                  <h5>Example</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span>
  <span class="hljs-attr">"adoptionStatus"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Available"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"breed"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"currentOwnerID"</span><span class="hljs-punctuation">:</span> <span class="hljs-number">4</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"gender"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Male"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"id"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"4"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"name"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"photo"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"shelterID"</span><span class="hljs-punctuation">:</span> <span class="hljs-number">4</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"type"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"Cat"</span>
<span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
              </div>
            </div>
          </section>
          <section id="definition-PetType" class="definition definition-enum" data-traverse-target="definition-PetType">
            <div class="definition-group-name">
              <a href="#group-Types">Types</a>
            </div>
            <h2 class="definition-heading">PetType</h2>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="definition-properties doc-copy-section">
                  <h5>Values</h5>
                  <table>
                    <thead>
                      <tr>
                        <th>Enum Value</th>
                        <th>Description</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr>
                        <td>
                          <p><code>Cat</code></p>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td>
                          <p><code>Dog</code></p>
                        </td>
                        <td>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
              <div class="doc-examples">
                <div class="example-section example-section-is-code">
                  <h5>Example</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-gql"><span class="hljs-symbol">"Cat"</span>
</code></pre>
                    </body>
                  </html>
                </div>
              </div>
            </div>
          </section>
          <section id="definition-Shelter" class="definition definition-object" data-traverse-target="definition-Shelter">
            <div class="definition-group-name">
              <a href="#group-Types">Types</a>
            </div>
            <h2 class="definition-heading">Shelter</h2>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="definition-properties doc-copy-section">
                  <h5>Fields</h5>
                  <table>
                    <thead>
                      <tr>
                        <th>Field Name</th>
                        <th>Description</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr>
                        <td data-property-name=""><span class="property-name"><code>contactInformation</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td data-property-name=""><span class="property-name"><code>description</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td data-property-name=""><span class="property-name"><code>id</code></span> - <span class="property-type"><a href="#definition-ID"><code>ID</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td data-property-name=""><span class="property-name"><code>location</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td data-property-name=""><span class="property-name"><code>logo</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td data-property-name=""><span class="property-name"><code>name</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td data-property-name=""><span class="property-name"><code>numberOfPets</code></span> - <span class="property-type"><a href="#definition-Int"><code>Int</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td data-property-name=""><span class="property-name"><code>ownerID</code></span> - <span class="property-type"><a href="#definition-ID"><code>ID</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td data-property-name=""><span class="property-name"><code>website</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
              <div class="doc-examples">
                <div class="example-section example-section-is-code">
                  <h5>Example</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span>
  <span class="hljs-attr">"contactInformation"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"description"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"id"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"4"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"location"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"logo"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"name"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"numberOfPets"</span><span class="hljs-punctuation">:</span> <span class="hljs-number">987</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"ownerID"</span><span class="hljs-punctuation">:</span> <span class="hljs-number">4</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"website"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span>
<span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
              </div>
            </div>
          </section>
          <section id="definition-String" class="definition definition-scalar" data-traverse-target="definition-String">
            <div class="definition-group-name">
              <a href="#group-Types">Types</a>
            </div>
            <h2 class="definition-heading">String</h2>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="definition-description doc-copy-section">
                  <h5>Description</h5>
                  <p>The <code>String</code> scalar type represents textual data, represented as UTF-8 character sequences. The String type is most often used by GraphQL to represent free-form human-readable text.</p>
                </div>
              </div>
              <div class="doc-examples">
                <div class="example-section example-section-is-code">
                  <h5>Example</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-gql"><span class="hljs-symbol">"xyz789"</span>
</code></pre>
                    </body>
                  </html>
                </div>
              </div>
            </div>
          </section>
          <section id="definition-User" class="definition definition-object" data-traverse-target="definition-User">
            <div class="definition-group-name">
              <a href="#group-Types">Types</a>
            </div>
            <h2 class="definition-heading">User</h2>
            <div class="doc-row">
              <div class="doc-copy">
                <div class="definition-properties doc-copy-section">
                  <h5>Fields</h5>
                  <table>
                    <thead>
                      <tr>
                        <th>Field Name</th>
                        <th>Description</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr>
                        <td data-property-name=""><span class="property-name"><code>avatar</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td data-property-name=""><span class="property-name"><code>bio</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td data-property-name=""><span class="property-name"><code>email</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td data-property-name=""><span class="property-name"><code>firstName</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td data-property-name=""><span class="property-name"><code>id</code></span> - <span class="property-type"><a href="#definition-ID"><code>ID</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td data-property-name=""><span class="property-name"><code>isActive</code></span> - <span class="property-type"><a href="#definition-Boolean"><code>Boolean</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td data-property-name=""><span class="property-name"><code>isAdmin</code></span> - <span class="property-type"><a href="#definition-Boolean"><code>Boolean</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td data-property-name=""><span class="property-name"><code>lastName</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td data-property-name=""><span class="property-name"><code>location</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                      <tr>
                        <td data-property-name=""><span class="property-name"><code>username</code></span> - <span class="property-type"><a href="#definition-String"><code>String</code></a></span>
                        </td>
                        <td>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
              <div class="doc-examples">
                <div class="example-section example-section-is-code">
                  <h5>Example</h5>
                  <html>
                    <head></head>
                    <body><pre><code class="hljs language-json"><span class="hljs-punctuation">{</span>
  <span class="hljs-attr">"avatar"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"bio"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"email"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"firstName"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"id"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"4"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"isActive"</span><span class="hljs-punctuation">:</span> <span class="hljs-literal"><span class="hljs-keyword">true</span></span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"isAdmin"</span><span class="hljs-punctuation">:</span> <span class="hljs-literal"><span class="hljs-keyword">true</span></span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"lastName"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"location"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"abc123"</span><span class="hljs-punctuation">,</span>
  <span class="hljs-attr">"username"</span><span class="hljs-punctuation">:</span> <span class="hljs-string">"xyz789"</span>
<span class="hljs-punctuation">}</span>
</code></pre>
                    </body>
                  </html>
                </div>
              </div>
            </div>
          </section>
          <div class="doc-row no-margin">
            <div class="doc-copy doc-separator">
              <a class="powered-by" href="https://github.com/anvilco/spectaql">Documentation by <span>Anvil SpectaQL</span></a>
            </div>
          </div>
        </article>
        <div class="drawer-overlay"></div>
      </div>
    </div>
  </body>
</html>


- Overview of the API
- Description of the different API endpoints
- Parameters, query strings, and headers required for each endpoint
- Sample requests and responses for each endpoint

## Codebase
- Overview of the codebase
- Description of the different modules and packages used in the project
- Explanation of the different functions and their roles in the system
- Explanation of any notable design patterns or frameworks used in the project

## Deployment
- Explanation of how to deploy the project to a production environment
- Description of the hosting and infrastructure requirements

## Maintenance and Support
- Explanation of how to maintain and support the project in a production environment
- Information on how to debug and troubleshoot common issues
- Contact information for support and feedback

## Conclusion
- Summary of the project
- Future plans and enhancements

( Add support for adding must columns in sql db implementations.
It's needed if we want to update column with zero value. )

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
            - database.go
    - logr/
        - logger.go
- models/
    - gqtypes/
        - pet.go
        - shelter.go
        - types.go
        - user.go
    - errors.go
    - pet.go
    - shelter.go
    - user.go
- pkg/
    - decode.go
    - hash.go
    - path.go
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
    - pet.go
    - shelter.go
    - user.go
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
