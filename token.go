package main

type TokenFunction func(string) string

func InitTokenizer() map[string]TokenFunction {
	functionMap := make(map[string]TokenFunction, 10)
	functionMap["json"] = JSONToken
	return functionMap
}

func AnalyzeData(payload string, tokens map[string]TokenFunction) string {
	return tokens["json"](payload)
}

func JSONToken(payload string) string {
	return payload
}
