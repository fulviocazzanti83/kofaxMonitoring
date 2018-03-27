package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args

	if len(args) != 5 {
		fmt.Println(`USE kofaxmonitoring --host myhost.com --pages 50`)
		os.Exit(0)
	}

	host := args[2]
	pages, err := strconv.Atoi(args[4])

	if err != nil {
		fmt.Println(`USE kofaxmonitoring --host myhost.com --pages 50`)
		os.Exit(0)
	}

	kd, _ := getKofaxDocuments(pages, host)
	if len(kd) == 0 {
		fmt.Println("No documents")
		return
	}
	fmt.Printf("%+v", kd)

}

func getKofaxDocuments(pageStop int, host string) ([]KofaxDocument, error) {
	selection, err := getSelection(host)
	if err != nil {
		fmt.Errorf(err.Error())
		return nil, err
	}

	kofaxpages, err := getData(selection, pageStop, host)
	if err != nil {
		fmt.Errorf(err.Error())
		return nil, err
	}

	var kofaxDocuments []KofaxDocument

	for _, page := range kofaxpages {

		for _, obj := range page.CObjectlist.CObject {

			kd := KofaxDocument{From: obj.CHeader.General.DisplayFrom,
				To:      obj.CHeader.General.DisplayTo,
				State:   obj.CHeader.General.Deliver.State,
				UUID:    obj.CCommon.CUUID,
				Subject: obj.CHeader.General.DisplaySubject,
			}

			kofaxDocuments = append(kofaxDocuments, kd)

		}
	}
	return kofaxDocuments, nil
}

func getSelection(host string) (string, error) {
	bodyString := ""

	resp, err := http.Get(fmt.Sprintf("http://%s:25086/call/fax/inhistory", host))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	bodyString = string(body[:])

	k, err := parseXML(bodyString)
	if err != nil {
		return "", err
	}
	return k.CObjectlist.CSelection, nil
}

func getData(selection string, pageStop int, host string) ([]KofaxXML, error) {

	var xmlTotal []KofaxXML

	pages := pageStop * 50

	for index := 0; index < pages; index += 50 {
		resp, err := http.Get(fmt.Sprintf("http://%s:25086/call/fax/inhistory?selection=%s&index=%s", host, selection, strconv.Itoa(index)))
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		bodyString := string(body[:])

		if len(bodyString) == 397 {
			return xmlTotal, nil
		}
		xmlPage, err := parseXML(bodyString)
		if err != nil {
			return nil, err
		}
		xmlTotal = append(xmlTotal, xmlPage)

	}
	return xmlTotal, nil
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

func parseXML(kofaxXml string) (KofaxXML, error) {

	result := KofaxXML{}
	data := kofaxXml
	data = strings.Replace(data, "<c:", "<", -1)
	data = strings.Replace(data, "</c:", "</", -1)

	err := xml.Unmarshal([]byte(data), &result)
	if err != nil {
		log.Printf("error: %v", err)
	}

	return result, nil
}
