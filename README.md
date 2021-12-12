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
