# Go Base API

Carrots Base API scaffolding for Go.

## Getting Started

Clone this project into your [$GOPATH](https://golang.org/cmd/go/#hdr-GOPATH_environment_variable).

The project by default is set up to point at `github.com/carrot/go-base-api`, but you can change that to whatever you'd like.  Just be sure to update all references in the code to match your new location.

#### Dependencies + Building

We manage our dependencies in this project with [gom](https://github.com/mattn/gom).  So to start off, you're going to have to install that:

```
go get github.com/mattn/gom
```

After you've installed gom, you can run the following command to install the dependencies:

```
gom install
```

After you have the dependencies installed, you can build the project:

```
gom build
```

An executable file with the name of the root folder should now appear in the current directory.  You can run this and navigate to `http://localhost:5000/`.  If you receive a page that says `Not Found` you're all set up!

## Main

## Controllers

## Database

## Middleware

This contains a set of commonly used middleware created for use with the Echo framework.

- `Recover` - Recovers from `panic` calls. It's based off of the Echo-provided middleware of the same name but updated to fit our specific JSON interface model.

## Models

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

Note: `AddError` may be called more than once to indicate multiple errors as could happen with form validations.

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
