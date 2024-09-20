package constant

type StatusCode struct {
	Status  int
	Code    string
	Message string
}

var Response = []StatusCode{
	{
		Status:  200,
		Code:    "SUCCESS",
		Message: "Request processed successfully",
	},
	{
		Status:  401,
		Code:    "UNAUTHORIZED",
		Message: "Unauthorized access. Please provide valid credentials.",
	},
	{
		Status:  401,
		Code:    "INVALID_JWT_TOKEN",
		Message: "Missing or invalid Authorization header",
	},
	{
		Status:  401,
		Code:    "EXPIRED_JWT_TOKEN",
		Message: "Missing or invalid Authorization header",
	},
	{
		Status:  420,
		Code:    "INVALID_PAYLOAD",
		Message: "Missing or invalid Body Payload",
	},
	{
		Status:  500,
		Code:    "GENERAL_ERROR",
		Message: "Internal Server Error",
	},
}
