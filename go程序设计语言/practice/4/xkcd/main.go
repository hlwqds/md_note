package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

)

func parseSingleComicDataIntoCache(data []byte, ptr *os.File) error {
	cache := &Comic{}

	json.Unmarshal(data, cache)

	if cache.Link == "" {
		cache.Link = baseURLPath + "/" + string(cache.Num)
	}

	//其实没有必要
	_, ok := numIndex[cache.Num]
	if ok {
		fmt.Println("duplicate")
		return fmt.Errorf("duplicate")
	}

	numIndex[cache.Num] = true
	time := fmt.Sprintf("%s-%s", cache.Year, cache.Month)
	timeIndex[time] = append(timeIndex[time], *cache)

	ptr.Write(data)
	ptr.Write([]byte{'\n'})
	return nil
}

func parseSingleComicData(data []byte) error {
	cache := &Comic{}

	json.Unmarshal(data, cache)

	if cache.Link == "" {
		cache.Link = baseURLPath + "/" + string(cache.Num)
	}

	//其实没有必要
	_, ok := numIndex[cache.Num]
	if ok {
		fmt.Println("duplicate")
		return fmt.Errorf("duplicate")
	}

	numIndex[cache.Num] = true
	time := fmt.Sprintf("%s-%s", cache.Year, cache.Month)
	timeIndex[time] = append(timeIndex[time], *cache)
	return nil
}

func getSingleComic(index int, ptr *os.File) error {
	url := baseURLPath + "/" + strconv.Itoa(index) + "/" + descriptionObj
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	//decoder虽然方便，但是过程不能控制，我们需要读取数据后进行两步操作
	//	decoder := json.NewDecoder(resp.Body)
	//	if err := decoder.Decode(cache); err != nil {
	//		resp.Body.Close()
	//		return err
	//	}

	err = parseSingleComicDataIntoCache(data, ptr)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func readInfoFromCache(ptr *os.File) error {
	data := bufio.NewScanner(ptr)
	for data.Scan() {
		err := parseSingleComicData(data.Bytes())
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return nil
}

func main() {
	if ptr, err := os.Open(cacheFile); err != nil {
		if ptr, err := os.Create(cacheFile); err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			for i := 1; i <= 100; i++ {
				err := getSingleComic(i, ptr)
				if err != nil {
					fmt.Println(err)
					continue
				}
			}
			ptr.Close()
		}
	} else {
		if err := readInfoFromCache(ptr); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	fmt.Printf("%#v\n", timeIndex)
}
