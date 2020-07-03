package sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/kirinlabs/HttpRequest"
	m "github.com/virzz/feishu-server-sdk/model"
)

// GetImage -
func GetImage(imageKey string) (err error) {
	res, err := HttpRequest.Debug(Debug).
		SetHeaders(map[string]string{"Authorization": GetAuthorization()}).
		Get(
			fmt.Sprintf("%s%s", FeishuAPI, APIGetImage),
			fmt.Sprintf("image_key=%s", imageKey),
		)
	if err != nil {
		return err
	}
	if res.StatusCode() != 200 {
		return errors.New("get image error")
	}
	return nil
}

func uploadImage(filename, imageType string) (imageKey string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", filename)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(part, file)
	writer.WriteField("image_type", imageType)
	err = writer.Close()
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", FeishuAPI, APIUploadImage), body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", GetAuthorization())
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	msg := &m.MsgResp{}
	if err = json.Unmarshal(respBytes, msg); err != nil {
		return "", err
	}
	if msg.Code > 0 {
		fmt.Println(string(respBytes))
	}
	return msg.Data.ImageKey, nil
}

// UploadMessageImage - 上传消息图片
func UploadMessageImage(filename string) (imageKey string, err error) {
	return uploadImage(filename, "message")
}

// UploadAvatarImage - 上传头像
func UploadAvatarImage(filename string) (imageKey string, err error) {
	return uploadImage(filename, "avatar")
}

// SendImageMsg Send Image Msg
func SendImageMsg(chatID, imageKey string) (err error) {
	data := &m.MsgReq{
		ChatID:  chatID,
		MsgType: "image",
		Content: &m.MsgContent{
			ImageKey: imageKey,
		},
	}
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}
	res, err := HttpRequest.Debug(Debug).
		SetHeaders(map[string]string{"Authorization": GetAuthorization()}).
		JSON().
		Post(
			fmt.Sprintf("%s%s", FeishuAPI, APIRobotSendMessage),
			string(dataJSON),
		)
	msgResp := &m.MsgResp{}
	if err = res.Json(msgResp); err != nil {
		return err
	}
	fmt.Println(msgResp)
	return nil
}
