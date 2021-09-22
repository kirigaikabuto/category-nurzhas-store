package category_nurzhas_store

import "github.com/google/uuid"

type CategoryService interface {
	CreateCategory(cmd *CreateCategoryCommand) (*Category, error)
	UpdateCategory(cmd *UpdateCategoryCommand) (*Category, error)
	ListCategory(cmd *ListCategoryCommand) ([]Category, error)
	GetCategory(cmd *GetCategoryCommand) (*Category, error)
	DeleteCategory(cmd *DeleteCategoryCommand) error
}

type categoryService struct {
	categoryStore CategoryStore
}

func NewCategoryService(cStore CategoryStore) CategoryService {
	return &categoryService{categoryStore: cStore}
}

func (c *categoryService) CreateCategory(cmd *CreateCategoryCommand) (*Category, error) {
	category := &Category{
		Id:               uuid.New().String(),
		Name:             cmd.Name,
		SmallDescription: cmd.SmallDescription,
		BigDescription:   cmd.BigDescription,
		ImageUrl:         cmd.ImageUrl,
	}
	return c.categoryStore.CreateCategory(category)
}

func (c *categoryService) UpdateCategory(cmd *UpdateCategoryCommand) (*Category, error) {
	categoryUpdate := &CategoryUpdate{Id: cmd.Id}
	oldCategory, err := c.categoryStore.GetCategory(cmd.Id)
	if err != nil {
		return nil, err
	}
	if cmd.Name != "" && oldCategory.Name != cmd.Name {
		categoryUpdate.Name = &cmd.Name
	}
	if cmd.SmallDescription != "" && oldCategory.SmallDescription != cmd.SmallDescription {
		categoryUpdate.SmallDescription = &cmd.SmallDescription
	}
	if cmd.BigDescription != "" && oldCategory.BigDescription != cmd.BigDescription {
		categoryUpdate.BigDescription = &cmd.BigDescription
	}
	if cmd.ImageUrl != "" && oldCategory.ImageUrl != cmd.ImageUrl {
		categoryUpdate.ImageUrl = &cmd.ImageUrl
	}
	return c.categoryStore.UpdateCategory(categoryUpdate)
}

func (c *categoryService) ListCategory(cmd *ListCategoryCommand) ([]Category, error) {
	return c.categoryStore.ListCategory()
}

func (c *categoryService) GetCategory(cmd *GetCategoryCommand) (*Category, error) {
	return c.categoryStore.GetCategory(cmd.Id)
}

func (c *categoryService) DeleteCategory(cmd *DeleteCategoryCommand) error {
	return c.categoryStore.DeleteCategory(cmd.Id)
}
