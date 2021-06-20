package models

import (
	"encoding/json"
	"strings"
)

type Template struct {
	ID              string            `json:"id"`
	RealID          int               `json:"real_id"`
	Name            string            `json:"name"`
	NameSlug        string            `json:"name_slug"`
	Lang            string            `json:"lang"`
	MetaDescription string            `json:"meta_description"`
	FullDescription string            `json:"full_description"`
	Category        string            `json:"category"`
	CategoryInfo    *TemplateCategory `json:"category_info"`
	Tags            map[string]string `json:"tags"`
	Mark            string            `json:"mark"`       //
	MarkCount       int               `json:"mark_count"` //
	Body            string            `json:"body"`       //
	Owner           string            `json:"owner"`
	Created         DateTimeType      `json:"created"`
	Preview         string            `json:"preview"`
	IsStructure     bool              `json:"is_structure"`
}

func (t *Template) UnmarshalJSON(data []byte) error {
	raw := string(data)
	data = []byte(strings.ReplaceAll(raw, `"category_info": []`, `"category_info": null`))

	type template Template
	if err := json.Unmarshal(data, (*template)(t)); err != nil {
		return err
	}
	return nil
}
