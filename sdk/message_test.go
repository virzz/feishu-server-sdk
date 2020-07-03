package sdk

import (
	"fmt"
	"testing"
)

// GetImage -
func TestGetImage(t *testing.T) {
	err := GetImage("img_6fb9ff0e-df73-409e-9bad-aea3c1d6985g")
	if err != nil {
		t.Error(err)
	}
}

func TestUploadMessageImage(t *testing.T) {
	key, err := UploadMessageImage("../tests/robot.png")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(key)
}

// SendImageMsg Send Image Msg
func TestSendImageMsg(t *testing.T) {
	err := SendImageMsg("oc_74fb9acaafbeb874fbd0d4b4c8bbeb59", "img_6fb9ff0e-df73-409e-9bad-aea3c1d6985g")
	if err != nil {
		t.Error(err)
	}
}
