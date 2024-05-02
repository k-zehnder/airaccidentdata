// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/accidents": {
            "get": {
                "description": "Get a list of all aviation accidents",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accidents"
                ],
                "summary": "Get a list of accidents",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of accidents per page",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AccidentPaginatedResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid parameters",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/aircrafts": {
            "get": {
                "description": "Retrieve a list of all aircrafts.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Aircrafts"
                ],
                "summary": "Get a list of aircrafts",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of aircraft per page",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Aircrafts data with pagination details",
                        "schema": {
                            "$ref": "#/definitions/models.AircraftPaginatedResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid parameters",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Accident": {
            "type": "object",
            "properties": {
                "aircraft_damage_description": {
                    "type": "string"
                },
                "aircraft_id": {
                    "type": "integer"
                },
                "aircraft_missing_flag": {
                    "type": "string"
                },
                "entry_date": {
                    "type": "string"
                },
                "event_local_date": {
                    "type": "string"
                },
                "event_local_time": {
                    "type": "string"
                },
                "event_type_description": {
                    "type": "string"
                },
                "far_part": {
                    "type": "string"
                },
                "fatal_flag": {
                    "type": "string"
                },
                "flight_activity": {
                    "type": "string"
                },
                "flight_number": {
                    "type": "string"
                },
                "flight_phase": {
                    "type": "string"
                },
                "fsdo_description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "location_id": {
                    "type": "integer"
                },
                "remark_text": {
                    "type": "string"
                },
                "updated": {
                    "type": "string"
                }
            }
        },
        "models.AccidentPaginatedResponse": {
            "type": "object",
            "properties": {
                "accidents": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Accident"
                    }
                },
                "limit": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "models.Aircraft": {
            "type": "object",
            "properties": {
                "aircraft_make_name": {
                    "type": "string"
                },
                "aircraft_model_name": {
                    "type": "string"
                },
                "aircraft_operator": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "registration_number": {
                    "type": "string"
                }
            }
        },
        "models.AircraftPaginatedResponse": {
            "type": "object",
            "properties": {
                "aircrafts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Aircraft"
                    }
                },
                "limit": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "models.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "AirAccidentData API",
	Description:      "API server for airaccidentdata.com",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
