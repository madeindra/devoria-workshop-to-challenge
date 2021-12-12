# Clean Architecture

## About
The project is about implementing [Uncle Bob's Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) in Golang.

![Uncle Bob's Clean Architecture](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)

### Project Structure
Main file is stored in **app**.

Features are stored in **domain** directory with following naming:

* Enterprise Business Rules: entitites
* Application Business Rules: usecase
* Interface Adapters: handler
* Framework & drivers: repositories (for database)

Other internally used files are stored in **internal**.

Postman Collections & other docs are stored in **docs**.

## Run
1.  Run dependencies
```
docker-compose -f docker/docker-compose.yml up -d
```

2. Setup `.env` by folliowing `.env.example`

3. Change the `private.key` and `public.key` to your own RSA key pair

4. Run
```
make run.dev
```

## Unit Test
Unit tests are created with `mockery` & `sqlmock`

To run it, run
```
make test
```

or to generate coverage report, run
```
make cover
```

## Accounts API
### Registration
* Endpoint: `/v1/accounts/registration`
* Method: `POST`
* Authorization: Basic
* Request Body
```
{
    "email": "example@example.com",
    "password": "test",
    "firstName": "test",
    "lastName": "test"
}
```

### Login
* Endpoint: `/v1/accounts/login`
* Method: `POST`
* Authorization: Basic
* Request Body
```
{
    "email": "example@example.com",
    "password": "test",
}
```
### Get Account
* Endpoint: `/v1/accounts/:id`
* Method: `GET`
* Authorization: Bearer
* Example
```
/v1/accounts/1
```

## Get All Articles
* Endpoint: `/v1/articles`
* Method: `GET`
* Authorization: Bearer

## Get an Article
* Endpoint: `/v1/articles/:id`
* Method: `GET`
* Authorization: Bearer
* Example
```
/v1/articles/1
```

## Create a new Article
* Endpoint: `/v1/articles`
* Method: `POST`
* Authorization: Bearer
* Request Body
```
{
    "title": "lorem",
    "subtitle": "ipsum",
    "content": "hello world",
    "isPublished": true
}
```

## Update an Article
* Endpoint: `/v1/articles`
* Method: `PATCH`
* Authorization: Bearer
* Example
```
/v1/articles/1
```
* Request Body
```
{
    "title": "lorem",
    "subtitle": "ipsum",
    "content": "hello world",
    "status": "ARCHIVED"
}
```
## Test Case
### Account Handler
* 2 functions tested
* 2 success test scenario
* 2 error test scenario

### Account Use Case
* 2 functions tested
* 2 success test scenarios
* 8 error test scenarios

### Account Repository
* 4 functions tested
* 4 success test scenario
* 4 error test scenario

## Credits
This repository is based on Devoria Workshop by [Sangian Patrick](https://github.com/sangianpatrick/devoria-article-service).
