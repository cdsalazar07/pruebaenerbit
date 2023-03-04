package handlers

import (
	"enerbit/prueba/db"
	"enerbit/prueba/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListEMeter	godoc
// @Summary 	Listado de todos los medidores de energía
// @Description Muestra todos los registros de medidores de energía que existan en la base de datos
// @Success 	200 {object} []models.Datameds{}
// @Router		/listemeter [get]
func ListEMeter(ctx *gin.Context) {
	data, err := db.ListEMeter()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.StandardResponse{
			Message: fmt.Sprintf("error: %s", err.Error()),
		})

		return
	}

	ctx.JSON(http.StatusOK, data)
}

// LastEMeter	godoc
// @Summary 	Último medidor de energía instalado
// @Description Muestra el último medidor de energía instalado en la base de datos
// @Param		brand path string true "Data"
// @Param		serial path string true "Data"
// @Success 	200 {object} models.Datameds{}
// @Router		/lastemeter/{brand}/{serial} [get]
func LastEMeter(ctx *gin.Context) {
	data := models.LastEMeterModel{
		Brand:  ctx.Param("brand"),
		Serial: ctx.Param("serial"),
	}

	response, err := db.LastEMeter(data.Serial, data.Brand)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.StandardResponse{
			Message: fmt.Sprintf("error: %s", err.Error()),
		})

		return
	}

	ctx.JSON(http.StatusOK, response)
}

// ListDisableEMeter	godoc
// @Summary 			Listado de todos los medidores de energía deshabilitados
// @Description 		Muestra todos los registros de medidores de energía deshabilitados que existan en la base de datos
// @Success 			200 {object} []models.Datameds{}
// @Router				/listdisableemeter [get]
func ListDisableEMeter(ctx *gin.Context) {
	data, err := db.ListDisableEMeter()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.StandardResponse{
			Message: fmt.Sprintf("error: %s", err.Error()),
		})

		return
	}

	ctx.JSON(http.StatusOK, data)
}

// NewEMeter	godoc
// @Summary 	Creador de medidor de energía
// @Description Crea un nuevo medidor de energía en la base de datos
// @Param		args body models.NewEMeterModel true "Data"
// @Success 	200 {object} models.StandardResponse{}
// @Router		/newemeter [post]
func NewEMeter(ctx *gin.Context) {
	var data models.NewEMeterModel
	err := ctx.BindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.StandardResponse{
			Message: fmt.Sprintf("error: %s", err.Error()),
		})

		return
	}

	err = db.NewEMeter(data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.StandardResponse{
			Message: fmt.Sprintf("error: %s", err.Error()),
		})

		return
	}

	ctx.JSON(http.StatusOK, models.StandardResponse{
		Message: "ok",
	})
}

// DeleteEMeter		godoc
// @Summary 		Elimina medidor de energía
// @Description 	Elimina un medidor de energía en la base de datos
// @Param			args body models.DeleteEMeterModel true "Data"
// @Success 		200 {object} models.StandardResponse{}
// @Router			/deleteemeter [post]
func DeleteEMeter(ctx *gin.Context) {
	var data models.DeleteEMeterModel
	err := ctx.BindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.StandardResponse{
			Message: fmt.Sprintf("error: %s", err.Error()),
		})

		return
	}

	err = db.DeleteEMeter(data.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.StandardResponse{
			Message: fmt.Sprintf("error: %s", err.Error()),
		})

		return
	}

	ctx.JSON(http.StatusOK, models.StandardResponse{
		Message: "ok",
	})
}

// UpdateEMeterStatus	godoc
// @Summary 			Actualiza medidor de energía
// @Description 		Actualiza el estado de un medidor de energía en la base de datos
// @Param				counter body models.UpdateEMeterStatusModel true "Update"
// @Success 			200 {object} models.StandardResponse{}
// @Router				/updateemeterstatus [post]
func UpdateEMeterStatus(ctx *gin.Context) {
	var data models.UpdateEMeterStatusModel
	err := ctx.BindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.StandardResponse{
			Message: fmt.Sprintf("error: %s", err.Error()),
		})

		return
	}

	err = db.UpdateEMeterStatus(data.ID, data.Status)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.StandardResponse{
			Message: fmt.Sprintf("error: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.StandardResponse{
		Message: "ok",
	})
}
