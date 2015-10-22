# Go Base API

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
