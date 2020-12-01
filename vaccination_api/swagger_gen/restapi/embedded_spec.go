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
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Digital infra for vaccination certificates",
    "title": "Divoc",
    "version": "1.0.0"
  },
  "host": "divoc.xiv.in",
  "basePath": "/divoc/api/v1",
  "paths": {
    "/authorize": {
      "post": {
        "security": [],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "login"
        ],
        "summary": "Establish token",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/LoginRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/LoginResponse"
            }
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      }
    },
    "/certify": {
      "post": {
        "security": [
          {
            "hasRole": [
              "facility-staff"
            ]
          }
        ],
        "description": "Certification happens asynchronously, this requires vaccinator athorization and vaccinator should be trained for the vaccination that is being certified.",
        "tags": [
          "certification"
        ],
        "summary": "Certify the one or more vaccination",
        "operationId": "certify",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/CertificationRequest"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/divoc/configuration": {
      "get": {
        "security": [
          {
            "isAdmin": []
          }
        ],
        "tags": [
          "configuration"
        ],
        "summary": "Get Meta information about the application flow",
        "operationId": "getConfiguration",
        "parameters": [
          {
            "type": "string",
            "name": "lastKnownVersion",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/ApplicationConfiguration"
            }
          }
        }
      }
    },
    "/identity/verify": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "tags": [
          "identity"
        ],
        "summary": "Validate identity if the person",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/IdentityVerificationRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "206": {
            "description": "Need OTP"
          }
        }
      }
    },
    "/ping": {
      "get": {
        "security": [],
        "description": "This operation shows how to override the global security defined above, as we want to open it up for all users.",
        "summary": "Server heartbeat operation",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/preEnrollments/facility": {
      "get": {
        "security": [
          {
            "hasRole": [
              "facility-staff"
            ]
          }
        ],
        "tags": [
          "vaccination"
        ],
        "summary": "Get all pre enrollments applicable to assigned facility",
        "operationId": "getPreEnrollmentsForFacility",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/PreEnrollment"
              }
            }
          }
        }
      }
    },
    "/preEnrollments/{preEnrollmentCode}": {
      "get": {
        "security": [
          {
            "hasRole": [
              "facility-staff"
            ]
          }
        ],
        "description": "Get pre enrollment data from api for vaccination",
        "tags": [
          "vaccination"
        ],
        "summary": "Get pre enrollment information",
        "operationId": "getPreEnrollment",
        "parameters": [
          {
            "type": "string",
            "name": "preEnrollmentCode",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/PreEnrollment"
            }
          }
        }
      }
    },
    "/programs/current": {
      "get": {
        "security": [
          {
            "hasRole": [
              "facility-staff"
            ]
          }
        ],
        "tags": [
          "configuration"
        ],
        "summary": "Get active vaccination programs",
        "operationId": "getCurrentPrograms",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Program"
              }
            }
          }
        }
      }
    },
    "/users/me": {
      "get": {
        "tags": [
          "vaccination"
        ],
        "summary": "Get User information",
        "operationId": "getLoggedInUserInfo",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/UserInfo"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ApplicationConfiguration": {
      "type": "object",
      "properties": {
        "navigation": {
          "type": "object"
        },
        "styles": {
          "type": "object"
        },
        "validation": {
          "type": "object"
        }
      }
    },
    "CertificationRequest": {
      "type": "object",
      "properties": {
        "facility": {
          "type": "object",
          "properties": {
            "address": {
              "type": "string"
            },
            "name": {
              "type": "string"
            }
          }
        },
        "preEnrollmentCode": {
          "type": "string"
        },
        "recipient": {
          "type": "object",
          "properties": {
            "dob": {
              "type": "string",
              "format": "date"
            },
            "gender": {
              "type": "string"
            },
            "identity": {
              "type": "string"
            },
            "name": {
              "type": "string"
            },
            "natinoality": {
              "type": "string"
            }
          }
        },
        "vaccination": {
          "type": "object",
          "properties": {
            "batch": {
              "type": "string"
            },
            "date": {
              "type": "string",
              "format": "date-time"
            },
            "effectiveStart": {
              "type": "string",
              "format": "date"
            },
            "effectiveUntil": {
              "type": "string",
              "format": "date"
            },
            "manufacturer": {
              "type": "string"
            },
            "name": {
              "type": "string"
            }
          }
        },
        "vaccinator": {
          "type": "object",
          "properties": {
            "name": {
              "type": "string"
            }
          }
        }
      }
    },
    "IdentityVerificationRequest": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "token": {
          "type": "string"
        }
      }
    },
    "LoginRequest": {
      "type": "object",
      "properties": {
        "mobile": {
          "type": "string"
        },
        "token2fa": {
          "type": "string"
        }
      }
    },
    "LoginResponse": {
      "type": "object",
      "properties": {
        "refreshToken": {
          "type": "string"
        },
        "token": {
          "type": "string"
        }
      }
    },
    "PreEnrollment": {
      "type": "object",
      "properties": {
        "code": {
          "type": "string"
        },
        "dob": {
          "type": "string",
          "format": "date"
        },
        "email": {
          "type": "string"
        },
        "enrollmentScopeId": {
          "type": "string"
        },
        "gender": {
          "type": "string",
          "enum": [
            "Male",
            "Female",
            "Other"
          ]
        },
        "meta": {
          "type": "object"
        },
        "name": {
          "type": "string"
        },
        "nationalId": {
          "type": "string"
        },
        "phone": {
          "type": "string"
        }
      }
    },
    "Program": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "medicines": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "name": {
          "type": "string"
        }
      }
    },
    "UserInfo": {
      "type": "object",
      "properties": {
        "firstName": {
          "type": "string"
        },
        "lastName": {
          "type": "string"
        },
        "mobile": {
          "type": "string"
        },
        "roles": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    }
  },
  "securityDefinitions": {
    "hasRole": {
      "type": "oauth2",
      "flow": "accessCode",
      "authorizationUrl": "https://divoc.xiv.in/keycloak/auth/realms/divoc/protocol/openid-connect/auth",
      "tokenUrl": "https://divoc.xiv.in/keycloak/auth/realms/divoc/protocol/openid-connect/token",
      "scopes": {
        "admin": "scope of super admin",
        "facility-staff": "scope of facility staff",
        "facillity-admin": "scope of facility admin"
      }
    }
  },
  "security": [
    {
      "isUser": []
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Digital infra for vaccination certificates",
    "title": "Divoc",
    "version": "1.0.0"
  },
  "host": "divoc.xiv.in",
  "basePath": "/divoc/api/v1",
  "paths": {
    "/authorize": {
      "post": {
        "security": [],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "login"
        ],
        "summary": "Establish token",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/LoginRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/LoginResponse"
            }
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      }
    },
    "/certify": {
      "post": {
        "security": [
          {
            "hasRole": [
              "facility-staff"
            ]
          }
        ],
        "description": "Certification happens asynchronously, this requires vaccinator athorization and vaccinator should be trained for the vaccination that is being certified.",
        "tags": [
          "certification"
        ],
        "summary": "Certify the one or more vaccination",
        "operationId": "certify",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/CertificationRequest"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/divoc/configuration": {
      "get": {
        "security": [
          {
            "isAdmin": []
          }
        ],
        "tags": [
          "configuration"
        ],
        "summary": "Get Meta information about the application flow",
        "operationId": "getConfiguration",
        "parameters": [
          {
            "type": "string",
            "name": "lastKnownVersion",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/ApplicationConfiguration"
            }
          }
        }
      }
    },
    "/identity/verify": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "tags": [
          "identity"
        ],
        "summary": "Validate identity if the person",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/IdentityVerificationRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "206": {
            "description": "Need OTP"
          }
        }
      }
    },
    "/ping": {
      "get": {
        "security": [],
        "description": "This operation shows how to override the global security defined above, as we want to open it up for all users.",
        "summary": "Server heartbeat operation",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/preEnrollments/facility": {
      "get": {
        "security": [
          {
            "hasRole": [
              "facility-staff"
            ]
          }
        ],
        "tags": [
          "vaccination"
        ],
        "summary": "Get all pre enrollments applicable to assigned facility",
        "operationId": "getPreEnrollmentsForFacility",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/PreEnrollment"
              }
            }
          }
        }
      }
    },
    "/preEnrollments/{preEnrollmentCode}": {
      "get": {
        "security": [
          {
            "hasRole": [
              "facility-staff"
            ]
          }
        ],
        "description": "Get pre enrollment data from api for vaccination",
        "tags": [
          "vaccination"
        ],
        "summary": "Get pre enrollment information",
        "operationId": "getPreEnrollment",
        "parameters": [
          {
            "type": "string",
            "name": "preEnrollmentCode",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/PreEnrollment"
            }
          }
        }
      }
    },
    "/programs/current": {
      "get": {
        "security": [
          {
            "hasRole": [
              "facility-staff"
            ]
          }
        ],
        "tags": [
          "configuration"
        ],
        "summary": "Get active vaccination programs",
        "operationId": "getCurrentPrograms",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Program"
              }
            }
          }
        }
      }
    },
    "/users/me": {
      "get": {
        "tags": [
          "vaccination"
        ],
        "summary": "Get User information",
        "operationId": "getLoggedInUserInfo",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/UserInfo"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ApplicationConfiguration": {
      "type": "object",
      "properties": {
        "navigation": {
          "type": "object"
        },
        "styles": {
          "type": "object"
        },
        "validation": {
          "type": "object"
        }
      }
    },
    "CertificationRequest": {
      "type": "object",
      "properties": {
        "facility": {
          "type": "object",
          "properties": {
            "address": {
              "type": "string"
            },
            "name": {
              "type": "string"
            }
          }
        },
        "preEnrollmentCode": {
          "type": "string"
        },
        "recipient": {
          "type": "object",
          "properties": {
            "dob": {
              "type": "string",
              "format": "date"
            },
            "gender": {
              "type": "string"
            },
            "identity": {
              "type": "string"
            },
            "name": {
              "type": "string"
            },
            "natinoality": {
              "type": "string"
            }
          }
        },
        "vaccination": {
          "type": "object",
          "properties": {
            "batch": {
              "type": "string"
            },
            "date": {
              "type": "string",
              "format": "date-time"
            },
            "effectiveStart": {
              "type": "string",
              "format": "date"
            },
            "effectiveUntil": {
              "type": "string",
              "format": "date"
            },
            "manufacturer": {
              "type": "string"
            },
            "name": {
              "type": "string"
            }
          }
        },
        "vaccinator": {
          "type": "object",
          "properties": {
            "name": {
              "type": "string"
            }
          }
        }
      }
    },
    "CertificationRequestFacility": {
      "type": "object",
      "properties": {
        "address": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "CertificationRequestRecipient": {
      "type": "object",
      "properties": {
        "dob": {
          "type": "string",
          "format": "date"
        },
        "gender": {
          "type": "string"
        },
        "identity": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "natinoality": {
          "type": "string"
        }
      }
    },
    "CertificationRequestVaccination": {
      "type": "object",
      "properties": {
        "batch": {
          "type": "string"
        },
        "date": {
          "type": "string",
          "format": "date-time"
        },
        "effectiveStart": {
          "type": "string",
          "format": "date"
        },
        "effectiveUntil": {
          "type": "string",
          "format": "date"
        },
        "manufacturer": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "CertificationRequestVaccinator": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        }
      }
    },
    "IdentityVerificationRequest": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "token": {
          "type": "string"
        }
      }
    },
    "LoginRequest": {
      "type": "object",
      "properties": {
        "mobile": {
          "type": "string"
        },
        "token2fa": {
          "type": "string"
        }
      }
    },
    "LoginResponse": {
      "type": "object",
      "properties": {
        "refreshToken": {
          "type": "string"
        },
        "token": {
          "type": "string"
        }
      }
    },
    "PreEnrollment": {
      "type": "object",
      "properties": {
        "code": {
          "type": "string"
        },
        "dob": {
          "type": "string",
          "format": "date"
        },
        "email": {
          "type": "string"
        },
        "enrollmentScopeId": {
          "type": "string"
        },
        "gender": {
          "type": "string",
          "enum": [
            "Male",
            "Female",
            "Other"
          ]
        },
        "meta": {
          "type": "object"
        },
        "name": {
          "type": "string"
        },
        "nationalId": {
          "type": "string"
        },
        "phone": {
          "type": "string"
        }
      }
    },
    "Program": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "medicines": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "name": {
          "type": "string"
        }
      }
    },
    "UserInfo": {
      "type": "object",
      "properties": {
        "firstName": {
          "type": "string"
        },
        "lastName": {
          "type": "string"
        },
        "mobile": {
          "type": "string"
        },
        "roles": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    }
  },
  "securityDefinitions": {
    "hasRole": {
      "type": "oauth2",
      "flow": "accessCode",
      "authorizationUrl": "https://divoc.xiv.in/keycloak/auth/realms/divoc/protocol/openid-connect/auth",
      "tokenUrl": "https://divoc.xiv.in/keycloak/auth/realms/divoc/protocol/openid-connect/token",
      "scopes": {
        "admin": "scope of super admin",
        "facility-staff": "scope of facility staff",
        "facillity-admin": "scope of facility admin"
      }
    }
  },
  "security": [
    {
      "isUser": []
    }
  ]
}`))
}
