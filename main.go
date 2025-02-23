package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"unsafe"
)

type Message struct {
	Path string `json:"path"`
	ID   string `json:"id"`
}

type PathResponse struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

//export receive_string
func receive_string(ptr *byte, length int) {
	data := unsafe.Slice(ptr, length) // Convert to slice
	str := string(data)               // Convert to Go string

	if str == "" {
		fmt.Println("")
	}

	fmt.Println("Received from JS:", str)
}

//export convert_string
func convert_string(e *byte, l int)

func send_string(e string) {
	firstByte := &(([]byte)(e)[0])
	convert_string(firstByte, len(e))
}

//export handle_path_response
func handle_path_response(e *byte, l int)

//export send_path_message
func send_path_message(ptr *byte, length int) {
	data := unsafe.Slice(ptr, length) // Convert to slice
	str := string(data)               // Convert to Go string

	if str == "" {
		fmt.Println("")
	}

	message := Message{}
	err := json.Unmarshal(data, &message)
	if err != nil {
		fmt.Println("Error parsing: ", str)
		return
	}

	process_message(message)
}

func process_message(message Message) {
	fmt.Println("parsing: ", message.Path)
	var b bytes.Buffer
	Test().Render(context.Background(), &b)
	str := b.String()

	path_response := PathResponse{
		ID:   message.ID,
		Text: str,
	}

	path_response_bytes, err := json.Marshal(path_response)
	if err != nil {
		fmt.Println("Error stringifying: ", message.ID)
		return
	}

	path_response_string := string(path_response_bytes)

	firstByte := &(([]byte)(path_response_string)[0])
	handle_path_response(firstByte, len(path_response_string))
}

//export handle_send_paths
func handle_send_paths(e *byte, l int)
func send_paths() {
	paths := []string{"/wasm/button", "/wasm/option", "/wasm/input"}
	str := strings.Join(paths, ",")
	firstByte := &(([]byte)(str)[0])
	handle_send_paths(firstByte, len(str))
}

func main() {
	send_paths()
	// send_path_response()
	// paths := []string{"one", "two"}
	// fmt.Println("yo from wasm")
	// send_string(strings.Join(paths, ","))
	//
	// var buffer bytes.Buffer
	//
	// Test().Render(context.Background(), &buffer)
	// send_string(buffer.String())
}
