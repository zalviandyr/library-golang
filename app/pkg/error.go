package pkg

type InternalServerError struct {
	Message string
}

type BadRequestError struct {
	Message string
}

type NotFoundError struct {
	Message string
}

type UnprocessableEntityError struct {
	Message string
	Field   map[string]any
}
