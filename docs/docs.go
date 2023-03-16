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
        "/auth/login": {
            "post": {
                "description": "Users can login to the application and obtain a JWT token through this endpoint",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Login",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AuthLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AuthLoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    }
                }
            }
        },
        "/auth/registration": {
            "post": {
                "description": "Registration endpoint for user new users to register for an account, after registering for an account, the user will be able to login to the system and obtain a JWT Token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Registration",
                "parameters": [
                    {
                        "description": "Registration",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AuthRegistrationRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AuthRegistrationResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    }
                }
            }
        },
        "/campaign": {
            "get": {
                "description": "Retrieve all campaigns, sorted by start date",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "campaign"
                ],
                "summary": "Retrieve all Campaigns",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer eyJhb...",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CampaignList"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    }
                }
            }
        },
        "/campaign/{campaign_id}": {
            "get": {
                "description": "Retrieve Campaign based on campaignId",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "campaign"
                ],
                "summary": "Retrieve Campaign based on campaignId",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer eyJhb...",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "campaign's id",
                        "name": "campaign_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Campaign"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    }
                }
            },
            "put": {
                "description": "Update Campaign based on campaignId",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "campaign"
                ],
                "summary": "Update Campaign based on campaignId",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer eyJhb...",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "campaign's id",
                        "name": "campaign_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "campaign",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CampaignList"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Campaign"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Campaign based on campaignId",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "campaign"
                ],
                "summary": "Delete Campaign based on campaignId",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer eyJhb...",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "campaign's id",
                        "name": "campaign_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    }
                }
            }
        },
        "/campaign/{user_id}": {
            "post": {
                "description": "Create campaigns",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "campaign"
                ],
                "summary": "Create Campaigns for Merchants",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer eyJhb...",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "user's id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "campaigns",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CampaignList"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Campaign"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Health Check Endpoint that doesn't require authentication",
                "tags": [
                    "health"
                ],
                "summary": "Health Check",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    }
                }
            }
        },
        "/transaction/{user_id}": {
            "get": {
                "description": "Retrieve transaction records of a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transaction"
                ],
                "summary": "Retrieve Transactions of User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer eyJhb...",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "user's id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "minimum": 0,
                        "type": "integer",
                        "default": 100,
                        "description": "maximum records per page",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "minimum": 0,
                        "type": "integer",
                        "default": 0,
                        "description": "page of records, starts from 0",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Transaction"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "Create transaction records",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transaction"
                ],
                "summary": "Create Transactions for User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer eyJhb...",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "user's id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "transactions",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TransactionList"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Transaction"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AuthLoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "sudo@soonann.dev"
                },
                "password": {
                    "type": "string",
                    "example": "supersecret"
                }
            }
        },
        "models.AuthLoginResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2023-03-02T13:10:23Z"
                },
                "email": {
                    "type": "string",
                    "example": "sudo@soonann.dev"
                },
                "first_name": {
                    "type": "string",
                    "example": "Soon Ann"
                },
                "last_name": {
                    "type": "string",
                    "example": "Tan"
                },
                "password": {
                    "type": "string",
                    "example": "hashedsupersecret"
                },
                "phone": {
                    "type": "string",
                    "example": "91234567"
                },
                "refresh_token": {
                    "type": "string",
                    "example": "eyJhb..."
                },
                "token": {
                    "type": "string",
                    "example": "eyJhb..."
                },
                "updated_at": {
                    "type": "string",
                    "example": "2023-03-02T13:10:23Z"
                },
                "user_id": {
                    "type": "string",
                    "example": "6400a..."
                }
            }
        },
        "models.AuthRegistrationRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "sudo@soonann.dev"
                },
                "first_name": {
                    "type": "string",
                    "example": "Soon Ann"
                },
                "last_name": {
                    "type": "string",
                    "example": "Tan"
                },
                "password": {
                    "type": "string",
                    "example": "supersecret"
                },
                "phone": {
                    "type": "string",
                    "example": "91234567"
                }
            }
        },
        "models.AuthRegistrationResponse": {
            "type": "object",
            "properties": {
                "InsertedID": {
                    "type": "string",
                    "example": "6400a..."
                }
            }
        },
        "models.Campaign": {
            "type": "object",
            "properties": {
                "campaign_id": {
                    "type": "string",
                    "example": "cmp00001"
                },
                "card_type": {
                    "type": "string",
                    "example": "super_miles_card"
                },
                "description": {
                    "type": "string"
                },
                "end_date": {
                    "type": "string",
                    "example": "2023-03-03T13:10:23Z"
                },
                "merchant": {
                    "type": "string",
                    "example": "7-11"
                },
                "start_date": {
                    "type": "string",
                    "example": "2023-03-02T13:10:23Z"
                }
            }
        },
        "models.CampaignList": {
            "type": "object",
            "properties": {
                "campaigns": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Campaign"
                    }
                }
            }
        },
        "models.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        },
        "models.Transaction": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number",
                    "example": 20.1
                },
                "card_id": {
                    "type": "string",
                    "example": "4111222233334444"
                },
                "card_pan": {
                    "type": "string",
                    "example": "xyz"
                },
                "card_type": {
                    "type": "string"
                },
                "currency": {
                    "type": "string",
                    "example": "USD"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "mcc": {
                    "type": "string",
                    "example": "5311"
                },
                "merchant": {
                    "type": "string",
                    "example": "7-11"
                },
                "transaction_date": {
                    "type": "string",
                    "example": "yyyy-mm-dd hh:mm:ss"
                },
                "transaction_id": {
                    "type": "string",
                    "example": "txn00001"
                }
            }
        },
        "models.TransactionList": {
            "type": "object",
            "properties": {
                "transactions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Transaction"
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "go-gin-backend",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
