package category_nurzhas_store

import (
	"encoding/json"
	"github.com/djumanoff/amqp"
	setdata_common "github.com/kirigaikabuto/setdata-common"
)

type CategoryAmqpEndpoints struct {
	ch setdata_common.CommandHandler
}

func NewCategoryAmqpEndpoints(ch setdata_common.CommandHandler) CategoryAmqpEndpoints {
	return CategoryAmqpEndpoints{ch: ch}
}

func (c *CategoryAmqpEndpoints) MakeCreateCategoryAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &CreateCategoryCommand{}
		err := json.Unmarshal(message.Body, cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := c.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonData}
	}
}

func (c *CategoryAmqpEndpoints) MakeGetCategoryAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &GetCategoryCommand{}
		err := json.Unmarshal(message.Body, cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := c.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonData}
	}
}

func (c *CategoryAmqpEndpoints) MakeUpdateCategoryAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &UpdateCategoryCommand{}
		err := json.Unmarshal(message.Body, cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := c.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonData}
	}
}

func (c *CategoryAmqpEndpoints) MakeDeleteCategoryAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &DeleteCategoryCommand{}
		err := json.Unmarshal(message.Body, cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := c.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonData}
	}
}

func (c *CategoryAmqpEndpoints) MakeListCategoryAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &ListCategoryCommand{}
		err := json.Unmarshal(message.Body, cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := c.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonData}
	}
}
