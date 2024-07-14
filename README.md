<div align="center">
    <img height="128" src="https://github.com/caesar-rocks/docs/raw/master/logo.svg" />
</div>

<div align="center">
    <h1>
        📜 Caesar Starter Kit
    </h1>
</div>

> _Caesar is a Go web framework, designed for productivity. It takes inspiration from traditional web frameworks such as Ruby on Rails, Laravel, Django, Phoenix, AdonisJS, etc._

## Setup

Once you have cloned this repository, you can run the following commands to get started:

```bash
task install
```

Run your migrations:

```bash
caesar migrations:migrate
```

and then start your development setup by launching these three commands (in separate terminals):

```bash
task air
```

```bash
task css
```

```bash
task templ
```


## Working with gRPC

We are using buf to help integrate gRPC into the framework.

Run the following to generate your proto file.

First make sure your buf.gen.yaml file has a the value set that it is your
go module's name.

Next make sure your proto file has the following added:

option go_package = '<modulename>/path/to/proto;<name>pb';

```bash
buf generate
```

Then check out the grpc/service.go file to see an exmaple of importing the handler for gRPC.

In config/routes.go can import your service.go file and then add the following within the RegisterRoutes function.

```go
// Get GRPC Service and add to the router
path, handler := v1connect.NewWeatherServiceHandler(&grpc.WeatherServiceServer{})
var extraHandler []caesar.ExtraHandler
extraHandler = append(extraHandler, caesar.ExtraHandler{
    Path:    path,
    Handler: handler,
})
router.CustomHandlers = extraHandler
```
