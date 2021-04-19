# go-enviro ![CI](https://github.com/mcneilcode/go-enviro/workflows/Builds/badge.svg) [![GoDoc](https://godoc.org/github.com/mcneilcode/go-enviro?status.svg)](https://godoc.org/github.com/mcneilcode/go-enviro) [![Go Report Card](https://goreportcard.com/badge/github.com/mcneilcode/go-enviro)](https://goreportcard.com/report/github.com/mcneilcode/go-enviro)

Environment helpers for go.

## Usage

```go
import "github.com/mcneilcode/go-enviro"
```

## Get

Returns the value for the environment variable if it exists or the provided default variable if it doesn't as a string.

```go
enviro.Get("MY_VAR", "some_default")
```

## GetInt

Returns the value for the environment variable as an integer if it exists or the provided default integer if it doesn't.

```go
enviro.GetInt("MY_INT_VAR", 2021)
```

## GetInt64

Returns the value for the environment variable as a 64-bit signed integer if it exists or the provided default int64 if it doesn't.

```go
enviro.GetInt64("MY_INT64_VAR", 2021)
```

## GetBool

Returns the value for the environment variable as a boolean value if it exists or the provided default boolean if it doesn't.

```go
enviro.GetBool("MY_BOOL_VAR", false)
```

## GetFloat32

Returns the value for the environment variable as a float32 value if it exists or the provided default float32 if it doesn't.

```go
enviro.GetFloat("MY_FLOAT32_VAR", 9.99)
```

## GetFloat64

Returns the value for the environment variable as a float64 value if it exists or the provided default float64 if it doesn't.

```go
enviro.GetFloat("MY_FLOAT64_VAR", 9.99)
```

## GetSlice

Returns the value for the environment variable as a string slice if it exists or the provided default string slice if it doesn't.

```go
enviro.GetSlice("MY_SLICE_VAR", ",", "default,none")
enviro.GetSlice("MY_SLICE_VAR", ":", "default:none")
```
