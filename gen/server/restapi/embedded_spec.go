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
  "info": {
    "description": "Debill API",
    "title": "debill-api",
    "contact": {
      "name": "Norbert Bondarczuk",
      "email": "nbondarczuk@yahoo.com"
    },
    "version": "1.0.0"
  },
  "paths": {
    "/api/account": {
      "get": {
        "tags": [
          "account"
        ],
        "summary": "Get account",
        "operationId": "getAccount",
        "parameters": [
          {
            "maximum": 255,
            "type": "string",
            "description": "Unique ID of an Account",
            "name": "Account",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/definitions/AccountResponse"
          }
        }
      }
    },
    "/api/address": {
      "get": {
        "tags": [
          "address"
        ],
        "summary": "Get address",
        "operationId": "getAddress",
        "parameters": [
          {
            "maximum": 255,
            "type": "string",
            "description": "Unique ID of an Address",
            "name": "Address",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/definitions/AddressResponse"
          }
        }
      }
    },
    "/api/customer": {
      "get": {
        "tags": [
          "customer"
        ],
        "summary": "Get customer",
        "operationId": "getCustomer",
        "parameters": [
          {
            "maximum": 255,
            "type": "string",
            "description": "Unique ID of a Customer",
            "name": "Customer",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/definitions/CustomerResponse"
          }
        }
      }
    },
    "/api/location": {
      "get": {
        "tags": [
          "location"
        ],
        "summary": "Get location",
        "operationId": "getLocation",
        "parameters": [
          {
            "maximum": 255,
            "type": "string",
            "description": "Unique ID of a Location",
            "name": "Location",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/definitions/LocationResponse"
          }
        }
      }
    }
  },
  "definitions": {
    "AccountResponse": {
      "type": "object",
      "required": [
        "account"
      ],
      "properties": {
        "location": {
          "description": "Unique ID of an Account",
          "type": "string",
          "format": "uuid",
          "example": "e760efc2-9407-4e34-adb2-d7cbc9f94ce2"
        },
        "name": {
          "description": "Descriptive name of an Account",
          "type": "string"
        }
      }
    },
    "AddressResponse": {
      "type": "object",
      "required": [
        "address"
      ],
      "properties": {
        "location": {
          "description": "Unique ID of an Address",
          "type": "string",
          "format": "uuid",
          "example": "e760efc2-9407-4e34-adb2-d7cbc9f94ce2"
        },
        "name": {
          "description": "Descriptive name of an Address",
          "type": "string"
        }
      }
    },
    "CustomerResponse": {
      "type": "object",
      "required": [
        "customer"
      ],
      "properties": {
        "location": {
          "description": "Unique ID of a Customer",
          "type": "string",
          "format": "uuid",
          "example": "e760efc2-9407-4e34-adb2-d7cbc9f94ce2"
        },
        "name": {
          "description": "Descriptive name of a Customer",
          "type": "string"
        }
      }
    },
    "LocationResponse": {
      "type": "object",
      "required": [
        "location"
      ],
      "properties": {
        "location": {
          "description": "Unique ID of a Location",
          "type": "string",
          "format": "uuid",
          "example": "e760efc2-9407-4e34-adb2-d7cbc9f94ce2"
        },
        "name": {
          "description": "Descriptive name of a Location",
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "API account endpoint",
      "name": "account"
    },
    {
      "description": "API address endpoint",
      "name": "address"
    },
    {
      "description": "API customer endpoint",
      "name": "customer"
    },
    {
      "description": "API location endpoint",
      "name": "location"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "info": {
    "description": "Debill API",
    "title": "debill-api",
    "contact": {
      "name": "Norbert Bondarczuk",
      "email": "nbondarczuk@yahoo.com"
    },
    "version": "1.0.0"
  },
  "paths": {
    "/api/account": {
      "get": {
        "tags": [
          "account"
        ],
        "summary": "Get account",
        "operationId": "getAccount",
        "parameters": [
          {
            "maximum": 255,
            "type": "string",
            "description": "Unique ID of an Account",
            "name": "Account",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        }
      }
    },
    "/api/address": {
      "get": {
        "tags": [
          "address"
        ],
        "summary": "Get address",
        "operationId": "getAddress",
        "parameters": [
          {
            "maximum": 255,
            "type": "string",
            "description": "Unique ID of an Address",
            "name": "Address",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        }
      }
    },
    "/api/customer": {
      "get": {
        "tags": [
          "customer"
        ],
        "summary": "Get customer",
        "operationId": "getCustomer",
        "parameters": [
          {
            "maximum": 255,
            "type": "string",
            "description": "Unique ID of a Customer",
            "name": "Customer",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        }
      }
    },
    "/api/location": {
      "get": {
        "tags": [
          "location"
        ],
        "summary": "Get location",
        "operationId": "getLocation",
        "parameters": [
          {
            "maximum": 255,
            "type": "string",
            "description": "Unique ID of a Location",
            "name": "Location",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        }
      }
    }
  },
  "definitions": {
    "AccountResponse": {
      "type": "object",
      "required": [
        "account"
      ],
      "properties": {
        "location": {
          "description": "Unique ID of an Account",
          "type": "string",
          "format": "uuid",
          "example": "e760efc2-9407-4e34-adb2-d7cbc9f94ce2"
        },
        "name": {
          "description": "Descriptive name of an Account",
          "type": "string"
        }
      }
    },
    "AddressResponse": {
      "type": "object",
      "required": [
        "address"
      ],
      "properties": {
        "location": {
          "description": "Unique ID of an Address",
          "type": "string",
          "format": "uuid",
          "example": "e760efc2-9407-4e34-adb2-d7cbc9f94ce2"
        },
        "name": {
          "description": "Descriptive name of an Address",
          "type": "string"
        }
      }
    },
    "CustomerResponse": {
      "type": "object",
      "required": [
        "customer"
      ],
      "properties": {
        "location": {
          "description": "Unique ID of a Customer",
          "type": "string",
          "format": "uuid",
          "example": "e760efc2-9407-4e34-adb2-d7cbc9f94ce2"
        },
        "name": {
          "description": "Descriptive name of a Customer",
          "type": "string"
        }
      }
    },
    "LocationResponse": {
      "type": "object",
      "required": [
        "location"
      ],
      "properties": {
        "location": {
          "description": "Unique ID of a Location",
          "type": "string",
          "format": "uuid",
          "example": "e760efc2-9407-4e34-adb2-d7cbc9f94ce2"
        },
        "name": {
          "description": "Descriptive name of a Location",
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "API account endpoint",
      "name": "account"
    },
    {
      "description": "API address endpoint",
      "name": "address"
    },
    {
      "description": "API customer endpoint",
      "name": "customer"
    },
    {
      "description": "API location endpoint",
      "name": "location"
    }
  ]
}`))
}
