package category_nurzhas_store

type Category struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
}

type CategoryUpdate struct {
	Id          string  `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	ImageUrl    *string `json:"image_url"`
}
