package models

type TemplateCategory struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	MetaDescription string `json:"meta_description"`
	FullDescription string `json:"full_description"`
	Code            string `json:"code"`
	Sort            int    `json:"sort"`
}
