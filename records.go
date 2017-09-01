package gandidnsclient

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type RecordType string

const (
	TypeA     = RecordType("A")
	TypeCName = RecordType("CName")
	TypeAAAA  = RecordType("AAAA")
	TypeCAA   = RecordType("CAA")
	TypeCDS   = RecordType("CDS")
	TypeDNAME = RecordType("DNAME")
	TypeDS    = RecordType("DS")
	TypeLOC   = RecordType("LOC")
	TypeMX    = RecordType("MX")
	TypeNS    = RecordType("NS")
	TypePTR   = RecordType("PTR")
	TypeSPF   = RecordType("SPF")
	TypeSRV   = RecordType("SRV")
	TypeSSHFP = RecordType("SSHFP")
	TypeTLSA  = RecordType("TLSA")
	TypeTXT   = RecordType("TXT")
	TypeWKS   = RecordType("WKS")
)

type ZoneRecord struct {
	Type       RecordType
	TimeToLive time.Duration
	Name       string
	Values     []string
}

func (c Client) ListRecords(zoneID string) ([]ZoneRecord, error) {
	req := c.request("GET", defaultBaseUrl+"/zones/"+zoneID+"/records", strings.Reader{""})
	resp, err := http.Client{}.Do(req)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	records := make([]ZoneRecord, 0)
	err := json.Unmarshal(body, &records)
	if err != nil {
		return err
	}
	return records
}
