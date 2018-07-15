package engine

type Return struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

type Config struct {
	Root string `json:"root"`
}
