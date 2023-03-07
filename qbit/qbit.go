package qbit

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/imroc/req/v3"
)

type Torrent struct {
	Name     string `json:"name"`
	Tracker  string `json:"tracker"`
	Category string `json:"category"`
	Hash     string `json:"hash"`
	SavePath string `json:"save_path"`
	Tags     string `json:"tags"`
}

type qbit struct {
	Cli *req.Client
}

func NewClient() (*qbit, error) {
	payload := map[string]string{
		"username": os.Getenv("QB_USER"),
		"password": os.Getenv("QB_PASS"),
	}

	cli := req.C().
		SetBaseURL(os.Getenv("QB_URL"))

	res, err := cli.R().
		SetFormData(payload).
		Post("/api/v2/auth/login")

	if err != nil {
		return nil, err
	}

	cookie := res.GetHeader("set-cookie")

	if cookie == "" {
		return nil, errors.New("No cookie found")
	}

	cli.SetCommonHeader("Cookie", cookie)

	return &qbit{cli}, nil
}

func (q *qbit) GetTorrentsInfo() ([]Torrent, error) {
	var torrents []Torrent

	res, err := q.Cli.R().
		Get("api/v2/torrents/info")

	if err != nil {
		return torrents, err
	}

	json.Unmarshal(res.Bytes(), &torrents)

	return torrents, nil
}

func (q *qbit) PostTorrentsAddTags(hashes []string, tags string) error {
	payload := map[string]string{
		"hashes": strings.Join(hashes, "|"),
		"tags":   tags,
	}

	_, err := q.Cli.R().
		SetFormData(payload).
		Post("api/v2/torrents/addTags")

	if err != nil {
		return err
	}

	return nil
}

func (q *qbit) PostTorrentsRemoveTags(hashes []string, tags string) error {
	payload := map[string]string{
		"hashes": strings.Join(hashes, "|"),
		"tags":   tags,
	}

	_, err := q.Cli.R().
		SetFormData(payload).
		Post("api/v2/torrents/removeTags")

	if err != nil {
		return err
	}

	return nil
}
