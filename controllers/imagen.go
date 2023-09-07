package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/plantillas_crud/models"
	service "github.com/udistrital/plantillas_crud/services/plantilla.service"
	"github.com/udistrital/utils_oas/time_bogota"
)

// ImagenController operations for Imagen
type ImagenController struct {
	beego.Controller
}

// URLMapping ...
func (c *ImagenController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Imagen
// @Param	body		body 	models.Imagen	true		"body for Imagen content"
// @Success 201 {object} models.Imagen
// @Failure 403 body is empty
// @router / [post]
func (c *ImagenController) Post() {
	horaRegistro := time_bogota.TiempoBogotaFormato()
	var v models.Plantilla
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.FechaCreacion = horaRegistro
		v.FechaModificacion = horaRegistro
		fmt.Println("resultado ", v)
		if err := service.Create(v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			logs.Error(err)
			//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = err
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get Imagen by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Imagen
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ImagenController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	v, err := service.ReadOne(idStr)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get Imagen
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Imagen
// @Failure 403
// @router / [get]
func (c *ImagenController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Imagen
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Imagen	true		"body for Imagen content"
// @Success 200 {object} models.Imagen
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ImagenController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Imagen
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ImagenController) Delete() {

}
