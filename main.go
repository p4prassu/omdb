//////////////////////////////////////////////////////////////////////////
// main go file is accepting  movie title from user as command line arg //
// and provide the ratings												//
//                                                           	   		//
// Author: Prasad Potipireddi               		                    //
// Date: Jan 13th, 2020.                                               	//
// Since: CX Cloud Release.                                            	//
// Copyright (c) 2018 Cisco Systems. All rights reserved.              	//
//                                                                      //
//////////////////////////////////////////////////////////////////////////

// Package main is driver
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	u "omdb/utils"
	"os"
)

func main() {
	/*
		// using flag module, trying to accept more than one arguments like confFileName, confType
		cf := flag.String("cfile", "config.json", "config file name")
		cfType := flag.String("ctype", "json", "config file type")
		cfile := "./conf/" + *cf
		u.LoadConfiguration(cfile, *cfType, &cfg)
	*/

	// flag module is used for accepting the command line arguments
	// In this case, we are expecting only title from user (No confFileName and confType)
	var title string
	flag.StringVar(&title, "title", "frozen", "a string var")
	flag.Parse()

	var cfg u.Configurations

	// setting of configuration file path and type
	cfile := "./conf/config.yml"
	ctype := "yaml"

	// LoadConfiguration loads the configurations from config file
	// by passing cfg object reference
	u.LoadConfiguration(cfile, ctype, &cfg)

	// Get URL string
	url, ok := u.GetURLString(title, cfg)
	if !ok {
		fmt.Println("Some or all Configuration attributes missing")
		os.Exit(1)
	}

	// Getting moving rating details from omdb get call response by using title
	bBytes, err := u.GetMovieRatingByTitle(url)
	if err != nil {
		u.ProcessError(err)
	}

	// Unmarshalling byte slice to nested map objects similar to nested json
	obj := map[string]interface{}{}
	if err := json.Unmarshal(bBytes, &obj); err != nil {
		u.ProcessError(err)
	}

	// Getting object matching with key Rating  ex. PKey - Ratings
	if rv, ok := u.FindObjFromMapByKey(obj, cfg.Server.PKey); ok {
		switch v := rv.(type) {
		case string:
			fmt.Printf("Rating as String -> %s\n", v)
		case fmt.Stringer:
			fmt.Printf(" Rating as stringer interface -> %s\n", v.String())
		case int:
			fmt.Printf(" Rating as int -> %d\n", v)
		case interface{}:
			// Rating object has slice of map objects ex. [map[string]string,map[string]string ...]
			if k, v, ok := u.FindValueFromMapByValue(v, cfg.Server.SKey); ok {
				fmt.Println(k + " : " + v)
			} else {
				fmt.Println("Some thing wrong: findValueFromMapByValue func could not able to find cfg.Server.SKey")
			}
		default:
			fmt.Printf("Rating = %v, ok = %v\n", v, ok)
		}
	}

}
