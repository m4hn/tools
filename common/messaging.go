package common

type Messenger interface {
	Send() ([]byte, error)
	SendFile() ([]byte, error)
	SendCustom(URL, message, title, content string) ([]byte, error)
	SendCustomFile(URL, message, fileName, title string, file []byte) ([]byte, error)
}
