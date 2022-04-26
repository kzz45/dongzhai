package k8s

import (
	"crypto/tls"
	"dongzhai/db"
	"dongzhai/models"
	k8s_model "dongzhai/models/k8s"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"gorm.io/gorm"
)

func GetRegistries(p *models.Pagination) ([]k8s_model.Registry, int64, error) {
	var registries []k8s_model.Registry
	if p.Size < 1 {
		p.Size = 15
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if err := db.GlobalGorm.Find(&registries).Count(&p.Total).Error; err != nil {
		return nil, 0, err
	}
	offset := p.Size * (p.Page - 1)
	if err := db.GlobalGorm.Limit(p.Size).Offset(offset).Find(&registries).Error; err != nil {
		return nil, 0, err
	}
	return registries, p.Total, nil
}

func GetRegistryById(id int) (k8s_model.Registry, error) {
	var registry k8s_model.Registry
	err := db.GlobalGorm.Where("id = ?", id).First(&registry).Error
	return registry, err
}

func CreateRegistry(registry k8s_model.Registry) error {
	var r k8s_model.Registry
	if !errors.Is(db.GlobalGorm.Where("name = ?", registry.Name).First(&r).Error, gorm.ErrRecordNotFound) {
		return errors.New("registry exist")
	}
	return db.GlobalGorm.Create(&registry).Error
}

func UpdateRegistry(registry k8s_model.Registry) error {
	return db.GlobalGorm.Where("id = ?", registry.ID).First(&k8s_model.Registry{}).Updates(&registry).Error
}

func DeleteRegistryById(id int) error {
	var registry k8s_model.Registry
	return db.GlobalGorm.Where("id = ?", id).Delete(&registry).Error
}

func ImageSearch(registry k8s_model.Registry, word string) []string {
	switch registry.Sign {
	case "Harbor":
		return searchHarbor(registry, word)
	case "aliyun":
		return searchAliYun()
	case "qcloud":
		return searchQCloud()
	case "DockerHub":
		return SearchDockerHub(word)
	}
	return nil
}

func GetImageTags(registry k8s_model.Registry, image_name string) []string {
	switch registry.Sign {
	case "Harbor":
		return getTagInHarbor(registry, image_name)
	case "aliyun":
		return searchAliYun()
	case "qcloud":
		return searchQCloud()
	}
	return nil
}

func searchHarbor(registry k8s_model.Registry, word string) []string {
	url := strings.TrimSuffix(registry.Addr, "/") + fmt.Sprintf("/api/search?q=%s", word)
	body, err := httpGet(url, registry.Username, registry.Password, false)
	if err != nil || len(body) == 0 {
		return nil
	}
	var repos k8s_model.HarborRepos
	repo_list := make([]string, 0, 100)
	err = json.Unmarshal(body, &repos)
	if err != nil {
		return nil
	}
	for _, repo := range repos.Repos {
		repo_list = append(repo_list, repo.RepoName)
	}

	return repo_list
}

func getTagInHarbor(registry k8s_model.Registry, image_name string) []string {
	url := strings.TrimSuffix(registry.Addr, "/") + fmt.Sprintf("/api/repositories/%s/tags", image_name)
	body, err := httpGet(url, registry.Username, registry.Password, false)
	if err != nil || len(body) == 0 {
		return nil
	}
	var tag_list []string
	err = json.Unmarshal(body, &tag_list)
	if err != nil {
		return nil
	}
	return tag_list
}

func searchAliYun() []string {
	return nil
}

func searchQCloud() []string {
	return nil
}

type dockerhubRepo struct {
	RepoName string `json:"repo_name"`
}
type dockerhubRepos struct {
	Repositories []dockerhubRepo `json:"results"`
}

func SearchDockerHub(word string) []string {
	url := fmt.Sprintf("https://hub.docker.com/v2/search/repositories/?page=1&query=%s&page_size=50", word)
	body, err := httpGet(url, "", "", true)
	if err != nil || len(body) == 0 {
		return nil
	}
	var repos dockerhubRepos
	err = json.Unmarshal(body, &repos)
	if err != nil {
		return nil
	}
	repo_list := make([]string, 0, 50)
	for _, repo := range repos.Repositories {
		repo_list = append(repo_list, repo.RepoName)
	}
	return repo_list
}

type dockerhubTag struct {
	TagName string `json:"name"`
}

type dockerhubTags struct {
	Tags []dockerhubTag `json:"results"`
}

func GetTagInDockerHub(image_name string) []string {
	if !strings.Contains(image_name, "/") {
		image_name = fmt.Sprintf("library/%s", image_name)
	}
	url := fmt.Sprintf("https://hub.docker.com/v2/repositories/%s/tags/?page=1&page_size=200", image_name)
	body, err := httpGet(url, "", "", true)
	if err != nil || len(body) == 0 {
		return nil
	}
	var tags dockerhubTags
	err = json.Unmarshal(body, &tags)
	if err != nil {
		return nil
	}

	tag_list := make([]string, 0, 200)
	for _, tag := range tags.Tags {
		tag_list = append(tag_list, tag.TagName)
	}

	return tag_list
}

func httpGet(url, username, password string, insecure bool) ([]byte, error) {
	var httpClient *http.Client
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if insecure {
		httpClient = &http.Client{}
	} else {
		request.SetBasicAuth(username, password)
		tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
		httpClient = &http.Client{Timeout: 20 * time.Second, Transport: tr}
	}
	resp, err := httpClient.Do(request)

	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if resp.StatusCode >= http.StatusBadRequest || err != nil {
		return nil, err
	}
	return body, nil
}
