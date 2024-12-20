package es

import (
	"fmt"
	"net/http"
)

func (c *Client) DeleteIndice(indice string) error {
	_, err := c.doRequest(http.MethodDelete, fmt.Sprintf("/%s", indice), nil)
	return err
}
