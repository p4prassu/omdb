//////////////////////////////////////////////////////////////////////////
// utils go file for omdb module.                          				//
//                                                           	   		//
// Author: Prasad Potipireddi                   		                //
// Date: Jan 13th, 2020.                                               	//
// Since: CX Cloud Release.                                            	//
// Copyright (c) 2018 Cisco Systems. All rights reserved.              	//
//                                                                      //
//////////////////////////////////////////////////////////////////////////

// Package utils is having common method being used in main.go
// all utils functions can be reused for further use cases.
package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

// Configurations struct is to used export the application configurations
// Unmarshall the []byte slice data into Configuration structure
type Configurations struct {
	Server struct {
		PKey    string `yaml:"pKey", json:"pKey", envconfig:"pKey"`          //Pkey primary key ex. Ratings
		SKey    string `yaml:"sKey", json:"sKey", envconfig:"sKey"`          //Secondary Key Ex. Rotan Tomotos
		OmdbURL string `yaml:"omdbURL", json:"omdbURL", envconfig:"omdbURL"` // URL of the omdb site
		APIKey  string `yaml:"apiKey", json:"apiKey", envconfig:"apiKey"`    // APIKey of omdb user
		OmdbID  string `yaml:"omdbID", json:"omdbID", envconfig:"omdbID"`    // omdbID is omdb user id
	} `yaml: "server", json:"server"` // Server details to connect
}

// type Configurations struct {
// 	pKey    string
// 	sKey    string
// 	omdbURL string
// 	apiKey  string
// 	omdbID  string
// }

// ProcessError print the error msg and exit with status code 2
// which is error
func ProcessError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

// LoadConfiguration load applciation configurations into Configurations struct
// it accepts the file name and file type (ex. yaml) and reference to Configuration type
// it support for both json and yaml parser
func LoadConfiguration(file string, fType string, cfg *Configurations) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		ProcessError(err)
	}
	if fType == "yaml" {
		decoder := yaml.NewDecoder(f)
		err = decoder.Decode(cfg)
	}
}

// GetURLString prepare the url based on confiuration information to pull omdb movie details
func GetURLString(title string, config Configurations) (string, bool) {
	if config.Server.PKey != "" && config.Server.OmdbURL != "" && config.Server.SKey != "" && config.Server.OmdbID != "" {
		return config.Server.OmdbURL + "i=" + config.Server.OmdbID + "&apiKey=" + config.Server.APIKey + "&t=" + title, true
	}
	return "", false
}

// GetMovieRatingByTitle returns the byte slice by taking URL string
// It captures the response from omdb movie get request.
func GetMovieRatingByTitle(url string) ([]byte, error) {
	bBytes := []byte{}
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Printf("response status code: %v", resp.StatusCode)
	} else {
		bBytes, err = ioutil.ReadAll(resp.Body)
	}
	return bBytes, err
}

// GetMapValuesAsSlice return slice of values as string from a given map
func GetMapValuesAsSlice(m map[string]interface{}) []string {
	values := []string{}
	for _, v := range m {
		values = append(values, v.(string))
	}
	return values
}

// FindItemFromSlice return the true, index if it finds item
// else return false, length of slice
func FindItemFromSlice(s []string, item string) (int, bool) {
	for i, v := range s {
		if v == item {
			return i, true
		}
	}
	return len(s), false
}

// FindObjFromMapByKey returns interface and bool
// It is a recurrsive function to iterate over the nested json until it finds
// interface object / item values for a given key
// hence it accepts interface and primary key ex.Ratings
func FindObjFromMapByKey(obj interface{}, key string) (interface{}, bool) {
	mobj, ok := obj.(map[string]interface{})
	if !ok {
		return nil, false
	}
	// return the value if it match key
	for k, v := range mobj {
		if k == key {
			return v, true
		}
		// return the map object for recursion
		if m, ok := v.(map[string]interface{}); ok {
			if res, ok := FindObjFromMapByKey(m, key); ok {
				return res, true
			}
		}
		// return the slice object for recursion
		if va, ok := v.([]interface{}); ok {
			for _, a := range va {
				if res, ok := FindObjFromMapByKey(a, key); ok {
					return res, true
				}
			}
		}
	}
	// return nil and false if not match key
	return nil, false
}

//FindValueFromMapByValue return values of map based on condition
// Ex. [map[string]string,map[string]string ...]/
func FindValueFromMapByValue(i interface{}, sKey string) (string, string, bool) {
	if va, ok := i.([]interface{}); ok {
		for _, items := range va {
			// getting values from a given map Ex. map[String]string -> []string
			values := GetMapValuesAsSlice(items.(map[string]interface{}))
			// finding the secondKey (Rotan Tomotos) in value slices
			if idx, ok := FindItemFromSlice(values, sKey); ok {
				// Printing the final results values
				//fmt.Println(values[idx] + " : " + values[(len(values)-1)-idx])
				return values[idx], values[(len(values)-1)-idx], true
			}

		}
	}
	return "", "", false
}
