package tencentcloud

func toHttpHeader(headers map[string]string) map[string][]string {
	header := make(map[string][]string, 0)
	for k, v := range headers {
		header[k] = []string{v}
	}
	return header
}
