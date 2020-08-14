package function

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

func ReadBinaryFile(file_name string) {
	file, err := os.Open(file_name)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	var m int8
	data := ReadNextBytes(file, 16)
	buffer := bytes.NewBuffer(data)
	err = binary.Read(buffer, binary.BigEndian, &m)
	if err != nil {
		log.Fatal("binary.Read failed", err)
	}
	fmt.Println(m)
}

func ReadNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)
	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

func WriteBinaryFile(file_name string, file_path string, file_size int64) {
	file_route := file_path + file_name
	file, err := os.Create(file_route)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
		fmt.Println("Cannot create the file")
	}
	var cero int64 = 0
	s := &cero
	var bin_buf bytes.Buffer
	binary.Write(&bin_buf, binary.BigEndian, s)
	WriteNextBytes(file, bin_buf.Bytes())
	file.Seek(file_size,0)
	var second_buffer bytes.Buffer
	binary.Write(&second_buffer, binary.BigEndian, s)
	WriteNextBytes(file, second_buffer.Bytes())
}

func WriteNextBytes(file *os.File, bytes []byte) {
	_, err := file.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
}