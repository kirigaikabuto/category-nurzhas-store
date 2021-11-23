package category_nurzhas_store

import (
	"fmt"
	"github.com/google/uuid"
)

type OrderService interface {
	CreateOrder(cmd *CreateOrderCommand) (*Order, error)
	ListOrder(cmd *ListOrderCommand) ([]Order, error)
}

type orderService struct {
	orderStore      OrderStore
	telegramService TelegramService
}

func NewOrderService(o OrderStore, t TelegramService) OrderService {
	return &orderService{orderStore: o, telegramService: t}
}

func (o *orderService) CreateOrder(cmd *CreateOrderCommand) (*Order, error) {
	order := &Order{}
	order.Id = uuid.New().String()
	order.BuildingType = cmd.BuildingType
	order.Color = cmd.Color
	order.PanelWidth = cmd.PanelType
	order.PanelDepth = cmd.PanelDepth
	order.PanelType = cmd.PanelType
	order.InsulationType = cmd.InsulationType
	order.LayoutType = cmd.LayoutType
	order.Height = cmd.Height
	order.Width = cmd.Width
	order.Length = cmd.Length
	message := ""
	message += fmt.Sprint("<pre><b>Калькулятор</b></pre>\n")
	message += fmt.Sprintf("<pre>%s:<b>%s</b></pre>\n", "Тип", cmd.PanelType)
	message += fmt.Sprintf("<pre>%s:<b>%s</b></pre>\n", "Утеплитель", cmd.InsulationType)
	message += fmt.Sprintf("<pre>%s:<b>%s</b></pre>\n", "Толщина панели", cmd.PanelDepth)
	message += fmt.Sprintf("<pre>%s:<b>%s</b></pre>\n", "Цвет панели", cmd.Color)
	err := o.telegramService.SendTelegramMessage("", message, "HTML")
	if err != nil {
		return nil, err
	}
	return o.orderStore.CreateOrder(order)
}

func (o *orderService) ListOrder(cmd *ListOrderCommand) ([]Order, error) {
	return o.orderStore.ListOrder()
}
