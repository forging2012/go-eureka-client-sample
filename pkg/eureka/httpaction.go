package eureka

type HttpAction struct {
	Method      string
	Url         string
	Body        string
	Template    string
	Accept      string
	ContentType string
	Headers     map[string]string
}
