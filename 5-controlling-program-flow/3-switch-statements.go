package main

type HTTPRequest struct {
	Method string
}

// In go we have implicit breaking
// (instead of implicit fallthrough in switch statements)
func main() {
	r := HTTPRequest{Method: "GET"}

	switch r.Method {
	case "GET":
		println("Get request")
	case "POST":
		println("Post request")
	case "PUT":
		println("Put request")
	case "DELETE":
		println("Delete request")
	default:
		println("Unhandled method")
	}
}
