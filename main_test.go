package main

import (
	"fmt"
	"testing"
)

func TestGetFirstHtml(t *testing.T) {
	firstPage, _ := getFirstHtml()
	if len(firstPage) == 0 {
		t.Errorf("Page lenght")
	}
}
func TestGetSelection(t *testing.T) {

	firstPage, _ := getFirstHtml()
	if len(firstPage) == 0 {
		t.Errorf("Page lenght")
	}

	selection, _ := getSelection(firstPage)
	if len(selection) == 0 {
		t.Errorf("Selection not findeds")
	}

}

func TestGetXML(t *testing.T) {

	firstPage, _ := getFirstHtml()
	if len(firstPage) == 0 {
		t.Errorf("Page not valid")
	}

	selection, _ := getSelection(firstPage)
	if len(selection) == 0 {
		t.Errorf("Selection not valid")
	}

	xml, _ := getXML(selection)
	if len(xml) == 0 {
		t.Errorf("XML not valid")
	}

}

func TestGetUUID(t *testing.T) {

	firstPage, _ := getFirstHtml()
	if len(firstPage) == 0 {
		t.Errorf("Page not valid")
	}

	selection, _ := getSelection(firstPage)
	if len(selection) == 0 {
		t.Errorf("Selection not valid")
	}

	xmlTotal, _ := getXML(selection)
	if len(xmlTotal) == 0 {
		t.Errorf("XML not valid")
	}

	uuids, _ := GetUUID(xmlTotal)
	if len(uuids) == 0 {
		t.Errorf("UUID not valid")
	}

	for count, uuid := range uuids {
		fmt.Println(count, " ", uuid)
	}

}

func TestGetSate(t *testing.T) {

	firstPage, _ := getFirstHtml()
	if len(firstPage) == 0 {
		t.Errorf("Page not valid")
	}

	selection, _ := getSelection(firstPage)
	if len(selection) == 0 {
		t.Errorf("Selection not valid")
	}

	xmlTotal, _ := getXML(selection)
	if len(xmlTotal) == 0 {
		t.Errorf("XML not valid")
	}

	uuids, _ := GetUUID(xmlTotal)
	if len(uuids) == 0 {
		t.Errorf("UUID not valid")
	}

	states, _ := GetState(xmlTotal)
	if len(states) != len(uuids) {
		t.Errorf("State not valid")
	}

	for count, uuid := range uuids {
		fmt.Println(count, " ", uuid, "	STATE:", states[count], "	SELECTION:", selection)
	}

}
