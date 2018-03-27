package main

type KofaxDocument struct {
	UUID    string `json:"uuid"`
	From    string `json:"from"`
	To      string `json:"to"`
	State   string `json:"state"`
	Subject string `json:"subject"`
}

type KofaxXML struct {
	CObjectlist CObjectlist `xml:"objectlist"`
}

type CObjectlist struct {
	CPaglen     string    `xml:"paglen"`
	CSelection  string    `xml:"selection"`
	CNtotal     string    `xml:"ntotal"`
	CListoffset string    `xml:"listoffset"`
	CNlisted    string    `xml:"nlisted"`
	CObject     []CObject `xml:"object"`
}

type CObject struct {
	CCommon CCommon `xml:"common"`
	CHeader CHeader `xml:"header"`
}

type CHeader struct {
	General  General  `xml:"general"`
	Specific Specific `xml:"specific"`
}

type CCommon struct {
	CUUID   string `xml:"uuid"`
	CBlocks string `xml:"blocks"`
	CUpdate string `xml:"update"`
}

type General struct {
	Act                  string  `xml:"Act"`
	Box                  string  `xml:"Box"`
	Type                 string  `xml:"Type"`
	Deliver              Deliver `xml:"deliver"`
	DisplayFrom          string  `xml:"DisplayFrom"`
	DisplayRecipientList string  `xml:"DisplayRecipientList"`
	DisplayTo            string  `xml:"DisplayTo"`
	DisplaySubject       string  `xml:"DisplaySubject"`
	TimeStart            string  `xml:"TimeStart"`
	TimeEnd              string  `xml:"TimeEnd"`
	ContentErrorLevel    string  `xml:"ContentErrorLevel"`
}

type Specific struct {
	Import Import `xml:"import"`
}

type Import struct {
	Folder          string   `xml:"Folder"`
	TriggerFileUsed string   `xml:"TriggerFileUsed"`
	FileList        FileList `xml:"FileList"`
}

type FileList struct {
	FilePath string `xml:"FilePath"`
}

type Deliver struct {
	State         string `xml:"State"`
	TimeScheduled string `xml:"TimeScheduled"`
	Response      string `xml:"Response"`
	Tries         string `xml:"Tries"`
}
