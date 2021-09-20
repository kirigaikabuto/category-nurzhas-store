package category_nurzhas_store

type CategoryStore interface {
	CreateCategory(category *Category) (*Category, error)
	UpdateCategory(category *CategoryUpdate) (*Category, error)
	ListCategory() ([]Category, error)
	GetCategory(id string) (*Category, error)
	DeleteCategory(id string) error
}
