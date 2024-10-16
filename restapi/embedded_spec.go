// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API definition of the code challenge proposed.\nUseful links - [GitHub project](https://github.com/AMAndres/iskaypetchallenge) - [Swagger PetStore editor](https://editor.swagger.io/?url=https://petstore.swagger.io/v2/swagger.yaml)",
    "title": "IskayPet Code Challenge",
    "termsOfService": "--",
    "contact": {
      "email": "arenas.macineira.andres@gmail.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "github.com",
  "basePath": "/AMAndres/iskaypetchallenge/docs",
  "paths": {
    "/pet": {
      "put": {
        "security": [
          {
            "petstore_auth": [
              "write:pets",
              "read:pets"
            ]
          }
        ],
        "consumes": [
          "application/json",
          "application/xml"
        ],
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Update an existing pet",
        "operationId": "updatePet",
        "parameters": [
          {
            "description": "Pet object that needs to be added to the store",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Pet"
            }
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Pet not found"
          },
          "405": {
            "description": "Validation exception"
          }
        }
      },
      "post": {
        "security": [
          {
            "petstore_auth": [
              "write:pets",
              "read:pets"
            ]
          }
        ],
        "consumes": [
          "application/json",
          "application/xml"
        ],
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Add a new pet to the store",
        "operationId": "addPet",
        "parameters": [
          {
            "description": "Pet object that needs to be added to the store",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Pet"
            }
          }
        ],
        "responses": {
          "405": {
            "description": "Invalid input"
          }
        }
      }
    },
    "/pet/findByStatus": {
      "get": {
        "security": [
          {
            "petstore_auth": [
              "write:pets",
              "read:pets"
            ]
          }
        ],
        "description": "Multiple status values can be provided with comma separated strings",
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Finds Pets by status",
        "operationId": "findPetsByStatus",
        "parameters": [
          {
            "type": "array",
            "items": {
              "enum": [
                "available",
                "pending",
                "sold"
              ],
              "type": "string",
              "default": "available"
            },
            "collectionFormat": "multi",
            "description": "Status values that need to be considered for filter",
            "name": "status",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Pet"
              }
            }
          },
          "400": {
            "description": "Invalid status value"
          }
        }
      }
    },
    "/pet/findByTags": {
      "get": {
        "security": [
          {
            "petstore_auth": [
              "write:pets",
              "read:pets"
            ]
          }
        ],
        "description": "Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.",
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Finds Pets by tags",
        "operationId": "findPetsByTags",
        "deprecated": true,
        "parameters": [
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi",
            "description": "Tags to filter by",
            "name": "tags",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Pet"
              }
            }
          },
          "400": {
            "description": "Invalid tag value"
          }
        }
      }
    },
    "/pet/{petId}": {
      "get": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "Returns a single pet",
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Find pet by ID",
        "operationId": "getPetById",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of pet to return",
            "name": "petId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Pet"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Pet not found"
          }
        }
      },
      "post": {
        "security": [
          {
            "petstore_auth": [
              "write:pets",
              "read:pets"
            ]
          }
        ],
        "consumes": [
          "application/x-www-form-urlencoded"
        ],
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Updates a pet in the store with form data",
        "operationId": "updatePetWithForm",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of pet that needs to be updated",
            "name": "petId",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "Updated name of the pet",
            "name": "name",
            "in": "formData"
          },
          {
            "type": "string",
            "description": "Updated status of the pet",
            "name": "status",
            "in": "formData"
          }
        ],
        "responses": {
          "405": {
            "description": "Invalid input"
          }
        }
      },
      "delete": {
        "security": [
          {
            "petstore_auth": [
              "write:pets",
              "read:pets"
            ]
          }
        ],
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Deletes a pet",
        "operationId": "deletePet",
        "parameters": [
          {
            "type": "string",
            "name": "api_key",
            "in": "header"
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "Pet id to delete",
            "name": "petId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Pet not found"
          }
        }
      }
    }
  },
  "definitions": {
    "ApiResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      }
    },
    "Category": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string"
        }
      },
      "xml": {
        "name": "Category"
      }
    },
    "Order": {
      "type": "object",
      "properties": {
        "complete": {
          "type": "boolean"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "petId": {
          "type": "integer",
          "format": "int64"
        },
        "quantity": {
          "type": "integer",
          "format": "int32"
        },
        "shipDate": {
          "type": "string",
          "format": "date-time"
        },
        "status": {
          "description": "Order Status",
          "type": "string",
          "enum": [
            "placed",
            "approved",
            "delivered"
          ]
        }
      },
      "xml": {
        "name": "Order"
      }
    },
    "Pet": {
      "type": "object",
      "required": [
        "name",
        "photoUrls"
      ],
      "properties": {
        "category": {
          "$ref": "#/definitions/Category"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string",
          "example": "doggie"
        },
        "photoUrls": {
          "type": "array",
          "items": {
            "type": "string",
            "xml": {
              "name": "photoUrl"
            }
          },
          "xml": {
            "wrapped": true
          }
        },
        "status": {
          "description": "pet status in the store",
          "type": "string",
          "enum": [
            "available",
            "pending",
            "sold"
          ]
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Tag",
            "xml": {
              "name": "tag"
            }
          },
          "xml": {
            "wrapped": true
          }
        }
      },
      "xml": {
        "name": "Pet"
      }
    },
    "Tag": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string"
        }
      },
      "xml": {
        "name": "Tag"
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "firstName": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "lastName": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "phone": {
          "type": "string"
        },
        "userStatus": {
          "description": "User Status",
          "type": "integer",
          "format": "int32"
        },
        "username": {
          "type": "string"
        }
      },
      "xml": {
        "name": "User"
      }
    }
  },
  "securityDefinitions": {
    "api_key": {
      "type": "apiKey",
      "name": "api_key",
      "in": "header"
    },
    "petstore_auth": {
      "type": "oauth2",
      "flow": "implicit",
      "authorizationUrl": "https://petstore.swagger.io/oauth/authorize",
      "scopes": {
        "read:pets": "read your pets",
        "write:pets": "modify pets in your account"
      }
    }
  },
  "tags": [
    {
      "description": "Everything about Pet entity",
      "name": "pet"
    },
    {
      "description": "Everything about Species entity",
      "name": "species"
    }
  ],
  "externalDocs": {
    "description": "Find out more about Swagger",
    "url": "http://swagger.io"
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API definition of the code challenge proposed.\nUseful links - [GitHub project](https://github.com/AMAndres/iskaypetchallenge) - [Swagger PetStore editor](https://editor.swagger.io/?url=https://petstore.swagger.io/v2/swagger.yaml)",
    "title": "IskayPet Code Challenge",
    "termsOfService": "--",
    "contact": {
      "email": "arenas.macineira.andres@gmail.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "github.com",
  "basePath": "/AMAndres/iskaypetchallenge/docs",
  "paths": {
    "/pet": {
      "put": {
        "security": [
          {
            "petstore_auth": [
              "read:pets",
              "write:pets"
            ]
          }
        ],
        "consumes": [
          "application/json",
          "application/xml"
        ],
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Update an existing pet",
        "operationId": "updatePet",
        "parameters": [
          {
            "description": "Pet object that needs to be added to the store",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Pet"
            }
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Pet not found"
          },
          "405": {
            "description": "Validation exception"
          }
        }
      },
      "post": {
        "security": [
          {
            "petstore_auth": [
              "read:pets",
              "write:pets"
            ]
          }
        ],
        "consumes": [
          "application/json",
          "application/xml"
        ],
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Add a new pet to the store",
        "operationId": "addPet",
        "parameters": [
          {
            "description": "Pet object that needs to be added to the store",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Pet"
            }
          }
        ],
        "responses": {
          "405": {
            "description": "Invalid input"
          }
        }
      }
    },
    "/pet/findByStatus": {
      "get": {
        "security": [
          {
            "petstore_auth": [
              "read:pets",
              "write:pets"
            ]
          }
        ],
        "description": "Multiple status values can be provided with comma separated strings",
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Finds Pets by status",
        "operationId": "findPetsByStatus",
        "parameters": [
          {
            "type": "array",
            "items": {
              "enum": [
                "available",
                "pending",
                "sold"
              ],
              "type": "string",
              "default": "available"
            },
            "collectionFormat": "multi",
            "description": "Status values that need to be considered for filter",
            "name": "status",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Pet"
              }
            }
          },
          "400": {
            "description": "Invalid status value"
          }
        }
      }
    },
    "/pet/findByTags": {
      "get": {
        "security": [
          {
            "petstore_auth": [
              "read:pets",
              "write:pets"
            ]
          }
        ],
        "description": "Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.",
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Finds Pets by tags",
        "operationId": "findPetsByTags",
        "deprecated": true,
        "parameters": [
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi",
            "description": "Tags to filter by",
            "name": "tags",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Pet"
              }
            }
          },
          "400": {
            "description": "Invalid tag value"
          }
        }
      }
    },
    "/pet/{petId}": {
      "get": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "Returns a single pet",
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Find pet by ID",
        "operationId": "getPetById",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of pet to return",
            "name": "petId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Pet"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Pet not found"
          }
        }
      },
      "post": {
        "security": [
          {
            "petstore_auth": [
              "read:pets",
              "write:pets"
            ]
          }
        ],
        "consumes": [
          "application/x-www-form-urlencoded"
        ],
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Updates a pet in the store with form data",
        "operationId": "updatePetWithForm",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of pet that needs to be updated",
            "name": "petId",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "Updated name of the pet",
            "name": "name",
            "in": "formData"
          },
          {
            "type": "string",
            "description": "Updated status of the pet",
            "name": "status",
            "in": "formData"
          }
        ],
        "responses": {
          "405": {
            "description": "Invalid input"
          }
        }
      },
      "delete": {
        "security": [
          {
            "petstore_auth": [
              "read:pets",
              "write:pets"
            ]
          }
        ],
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Deletes a pet",
        "operationId": "deletePet",
        "parameters": [
          {
            "type": "string",
            "name": "api_key",
            "in": "header"
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "Pet id to delete",
            "name": "petId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Pet not found"
          }
        }
      }
    }
  },
  "definitions": {
    "ApiResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      }
    },
    "Category": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string"
        }
      },
      "xml": {
        "name": "Category"
      }
    },
    "Order": {
      "type": "object",
      "properties": {
        "complete": {
          "type": "boolean"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "petId": {
          "type": "integer",
          "format": "int64"
        },
        "quantity": {
          "type": "integer",
          "format": "int32"
        },
        "shipDate": {
          "type": "string",
          "format": "date-time"
        },
        "status": {
          "description": "Order Status",
          "type": "string",
          "enum": [
            "placed",
            "approved",
            "delivered"
          ]
        }
      },
      "xml": {
        "name": "Order"
      }
    },
    "Pet": {
      "type": "object",
      "required": [
        "name",
        "photoUrls"
      ],
      "properties": {
        "category": {
          "$ref": "#/definitions/Category"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string",
          "example": "doggie"
        },
        "photoUrls": {
          "type": "array",
          "items": {
            "type": "string",
            "xml": {
              "name": "photoUrl"
            }
          },
          "xml": {
            "wrapped": true
          }
        },
        "status": {
          "description": "pet status in the store",
          "type": "string",
          "enum": [
            "available",
            "pending",
            "sold"
          ]
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Tag",
            "xml": {
              "name": "tag"
            }
          },
          "xml": {
            "wrapped": true
          }
        }
      },
      "xml": {
        "name": "Pet"
      }
    },
    "Tag": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string"
        }
      },
      "xml": {
        "name": "Tag"
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "firstName": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "lastName": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "phone": {
          "type": "string"
        },
        "userStatus": {
          "description": "User Status",
          "type": "integer",
          "format": "int32"
        },
        "username": {
          "type": "string"
        }
      },
      "xml": {
        "name": "User"
      }
    }
  },
  "securityDefinitions": {
    "api_key": {
      "type": "apiKey",
      "name": "api_key",
      "in": "header"
    },
    "petstore_auth": {
      "type": "oauth2",
      "flow": "implicit",
      "authorizationUrl": "https://petstore.swagger.io/oauth/authorize",
      "scopes": {
        "read:pets": "read your pets",
        "write:pets": "modify pets in your account"
      }
    }
  },
  "tags": [
    {
      "description": "Everything about Pet entity",
      "name": "pet"
    },
    {
      "description": "Everything about Species entity",
      "name": "species"
    }
  ],
  "externalDocs": {
    "description": "Find out more about Swagger",
    "url": "http://swagger.io"
  }
}`))
}
