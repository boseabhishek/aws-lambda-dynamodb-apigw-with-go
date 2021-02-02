package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	ID    float64 `json:"id"`
	Value string  `json:"value"`
}

type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

// Lamdba handlers can take a variety of different signatures
// and reflection is used to determine exactly which signature
// you're using. 

// The full list of supported forms is…

// func()
// func() error
// func(TIn) error
// func() (TOut, error)
// func(TIn) (TOut, error)
// func(context.Context) error
// func(context.Context, TIn) error
// func(context.Context) (TOut, error)
// func(context.Context, TIn) (TOut, error)

func Handler(request Request) (Response, error) {
	return Response{
		Message: fmt.Sprintf("Process Request ID %f", request.ID),
		Ok:      true,
	}, nil
}

// he main() function we call lambda.Start() and
// pass in the show function as the lambda handler.
func main() {
	lambda.Start(Handler)
}
