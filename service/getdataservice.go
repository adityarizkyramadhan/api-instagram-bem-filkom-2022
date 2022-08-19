package service

import (
	"api-instagram-bem-filkom-2022/config"
	"api-instagram-bem-filkom-2022/entity"
	"api-instagram-bem-filkom-2022/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"gorm.io/gorm"
)

func UpdateResponseFromIG(db *gorm.DB) ([]model.DataIG, error) {
	db_manual := config.GetDB()
	test, err := GetDataBemFilkom(db)
	if err != nil {
		return []model.DataIG{}, err
	}
	if test != nil {
		db_manual.ExecContext(context.Background(), "TRUNCATE TABLE data_igs;")
	}
	var data entity.ResponseKedua
	link := "https://instagram47.p.rapidapi.com/user_posts?username=bemfilkomub"

	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return []model.DataIG{}, err
	}

	req.Header.Add("X-RapidAPI-Host", "instagram47.p.rapidapi.com'")
	req.Header.Add("X-RapidAPI-Key", os.Getenv("RAPID_API_KEY"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []model.DataIG{}, err
	}
	// fmt.Println(res.StatusCode)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	if err != nil || body == nil {
		return []model.DataIG{}, err
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return []model.DataIG{}, err
	}
	var responses []model.DataIG
	for _, response := range data.Body["items"].([]interface{}) {
		if response == nil {
			break
		}
		var temp model.DataIG
		temp.ThumbnailSrc = strings.Replace(response.(map[string]interface{})["image_versions2"].(map[string]interface{})["candidates"].([]interface{})[0].(map[string]interface{})["url"].(string), "\\", "", 1)
		temp.Caption = response.(map[string]interface{})["caption"].(map[string]interface{})["text"].(string)
		dumb := time.Unix(int64(response.(map[string]interface{})["taken_at"].(float64)), 0)
		temp.Tanggal = fmt.Sprintf("%d %s %d", dumb.Day(), dumb.Month().String(), dumb.Year())
		temp.Hari = dumb.Weekday().String()
		temp.LinkMedia = fmt.Sprintf("https://www.instagram.com/p/%s/", response.(map[string]interface{})["code"].(string))
		responses = append(responses, temp)
	}
	// if responses != nil {
	// 	//delete all data from table
	// 	if err := db.Delete(&model.DataIG{}).Error; err != nil {
	// 		return []model.DataIG{}, err
	// 	}
	// }
	if err := db.Create(&responses).Error; err != nil {
		return []model.DataIG{}, err
	}
	return responses, nil
}

func UpdateResponseFromHastag(db *gorm.DB) ([]model.DataIGSjw, error) {
	db_manual := config.GetDB()
	test, err := GetDataBemFilkom(db)
	if err != nil {
		return []model.DataIGSjw{}, err
	}
	if test != nil {
		db_manual.ExecContext(context.Background(), "TRUNCATE TABLE data_ig_sjws;")
	}
	link := "https://instagram47.p.rapidapi.com/hashtag_post?hashtag=sjwfilkom"
	// link := fmt.Sprintf("https://instagram47.p.rapidapi.com/user_posts?username=%s", username)
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return []model.DataIGSjw{}, err
	}

	req.Header.Add("X-RapidAPI-Host", "instagram47.p.rapidapi.com'")
	req.Header.Add("X-RapidAPI-Key", "72952e64bdmshe2e87b888f28b6dp1c9e44jsn7954eb2f9159")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []model.DataIGSjw{}, err
	}
	defer res.Body.Close()
	// jsonFile, err := os.Open("../response.json")
	// if err != nil {
	// 	return err
	// }
	body, err := ioutil.ReadAll(res.Body)
	if err != nil || body == nil {
		return []model.DataIGSjw{}, err
	}
	fmt.Println(string(body))
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return []model.DataIGSjw{}, err
	}
	fmt.Println(data)
	if data["body"] == nil {
		return []model.DataIGSjw{}, errors.New("data is empty")
	}
	var responses []model.DataIGSjw
	for _, response := range data["body"].(map[string]interface{})["edge_hashtag_to_media"].(map[string]interface{})["edges"].([]interface{}) {
		var temp model.DataIGSjw
		temp.ThumbnailSrc = strings.Replace(response.(map[string]interface{})["node"].(map[string]interface{})["thumbnail_src"].(string), "\\", "", 1)
		temp.Caption = response.(map[string]interface{})["node"].(map[string]interface{})["edge_media_to_caption"].(map[string]interface{})["edges"].([]interface{})[0].(map[string]interface{})["node"].(map[string]interface{})["text"].(string)
		dumb := time.Unix(int64(data["body"].(map[string]interface{})["edge_hashtag_to_media"].(map[string]interface{})["edges"].([]interface{})[0].(map[string]interface{})["node"].(map[string]interface{})["taken_at_timestamp"].(float64)), 0)
		temp.Tanggal = fmt.Sprintf("%d %s %d", dumb.Day(), dumb.Month().String(), dumb.Year())
		temp.Hari = dumb.Weekday().String()
		temp.LinkMedia = fmt.Sprintf("https://www.instagram.com/p/%s/", response.(map[string]interface{})["node"].(map[string]interface{})["shortcode"].(string))
		responses = append(responses, temp)
	}
	// if responses != nil {
	// 	err := db.Delete(&model.DataIGSjw{}).Unscoped().Error
	// 	if err != nil {
	// 		return []model.DataIGSjw{}, err
	// 	}
	// }
	if err := db.Create(&responses).Error; err != nil {
		return []model.DataIGSjw{}, err
	}
	return responses, nil
}

func GetDataBemFilkom(db *gorm.DB) ([]model.DataIG, error) {
	var data []model.DataIG
	if err := db.Find(&data).Error; err != nil {
		return []model.DataIG{}, err
	}
	return data, nil
}

func GetDataSjwFilkom(db *gorm.DB) ([]model.DataIGSjw, error) {
	var data []model.DataIGSjw
	if err := db.Find(&data).Error; err != nil {
		return []model.DataIGSjw{}, err
	}
	return data, nil
}
