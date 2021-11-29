package category_nurzhas_store

type Order struct {
	Id             string `json:"id"`
	BuildingType   string `json:"building_type"`
	Width          string `json:"width"`
	Height         string `json:"height"`
	Length         string `json:"length"`
	PanelType      string `json:"panel_type"`
	InsulationType string `json:"insulation_type"`
	PanelDepth     string `json:"panel_depth"`
	LayoutType     string `json:"layout_type"`
	PanelWidth     string `json:"panel_width"`
	Color          string `json:"color"`
	TotalSum       string `json:"total_sum"`
}
