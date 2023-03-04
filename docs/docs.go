// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/deleteemeter": {
            "post": {
                "description": "Elimina un medidor de energía en la base de datos",
                "summary": "Elimina medidor de energía",
                "parameters": [
                    {
                        "description": "Data",
                        "name": "args",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.DeleteEMeterModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.StandardResponse"
                        }
                    }
                }
            }
        },
        "/lastemeter/{brand}/{serial}": {
            "get": {
                "description": "Muestra el último medidor de energía instalado en la base de datos",
                "summary": "Último medidor de energía instalado",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Data",
                        "name": "brand",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Data",
                        "name": "serial",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Datameds"
                        }
                    }
                }
            }
        },
        "/listdisableemeter": {
            "get": {
                "description": "Muestra todos los registros de medidores de energía deshabilitados que existan en la base de datos",
                "summary": "Listado de todos los medidores de energía deshabilitados",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Datameds"
                            }
                        }
                    }
                }
            }
        },
        "/listemeter": {
            "get": {
                "description": "Muestra todos los registros de medidores de energía que existan en la base de datos",
                "summary": "Listado de todos los medidores de energía",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Datameds"
                            }
                        }
                    }
                }
            }
        },
        "/newemeter": {
            "post": {
                "description": "Crea un nuevo medidor de energía en la base de datos",
                "summary": "Creador de medidor de energía",
                "parameters": [
                    {
                        "description": "Data",
                        "name": "args",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.NewEMeterModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.StandardResponse"
                        }
                    }
                }
            }
        },
        "/updateemeterstatus": {
            "post": {
                "description": "Actualiza el estado de un medidor de energía en la base de datos",
                "summary": "Actualiza medidor de energía",
                "parameters": [
                    {
                        "description": "Update",
                        "name": "counter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateEMeterStatusModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.StandardResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Datameds": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "brand": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "installationDate": {
                    "type": "string"
                },
                "isActive": {
                    "type": "boolean"
                },
                "lines": {
                    "type": "integer"
                },
                "retirementDate": {
                    "type": "string"
                },
                "serial": {
                    "type": "string"
                }
            }
        },
        "models.DeleteEMeterModel": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "models.NewEMeterModel": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "brand": {
                    "type": "string"
                },
                "installationDate": {
                    "type": "string"
                },
                "lines": {
                    "type": "integer"
                },
                "serial": {
                    "type": "string"
                }
            }
        },
        "models.StandardResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.UpdateEMeterStatusModel": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Prueba técnica Enerbit",
	Description:      "Servicio de gestion contadores",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
