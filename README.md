# Todo-Backend: Go kit

![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/sagikazarmark/todobackend-go-kit/ci.yaml?style=flat-square)
[![OpenSSF Scorecard](https://api.securityscorecards.dev/projects/github.com/sagikazarmark/todobackend-go-kit/badge)](https://api.securityscorecards.dev/projects/github.com/sagikazarmark/todobackend-go-kit)
[![Go Report Card](https://goreportcard.com/badge/github.com/sagikazarmark/todobackend-go-kit?style=flat-square)](https://goreportcard.com/report/github.com/sagikazarmark/todobackend-go-kit)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/mod/github.com/sagikazarmark/todobackend-go-kit)

A simple [Todo-Backend](http://todobackend.com/) application written using [Go kit](https://gokit.io/).

*Code extracted from [Modern Go Application](https://github.com/sagikazarmark/modern-go-application/tree/558c8cf1844fd76399f1e086b4df1385bf6ea439).*

[Try the client](http://todobackend.com/client/index.html?https://todobackend-go-kit.fly.dev/todos) or
[run the tests](http://todobackend.com/specs/index.html?https://todobackend-go-kit.fly.dev/todos).

In addition to the [Todo-Backend](http://todobackend.com/) implementation, this project comes with:

- BDD style tests using [GoBDD](https://go-bdd.github.io/gobdd/) *(Specifications are created based on [this](https://paulhammant.com/2017/05/14/todomvc-and-given-when-then-scenarios/) article.)*
- [Modern Go Application](https://github.com/sagikazarmark/modern-go-application) practices
- Separate [API SDK](api/) module
- [gRPC](https://grpc.io/) and [GraphQL](https://graphql.org/) implementation


## Development

1. Clone the repository
1. `make run`

For the best developer experience, install [Nix](https://builtwithnix.org/) and [direnv](https://direnv.net/).

Alternatively, install Go manually or using a package manager. Install the rest of the dependencies by running `make deps`.

[Try the client](http://todobackend.com/client/index.html?http://localhost:8000/todos) or
[run the tests](http://todobackend.com/specs/index.html?http://localhost:8000/todos) locally.


## License

The project is licensed under the [MIT License](LICENSE).
