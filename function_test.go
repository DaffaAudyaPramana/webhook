package webhook

import (
	"fmt"
	"os"
	"testing"

	"github.com/aiteung/atapi"
	"github.com/aiteung/atmessage"
	"github.com/whatsauth/wa"
)

func TestGetENVToken(t *testing.T) {
	fmt.Println(os.Getenv("TOKEN"))
}

func TestUpdateGetData(t *testing.T) {
	dt := &wa.TextMessage{
		To:       "6282117252716",
		IsGroup:  false,
		Messages: " gi apa  " + "daff",
	}
	//url := "https://api.wa.my.id/api/send/message/text"
	resp, error := atapi.PostStructWithToken[atmessage.Response]("Token", os.Getenv("TOKEN"), dt, "https://api.wa.my.id/api/send/message/text")
	fmt.Println(resp, error)
}
