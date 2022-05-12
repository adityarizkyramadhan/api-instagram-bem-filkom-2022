package service

import (
	"api-instagram-bem-filkom-2022/entity"
	"api-instagram-bem-filkom-2022/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func GetResponseFromIG() error {
	fmt.Println("GetResponseFromIG")
	var data entity.ResponseKedua
	link := "https://instagram47.p.rapidapi.com/user_posts?username=bemfilkomub"
	// link := fmt.Sprintf("https://instagram47.p.rapidapi.com/user_posts?username=%s", username)
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return err
	}

	req.Header.Add("X-RapidAPI-Host", "instagram47.p.rapidapi.com'")
	req.Header.Add("X-RapidAPI-Key", "72952e64bdmshe2e87b888f28b6dp1c9e44jsn7954eb2f9159")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	fmt.Println(res.StatusCode)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}
	if data.Body == nil {
		os.Exit(1)
	}

	var responses []model.DataIG
	for _, response := range data.Body["items"].([]interface{}) {
		var temp model.DataIG
		temp.ThumbnailSrc = strings.Replace(response.(map[string]interface{})["image_versions2"].(map[string]interface{})["candidates"].([]interface{})[0].(map[string]interface{})["url"].(string), "\\", "", 1)
		temp.Caption = response.(map[string]interface{})["caption"].(map[string]interface{})["text"].(string)
		dumb := time.Unix(int64(response.(map[string]interface{})["taken_at"].(float64)), 0)
		temp.Tanggal = fmt.Sprintf("%d %s %d", dumb.Day(), dumb.Month().String(), dumb.Year())
		temp.Hari = dumb.Weekday().String()
		temp.LinkMedia = fmt.Sprintf("https://www.instagram.com/p/%s/", response.(map[string]interface{})["code"].(string))
		responses = append(responses, temp)
	}
	if err := addData(responses); err != nil {
		return err
	}
	return nil
}
