package system

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

func Png2Base64(filePath string) string {
	imgFile, err := os.Open(filePath) // a QR code image
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer imgFile.Close()
	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)
	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	imgBase64Str := base64.StdEncoding.EncodeToString(buf)
	return "data:image/png;base64," + imgBase64Str
}
