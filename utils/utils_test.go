package utils

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func getConfig() Configurations {
	cfile := "../conf/config.yml"
	ctype := "yaml"
	var cfg Configurations
	LoadConfiguration(cfile, ctype, &cfg)
	return cfg
}

func TestLoadConfiguration(t *testing.T) {
	cfg := getConfig()
	if cfg.Server.PKey != "Ratings" {
		t.Errorf("LoadConfiguration() FAIL, Expected %v but found value %v ", "Ratings", cfg.Server.PKey)
	} else {
		t.Logf("LoadConfiguration() PASS, Expected %v and got value %v ", "Rating", cfg.Server.PKey)
	}
}

func TestGetURLString(t *testing.T) {
	cfg := getConfig()
	title := "frozen"
	url, ok := GetURLString(title, cfg)
	if ok && !strings.Contains(url, "92536907") {
		t.Errorf("GetURLString() FAIL, Got URL string without APIKey value 92536907")
	}

	cfg.Server.PKey = "1234"
	cfg.Server.SKey = "Rotten"
	cfg.Server.APIKey = ""
	cfg.Server.OmdbURL = ""
	url, ok = GetURLString(title, cfg)
	if !ok && url != "" {
		t.Errorf("GetURLString() FAIL, Got URL string without APIKey")
	}
}

func TestFindObjFromMapByKey(t *testing.T) {
	respBody := "{\"Title\":\"frozen\",\"Year\":\"2002\",\"Ratings\":[{\"Source\":\"Internet Movie Database\",\"Value\":\"5.9/10\"},{\"Source\":\"Rotten Tomatoes\",\"Value\":\"48%\"}],\"Response\":\"True\"}"
	//respByte := []byte(respBody)
	pKey := "Ratings"
	obj := map[string]interface{}{}
	if err := json.Unmarshal([]byte(respBody), &obj); err != nil {
		ProcessError(err)
	}
	rv, ok := FindObjFromMapByKey(obj, pKey)
	fmt.Println("Rating :", rv)
	if !ok {
		t.Errorf("FindObjFromMapByKey() FAIL, Expecting key match and result should not Nil")
	}

	respBody = "{\"Title\":\"frozen\",\"Year\":\"2002\",\"Response\":\"True\"}"
	//respByte = []byte(respBody)
	obj = map[string]interface{}{}
	if err := json.Unmarshal([]byte(respBody), &obj); err != nil {
		ProcessError(err)
	}
	pKey = "Year"
	_, ok = FindObjFromMapByKey(obj, pKey)
	if !ok {
		t.Errorf("FindObjFromMapByKey() FAIL, Expected result should be Nil")
	}

}

func TestGetMovieRatingByTitle(t *testing.T) {
	url := "http://www.omdbapi.com/?i=tt3896198&apikey=92536907&t=frozen"
	_, err := GetMovieRatingByTitle(url)
	if err != nil {
		t.Errorf("TestGetMovieRatingByTitle() FAIL, Expecting response body")
	}
}

func TestFindValueFromMapByValue(t *testing.T) {
	input := "[{\"Source\":\"Internet Movie Database\",\"Value\":\"5.9/10\"},{\"Source\":\"Rotten Tomatoes\",\"Value\":\"48%\"}]"
	obj := []interface{}{}
	if err := json.Unmarshal([]byte(input), &obj); err != nil {
		ProcessError(err)
	}
	sKey := "Rotten Tomatoes"
	_, v, _ := FindValueFromMapByValue(obj, sKey)
	if v != "48%" {
		t.Errorf("TestFindValueFromMapByValue() FAIL, Expecting value 48 but found %v", v)
	}
}
