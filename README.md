# go-enviro

Environment helper for go.

# Usage

```
import "github.com/mcneilcode/go-enviro"
```

## Get

Returns the value for the environment variable if it exists or the provided default variable if it doesn't.

```go
enviro.Get("MY_STR_VAR", "some_default")
```

## GetInt

Returns the value for the environment variable as an integer if it exists or the provided default integer if it doesn't.

```go
enviro.GetInt("MY_INT_VAR", 2020)
```

## GetBool

Returns the value for the environment variable as a boolean value if it exists or the provided default boolean if it doesn't.

```go
enviro.GetBool("MY_BOOL_VAR", false)
```

## GetSlice

Returns the value for the environment variable as a string slice if it exists or the provided default string slice if it doesn't.

```go
enviro.GetSlice("MY_SLICE_VAR", ",", "default,none")
enviro.GetSlice("MY_SLICE_VAR", ":", "default:none")
```
