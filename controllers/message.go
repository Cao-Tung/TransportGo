package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"svGo/models"
	"fmt"
)

// Operations about message
type MessageController struct {
	beego.Controller
}

// @Title Rec Message
// @Description rec message
// @Param	body		body 	models.Message	true		"The message content"
// @Success 200 {object} models.Message
// @Failure 403 body is empty
// @router / [post]
func (ms *MessageController) Post() {
	var message models.Message
	json.Unmarshal(ms.Ctx.Input.RequestBody, &message)
	models.RecMessage(message)
	ms.Data["json"] = message
	fmt.Println(message)
	ms.ServeJSON()
}


