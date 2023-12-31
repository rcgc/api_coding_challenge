// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Roberto Guzmán",
            "email": "roberto140298@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "https://opensource.org/license/mit/"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/cars": {
            "get": {
                "description": "Gets all the cars from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "car"
                ],
                "summary": "Get all cars",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Car"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Updates an existing car from the database corresponding to the id sent. Otherwise, returns error",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "car"
                ],
                "summary": "Update a car",
                "parameters": [
                    {
                        "description": "Car JSON Object",
                        "name": "car",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Car"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Car"
                        }
                    },
                    "400": {
                        "description": "BadRequest",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new car in the database. In case of existing id returns error",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "car"
                ],
                "summary": "Create a new car",
                "parameters": [
                    {
                        "description": "Car JSON Object",
                        "name": "car",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Car"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Car"
                        }
                    },
                    "400": {
                        "description": "BadRequest",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/cars/{id}": {
            "get": {
                "description": "Gets a single car from the database corresponding to the id in the path. Otherwise, returns error",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "car"
                ],
                "summary": "Get a car",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Car Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Car"
                        }
                    },
                    "404": {
                        "description": "NotFound",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes an existing car from the database corresponding to the id in the path. Otherwise, returns error",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "car"
                ],
                "summary": "Delete a car",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Car Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "NoContent",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "NotFound",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Car": {
            "description": "car information",
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "color": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "make": {
                    "type": "string"
                },
                "mileage": {
                    "type": "number"
                },
                "model": {
                    "type": "string"
                },
                "package": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "year": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Cars Restful API with Swagger",
	Description:      "Simple swagger implementation in Go HTTP",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
