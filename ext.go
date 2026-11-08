package htmx

const (
	HXDebug                  = "debug"                 // The debug extension allows you to debug htmx requests.
	HxExtClassTools          = "class-tools"           // The class-tools extension allows you to add and remove classes from elements.
	HxExtClientSideTemplates = "client-side-templates" // The client-side-templates extension allows you to use client-side templates to render the response.
	HxExtIgnoreDebug         = "ignore:debug"          // The ignore:debug extension allows you to ignore the debug header.
	HxExtJSON                = "json-enc"              // The json-enc extension allows you to specify the JSON encoding for the response.
	HxExtMultiSwap           = "multi-swap"            // The multi-swap extension allows you to swap multiple elements in a single response.
	HxExtPathDeps            = "path-deps"             // The path-deps extension allows you to specify the dependencies for a path.
	HxExtSSE                 = "sse"                   // The sse extension allows you to use Server-Sent Events (SSE) to stream updates to the client.
	HxResponseTargets        = "response-targets"      // The response-target extension allows you to specify the target for the response.
)

// MustacheTemplate sets the mustache-template attribute to specify the mustache template for the response.
func MustacheTemplate(v string) Node {
	return Attribute("mustache-template", v)
}

// HandlebarsTemplate sets the handlebars-template attribute to specify the handlebars template for the response.
func HandlebarsTemplate(v string) Node {
	return Attribute("handlebars-template", v)
}

// NunjucksTemplate sets the nunjucks-template attribute to specify the nunjucks template for the response.
func NunjucksTemplate(v string) Node {
	return Attribute("nunjucks-template", v)
}

// XSLTTemplat sets the xslt-template attribute to specify the xslt template for the response.
func XSLTTemplat(v string) Node {
	return Attribute("xslt-template", v)
}

// HyperScript sets the _ attribute to specify the hyperscript for the response.
func HyperScript(v string) Node {
	return Attribute("_", v)
}

// HxSSE sets the sse attribute to specify the Server-Sent Events (SSE) URL.
func HxSSE() Node {
	return HxExt(HxExtSSE)
}

// HxSSEConnect sets the sse-connect attribute to specify the Server-Sent Events (SSE) URL.
func HxSSEConnect(url string) Node {
	return Attribute("sse-connect", url)
}

// HxSSESwap sets the sse-swap attribute to specify the Server-Sent Events (SSE) swap.
func HxSSESwap(target string) Node {
	return Attribute("sse-swap", target)
}
