package proto

// DiagnosticFromError ...
func DiagnosticFromError(err error) *Diagnostic {
	return &Diagnostic{
		Severity: Diagnostic_ERROR,
		Summary:  err.Error(),
	}
}
