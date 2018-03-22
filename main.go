package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	host = "hitmais0010"
	//host = "172.16.80.95"
)

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
		fmt.Println("State not valid, finded:", len(states), "  expected (uuids):", len(uuids))
	}

	// for count, uuid := range uuids {
	// 	fmt.Println(count, " ", uuid, "	STATE:", states[count], "	SELECTION:", selection, "\r")
	// }

	for count, uuid := range uuids {
		fmt.Println(count, " ", uuid, "	SELECTION:", selection, "\n")
	}

	for count, state := range states {
		fmt.Println(count, "	STATE:", state, "	SELECTION:", selection)
	}

}

func getFirstHtml() (string, error) {
	bodyString := ""
	for index := 0; index < 10; index++ {
		resp, err := http.Get(fmt.Sprintf("http://%s:25086/call/fax/inhistory", host))
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		ioutil.WriteFile("./firstpage.xml", body, 0644)

		bodyString = string(body[:])

	}

	return bodyString, nil
}

func getSelection(firstPageBody string) (string, error) {
	re := regexp.MustCompile("<c:selection>\\d{10}|<c:selection>(-)\\d{10}")
	//fmt.Println(re.FindString(firstPageBody))
	selection := re.FindString(firstPageBody)
	selection = strings.Replace(selection, "<c:selection>", "", -1)
	return selection, nil
}

func getXML(selection string) (string, error) {

	xmlTotal := ""
	for index := 0; index < 5000; index += 50 {
		resp, err := http.Get(fmt.Sprintf("http://%s:25086/call/fax/inhistory?selection=%s&index=%s", host, selection, strconv.Itoa(index)))
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
			WriteStringToFile("./all.xml", xmlTotal)
			return xmlTotal, nil
		}
		xmlTotal += bodyString

	}
	WriteStringToFile("./all.xml", xmlTotal)
	return xmlTotal, nil
}

func GetUUID(xmlTotal string) ([]string, error) {

	re := regexp.MustCompile("<c:uuid>[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}")
	uuid := re.FindAllString(xmlTotal, -1)
	for c, u := range uuid {
		uuid[c] = strings.Replace(u, "<c:uuid>", "", -1)
	}
	return uuid, nil
}

func GetState(xmlTotal string) ([]string, error) {

	re := regexp.MustCompile("<State>\\d{0,3}")
	//fmt.Println(re.FindAllString(xmlTotal))
	state := re.FindAllString(xmlTotal, -1)
	for c, s := range state {
		state[c] = strings.Replace(s, "<State>", "", -1)
	}

	return state, nil
}

func WriteStringToFile(filepath, s string) error {
	fo, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer fo.Close()

	_, err = io.Copy(fo, strings.NewReader(s))
	if err != nil {
		return err
	}

	return nil
}
