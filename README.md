# Burrow

[![Godoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/carrot/burrow)  [![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](/LICENSE.md) [![Version](https://img.shields.io/github/release/carrot/burrow.svg?style=flat-square)](https://github.com/carrot/burrow/releases) [![Build Status](http://img.shields.io/travis/carrot/burrow.svg?style=flat-square)](https://travis-ci.org/carrot/burrow) [![Gitter](https://img.shields.io/badge/gitter-join%20chat-brightgreen.svg?style=flat-square)](https://gitter.im/carrot/burrow)

Burrow is a base API scaffolding for Go.  We use [echo](https://github.com/labstack/echo) as our base framework.

This is a Go implementation of Carrot's [RESTful API Spec](https://github.com/carrot/restful-api-spec).

## Getting Started

Clone this project into your [$GOPATH](https://golang.org/cmd/go/#hdr-GOPATH_environment_variable).

By default, this project is set up to point at `github.com/carrot/burrow`, but you can change that to whatever you'd like.  Just be sure to update all references in the code to match your new location.

#### Multiple Environment Support

This base supports multiple environments:

- development
- testing
- staging
- production

You will run this project with `./burrow {target-environment}`.  For example:

```
./burrow development
```

#### Environment Variables

Burrow uses [godotenv](https://github.com/joho/godotenv) (a Go port of [bkeepers/dotenv](https://github.com/bkeepers/dotenv))to manage environment variables.

Copy `.env.sample` to `.env.{target-environment}`, and update the values in the `.env.{target-environment}` file.

You'll also need to globally set this environment variable:

```sh
# Always set as 1, to manage dependencies
export GO15VENDOREXPERIMENT=1
```

#### Database

Burrow manages its database migrations with [Originator](https://github.com/DigitalCitadel/originator).  To start off, [get that installed](https://github.com/DigitalCitadel/originator#installation) on your machine.

After you have Originator installed, cd into `originator-files/config` and run `mkdir $(hostname)` to create a directory to hold your machine-specific config.

The only configuration file that must be added to your machine-specific config is the `database_config.bash` file.  Copy that from the default folder into your machine specific configuration folder.  Update your database config file to match your specific database setup.

After you've updated the file, navigate to the root of the project and run `originator migrate` to execute all of the current migrations. Your database should now be set up.

#### Dependencies + Building

Burrow manages its dependencies with [gom](https://github.com/mattn/gom).  To start off, you'll need to install that:

```sh
go get github.com/mattn/gom
```

After you've installed gom, you can run the following command to install the dependencies:

```sh
gom install
```

Once the dependencies are installed, you can build the project with the following command:

```sh
gom build
```

An executable file with the name of the root folder should now appear in the current directory.  Run the executable and navigate to `http://localhost:5000/`.  If you receive a page that says `Not Found` you're all set up!

#### Testing

To run all tests, run `gom test`.

Tests should be written to automatically load `.env.testing` (but not fail if it's not there) and as a user of the tests you should have your `.env.testing` file filled out.

If you're testing the project on Travis CI, set up the environment variables as you traditionally would with Travis and run normally.

## Controllers

Controllers are responsible for directly managing what happens during a request.  Every endpoint maps to a controller method.

To keep things clean, Burrow uses one controller per model (with the name {Model}Controller) and all handlers are methods.

We try to follow [CRUD](https://en.wikipedia.org/wiki/Create,_read,_update_and_delete) as the naming convention for all of our controller methods, with the exception of Read, which we use `Index` for bulk fetches, and `Show` for single fetches.

So, if we had a model named `people`, our methods would map a little something like this:

```
PeopleController.Create  ->  [POST]   /people
PeopleController.Show    ->  [GET]    /people/{id}
PeopleController.Index   ->  [GET]    /people
PeopleController.Update  ->  [PUT]    /people/{id}
PeopleController.Delete  ->  [DELETE] /people/{id}
```

#### Nested Controllers

In the event that you have nested endpoints, that look like this:

```
[GET] /people/{id}/pets
```

You're going to want to create a new controller to handle these relations.

Following our example above, we would create a `PeoplePetsController`.

## Routes

After you've set up a controller with some handlers, you now likely want to hook up the controller so we can actually call the handler from an HTTP request.

The routes are managed in the [request.go](/request/request.go) file.  In the `BuildEcho` function, you'll find a few sections.  One is for `Controllers`, which you should there create an instance of your controller.  There's another section for `Endpoints`, which is where you will register each controller method.

After both of these are done, you should now be able to send an HTTP request to run your controller method.

## Database

Burrow has configurations for both Redis and PostgreSQL.

### PostgreSQL

There are details of how to get PostgreSQL up and running in the [Getting Started](#getting-started) section of this README.

Burrow uses [lib/pq](https://github.com/lib/pq) as a database driver, but you really don't have to know that as it's already been abstracted away in the `db/postgres` package.  You will simply be interfacing with Go's [database/sql](https://golang.org/pkg/database/sql/).

## Middleware

This contains a set of commonly used middleware created for use with the Echo framework.

- `Recover` - Recovers from `panic` calls. It's based off of the Echo-provided middleware of the same name but updated to fit Burrows specific JSON interface model.

## Models

Models are responsible for storing, updating data, and exposing data to the application.  They are the interface between the database and the rest of the application.

We use one file per model, and store them in the `models/` directory and the name of the file is the `snake_case` version of the primary struct inside.

We use [methods](https://gobyexample.com/methods) for fetching/manipulating single model structs, and [functions](https://gobyexample.com/functions) for bulk fetching models.

We try to follow the naming conventions as described in this interface for all methods.  Burrow doesn't actually enforce the interface in code, as most applications don't require every one of these methods for all models.

```go
type Model interface {
    Load(id int64) error    // Loads the contents of model entry with ID into current struct
    Insert() error          // Inserts the state of the current struct into the DB
    Update() error          // Updates the state of the current struct to the DB
    Delete() error          // Removes the current struct from the DB
}
```

## Responses

The `response` package contains both consistent error code/messages as well as helpers to format JSON responses.

### Usage

#### Success

```go
func HomePage(c *echo.Context) error {
  resp := response.New(c)
  defer resp.Render()

  content := SimpleLogic()

  resp.SetResponse(http.StatusOK, content)
  return nil
}
```

#### Error

> Note: `AddError` may be called more than once to indicate multiple errors as could happen with form validations.

```go
func HomePage(c *echo.Context) error {
  resp := response.New(c)
  defer resp.Render()

  content, err := ComplexLogic()

  if err != nil {
    resp.AddError(response.ErrorInternalServerError)
    resp.SetResponse(http.StatusInternalServerError, nil)
    return nil
  }

  resp.SetResponse(http.StatusOK, content)
  return nil
}

```

#### Defaults

By using `defer` on `Render()` we can ensure that, even in the case of a `panic`, the response will still be rendered.
As a catch-all, the default response is set to the following:

```go
Response{
  Success: false,
  StatusCode: 500,
  StatusText: "Internal Server Error",
  Errors: [],
  Content: nil,
}
```

## Testing

Tests are run on Travis CI against Go versions:

- 1.5
- 1.5.1
- tip (failures allowed)

## License

[MIT](/LICENSE.md)
