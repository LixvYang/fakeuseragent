package parse

var (
	browserMap = make(map[string]string)
)

// Some short code for browsers
func init() {
	browserMap = map[string]string{
		"internet explorer": "internetexplorer",
		"ie":                "internetexplorer",
		"msie":              "internetexplorer",
		"edge":              "internetexplorer",
		"google":            "chrome",
		"googlechrome":      "chrome",
		"ff":                "firefox",
		"sf":                "safari",
		"op":                "opera",
	}
}

// Statistic return a string list
func Statistics() []string {
	browserList := make([]string, 1)
	for _, value := range browserMap {
		browserList = append(browserList, value)
	}
	return browserList
}