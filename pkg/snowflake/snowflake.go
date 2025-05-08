package snowflake

import (
	"github.com/bwmarrin/snowflake"
)

var orderNode *snowflake.Node
var paymentNode *snowflake.Node

func Init() {
	var err error
	orderNode, err = snowflake.NewNode(1)
	if err != nil {
		panic(err.Error())
	}
	paymentNode, err = snowflake.NewNode(2)
	if err != nil {
		panic(err.Error())
	}
}
func GenerateOrderID() string {
	if orderNode == nil {
		Init()
	}
	return orderNode.Generate().String()
}

func GeneratePaymentID() string {
	if paymentNode == nil {
		Init()
	}
	return paymentNode.Generate().String()
}
