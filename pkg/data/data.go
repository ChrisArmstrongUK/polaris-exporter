package data

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	polaris "github.com/fairwindsops/polaris/pkg/validator"
)

type Data struct {
	AuditData polaris.AuditData
}

func (d *Data) JSON() (string, error) {
	json, err := json.MarshalIndent(d.AuditData, "", "  ")
	if err != nil {
		return "", fmt.Errorf("STRING: %v", err)
	}
	return string(json), nil
}

func (d *Data) MonitorTarget(target url.URL, interval time.Duration, timeout time.Duration) {
	log.Printf("url:%v\nscheme:%v host:%v Path:%v\n\n", target, target.Scheme, target.Host, target.Path)
	switch target.Scheme {
	case "http", "https":
		log.Println("Target is http")
		d.MonitorHTTP(interval, target, timeout)
	case "file":
		log.Println("Target is a file")
		d.MonitorFile(interval, fmt.Sprintf("%s%s", target.Host, target.Path))
	default:
		log.Println("Invalid target")
	}
}

func (d *Data) MarshalFromFile(filepath string) error {
	contents, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	ad := &polaris.AuditData{}

	json.Unmarshal(contents, ad)
	d.AuditData = *ad
	return nil
}

func (d *Data) MonitorFile(interval time.Duration, filepath string) {
	go func() {
		for {
			log.Println("Reading file")
			err := d.MarshalFromFile(filepath)
			if err != nil {
				log.Println(err)
			}

			time.Sleep(interval)
		}
	}()
}

func (d *Data) MarshalFromHTTP(target url.URL, httpTimeout time.Duration) error {
	httpClient := http.Client{Timeout: httpTimeout}

	req, err := http.NewRequest(http.MethodGet, target.String(), nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "polaris-exporter")

	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	ad := &polaris.AuditData{}

	err = json.Unmarshal(body, ad)
	if err != nil {
		return err
	}

	d.AuditData = *ad

	return nil
}

func (d *Data) MonitorHTTP(interval time.Duration, target url.URL, httpTimeout time.Duration) {
	go func() {
		for {
			log.Println("Reading target")
			err := d.MarshalFromHTTP(target, httpTimeout)
			if err != nil {
				log.Println(err)
			}
			time.Sleep(interval)
		}
	}()
}
