package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
	"slices"
	"strings"
)

type Citation struct {
	Quote   string
	Movie   string
	Speaker string
}

func FetchPage(ctx context.Context) (pc pageConsumer, err error) {
	resp, err := http.Get(baseURLCitation)
	if err != nil {
		return
	}

	info, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	info_str := html.UnescapeString(string(info))

	var consumer pageConsumer
	err = xml.Unmarshal([]byte(info_str), &consumer)
	if err != nil {
		return pageConsumer{}, err
	}

	return consumer, nil
}

func GetCitation(ctx context.Context) (cit Citation, err error) {
	pc, err := FetchPage(ctx)
	if err != nil {
		return
	}

	idx_corps := slices.IndexFunc(pc.Total, func(c divConsumer) bool { return c.ID == corps_id })
	idx_signature := slices.IndexFunc(pc.Total, func(c divConsumer) bool { return c.ID == signature_id })

	corps := pc.Total[idx_corps]
	signature := pc.Total[idx_signature]

	cit = Citation{
		Quote:   corps.Text,
		Movie:   strings.TrimSuffix(strings.TrimSpace(signature.Title), "\n"),
		Speaker: strings.TrimSuffix(strings.TrimSpace(signature.Text), "\n"),
	}
	return
}
