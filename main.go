package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	smbJobChannel := make(chan SMBQueryRequest)

	smb := NewSMBClient(os.Getenv("SMB_HOST"), os.Getenv("SMB_PORT"), os.Getenv("SMB_USER"), os.Getenv("SMB_PASS"), smbJobChannel)
	go smb.RunSession()

}
