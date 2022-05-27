package initconfig

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"readCommunity/internal/pkg/utils/cryptutil"
	"strings"
)

func SetConfig() {
	encrypt := cryptutil.Encrypt([]byte("876543210.adfjadf"))
	fmt.Println("encrypted data:>>>", hex.EncodeToString(encrypt))
	file,err := os.OpenFile("./configs/config.logconf.yaml", os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("open file config.logconf.yaml failed,err:%s", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	pos := int64(0)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("read file failed,err:%v", err)
				return
			}
		}
		if strings.Contains(line, "CryptKey") {
			bytes := []byte("CryptKey: " + hex.EncodeToString(encrypt) + "\n")
			file.WriteAt(bytes, pos)
		}
		pos += int64(len(line))
	}
}
