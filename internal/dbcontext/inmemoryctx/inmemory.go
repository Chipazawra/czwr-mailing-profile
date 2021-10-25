package inmemoryctx

import (
	"fmt"
)

type ctx struct {
	receiversList map[string][]string
}

func New() *ctx {
	return &ctx{make(map[string][]string)}
}

func (c *ctx) Create(usr, receiver string) (uint, error) {

	c.receiversList[usr] = append(c.receiversList[usr], receiver)

	return uint(len(c.receiversList[usr]) - 1), nil
}

func (c *ctx) Delete(usr string, id uint) error {

	if _, exist := c.receiversList[usr]; !exist {
		return fmt.Errorf("There is no receiver list for user %v", usr)
	} else if int(id) > cap(c.receiversList[usr])-1 {
		return fmt.Errorf("There is no receiver with index %v", id)
	}

	copy(c.receiversList[usr][id:], c.receiversList[usr][id+1:])
	c.receiversList[usr][len(c.receiversList[usr])-1] = ""
	c.receiversList[usr] = c.receiversList[usr][:len(c.receiversList[usr])-1]

	return nil
}

func (c *ctx) Update(usr string, id uint, receiver string) error {

	if _, exist := c.receiversList[usr]; exist {

		if int(id) > cap(c.receiversList[usr])-1 {
			return fmt.Errorf("There is no receiver with index %v", id)
		}

	} else {
		return fmt.Errorf("There is no receiver list for user %v", usr)
	}

	c.receiversList[usr][id] = receiver

	return nil
}

func (c *ctx) Read(usr string) ([]string, error) {

	if _, exist := c.receiversList[usr]; !exist {
		return nil, fmt.Errorf("There is no receiver list for user %v", usr)
	}

	return c.receiversList[usr], nil
}
