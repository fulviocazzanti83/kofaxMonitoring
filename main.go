package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	// requiredEnvs := []string{
	// 	"DB_HOST",
	// 	"DB_USER",
	// 	"DB_NAME",
	// 	"PUBSUB_TOPIC",
	// 	"GCE_PROJECT",
	// 	"APP_PORT",
	// }

	// for _, val := range requiredEnvs {
	// 	if os.Getenv(val) == "" {
	// 		log.Fatalf("%s ENV NOT FOUND", val)
	// 	}
	// }

}

func main() {
	firstPage, _ := getFirstHtml()
	if len(firstPage) == 0 {
		panic("Page not valid")
	}

	selection, _ := getSelection(firstPage)
	if len(selection) == 0 {
		panic("Selection not valid")
	}

	xmlTotal, _ := getXML(selection)
	if len(xmlTotal) == 0 {
		panic("XML not valid")
	}

	uuids, _ := GetUUID(xmlTotal)
	if len(uuids) == 0 {
		panic("UUID not valid")
	}

	states, _ := GetState(xmlTotal)
	if len(states) != len(uuids) {
		panic("State not valid")
	}

	for count, uuid := range uuids {
		fmt.Println(count, " ", uuid, "	STATE:", states[count], "	SELECTION:", selection)
	}

}

func getFirstHtml() (string, error) {
	bodyString := ""
	for index := 0; index < 10; index++ {
		resp, err := http.Get("http://172.16.80.95:25086/call/fax/inhistory")
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		//err = ioutil.WriteFile("/tmp/firstpage.xml", body, 0644)

		bodyString = string(body[:])

	}
	return bodyString, nil
}

func getSelection(firstPageBody string) (string, error) {
	re := regexp.MustCompile("<c:selection>\\d{10}")
	//fmt.Println(re.FindString(firstPageBody))
	selection := re.FindString(firstPageBody)
	selection = strings.Replace(selection, "<c:selection>", "", -1)
	return selection, nil
}

func getXML(selection string) (string, error) {

	xmlTotal := ""
	for index := 0; index < 5000; index += 50 {
		resp, err := http.Get(fmt.Sprintf("http://172.16.80.95:25086/call/fax/inhistory?selection=%s&index=%s", selection, strconv.Itoa(index)))
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		bodyString := string(body[:])

		if len(bodyString) == 397 {
			return xmlTotal, nil
		}
		xmlTotal += bodyString

	}
	return xmlTotal, nil
}

func GetUUID(xmlTotal string) ([]string, error) {

	re := regexp.MustCompile("[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}")
	//fmt.Println(re.FindAllString(xmlTotal))
	uuid := re.FindAllString(xmlTotal, -1)
	return uuid, nil
}

func GetState(xmlTotal string) ([]string, error) {

	re := regexp.MustCompile("<State>\\d{3}")
	//fmt.Println(re.FindAllString(xmlTotal))
	state := re.FindAllString(xmlTotal, -1)
	for c, s := range state {
		state[c] = strings.Replace(s, "<State>", "", -1)
	}

	return state, nil
}
