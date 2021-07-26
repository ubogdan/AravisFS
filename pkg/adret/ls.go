package adret

import (
	"encoding/base64"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/ariary/AravisFS/pkg/encrypt"
	"github.com/ariary/AravisFS/pkg/filesystem"
)

// Parse the directory content retrieve from encrypted fs (w/ ubac for example)
// First decrypt the whole content
// Then parse the content as each resources is separated by "\"
func ParseLsDirectoryContent(content string, key string) []string {
	contentDecoded, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		log.Fatal("error:", err)
	}
	lsResult := string(encrypt.DecryptByte(contentDecoded, key))
	//TO DO: add info if file is directory or not
	return strings.Split(lsResult, "\\")
}

// Parse the ls result received from ubac utility
// result structure: resourceType:base64(Encrypt(Content))
func ParseLsContent(result string, key string) string {
	//file or directory output?
	resultParsed := strings.SplitN(result, ":", 2)
	resourceType := resultParsed[0]
	content := resultParsed[1]
	if resourceType == filesystem.DIRECTORY {
		// decrypt content, parse it to have all resource under and print it
		files := ParseLsDirectoryContent(content, key)
		//take basename
		// TO DO propose to print the full path (ie don't call Base ~ ls -d $PWD/*)
		for i := range files {
			files[i] = filepath.Base(files[i])
		}

		//TO DO: specify if resource is a file or directory
		lsOutput := strings.Join(files, " ")
		return lsOutput

	} else if resourceType == filesystem.FILE {
		// decrypt content and print it (it must be the filename)
		lsOutput := string(encrypt.DecryptString(content, key))
		return lsOutput
	} else {
		log.Fatal("Failed to retrieve content after decrypting result")
	}
	return ""
}

func PrintLS(result string, key string) {
	output := ParseLsContent(result, key)
	if output != "" {
		fmt.Println(output)
	} else {
		log.Fatal("Failed to parse ls result")
	}

}