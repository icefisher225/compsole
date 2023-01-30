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
        "/auth/local/login": {
            "post": {
                "description": "Login with a local account",
                "consumes": [
                    "application/json",
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login with a local account",
                "parameters": [
                    {
                        "description": "User account details",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.UserLoginVals"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.User"
                        },
                        "headers": {
                            "Cookie": {
                                "type": "string",
                                "description": "` + "`" + `auth-cookie` + "`" + ` contains the session token"
                            }
                        }
                    }
                }
            }
        },
        "/auth/service/login": {
            "post": {
                "description": "Login with a service account",
                "consumes": [
                    "application/json",
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login with a service account",
                "parameters": [
                    {
                        "description": "Service account details",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.ServiceLoginVals"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.ServiceLoginResult"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/auth.ServiceLoginError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "action.Type": {
            "type": "string",
            "enum": [
                "SIGN_IN",
                "FAILED_SIGN_IN",
                "SIGN_OUT",
                "API_CALL",
                "CONSOLE_ACCESS",
                "REBOOT",
                "SHUTDOWN",
                "POWER_ON",
                "POWER_OFF",
                "CHANGE_SELF_PASSWORD",
                "CHANGE_PASSWORD",
                "CREATE_OBJECT",
                "UPDATE_OBJECT",
                "DELETE_OBJECT",
                "UPDATE_LOCKOUT"
            ],
            "x-enum-varnames": [
                "TypeSIGN_IN",
                "TypeFAILED_SIGN_IN",
                "TypeSIGN_OUT",
                "TypeAPI_CALL",
                "TypeCONSOLE_ACCESS",
                "TypeREBOOT",
                "TypeSHUTDOWN",
                "TypePOWER_ON",
                "TypePOWER_OFF",
                "TypeCHANGE_SELF_PASSWORD",
                "TypeCHANGE_PASSWORD",
                "TypeCREATE_OBJECT",
                "TypeUPDATE_OBJECT",
                "TypeDELETE_OBJECT",
                "TypeUPDATE_LOCKOUT"
            ]
        },
        "auth.ServiceLoginError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "auth.ServiceLoginResult": {
            "type": "object",
            "properties": {
                "expires_at": {
                    "type": "integer"
                },
                "refresh_token": {
                    "type": "string"
                },
                "session_token": {
                    "type": "string"
                }
            }
        },
        "auth.ServiceLoginVals": {
            "type": "object",
            "required": [
                "api_key",
                "api_secret"
            ],
            "properties": {
                "api_key": {
                    "type": "string"
                },
                "api_secret": {
                    "type": "string"
                }
            }
        },
        "auth.UserLoginVals": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "example": "password123"
                },
                "username": {
                    "type": "string",
                    "example": "admin"
                }
            }
        },
        "ent.Action": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the ActionQuery when eager-loading is set.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.ActionEdges"
                        }
                    ]
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "ip_address": {
                    "description": "IPAddress holds the value of the \"ip_address\" field.",
                    "type": "string"
                },
                "message": {
                    "description": "Message holds the value of the \"message\" field.",
                    "type": "string"
                },
                "performed_at": {
                    "description": "PerformedAt holds the value of the \"performed_at\" field.",
                    "type": "string"
                },
                "type": {
                    "description": "Type holds the value of the \"type\" field.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/action.Type"
                        }
                    ]
                }
            }
        },
        "ent.ActionEdges": {
            "type": "object",
            "properties": {
                "ActionToServiceAccount": {
                    "description": "ActionToServiceAccount holds the value of the ActionToServiceAccount edge.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.ServiceAccount"
                        }
                    ]
                },
                "ActionToUser": {
                    "description": "ActionToUser holds the value of the ActionToUser edge.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.User"
                        }
                    ]
                }
            }
        },
        "ent.Competition": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the CompetitionQuery when eager-loading is set.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.CompetitionEdges"
                        }
                    ]
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.\n[REQUIRED] The unique name (aka. slug) for the competition.",
                    "type": "string"
                }
            }
        },
        "ent.CompetitionEdges": {
            "type": "object",
            "properties": {
                "CompetitionToProvider": {
                    "description": "CompetitionToProvider holds the value of the CompetitionToProvider edge.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.Provider"
                        }
                    ]
                },
                "CompetitionToTeams": {
                    "description": "CompetitionToTeams holds the value of the CompetitionToTeams edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Team"
                    }
                }
            }
        },
        "ent.Provider": {
            "type": "object",
            "properties": {
                "config": {
                    "description": "Config holds the value of the \"config\" field.\n[REQUIRED] This is the JSON configuration for the provider.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the ProviderQuery when eager-loading is set.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.ProviderEdges"
                        }
                    ]
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.\n[REQUIRED] The unique name (aka. slug) for the provider.",
                    "type": "string"
                },
                "type": {
                    "description": "Type holds the value of the \"type\" field.\n[REQUIRED] The unique name (aka. slug) for the provider.",
                    "type": "string"
                }
            }
        },
        "ent.ProviderEdges": {
            "type": "object",
            "properties": {
                "ProviderToCompetition": {
                    "description": "ProviderToCompetition holds the value of the ProviderToCompetition edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Competition"
                    }
                }
            }
        },
        "ent.ServiceAccount": {
            "type": "object",
            "properties": {
                "active": {
                    "description": "Active holds the value of the \"active\" field.\nDetermines whether or not the service account is active or not",
                    "allOf": [
                        {
                            "$ref": "#/definitions/serviceaccount.Active"
                        }
                    ]
                },
                "api_key": {
                    "description": "APIKey holds the value of the \"api_key\" field.\n[REQUIRED] The API key for the service account. Equivalent to a username.",
                    "type": "string"
                },
                "api_secret": {
                    "description": "APISecret holds the value of the \"api_secret\" field.\n[REQUIRED] The API secret for the service account. This value MUST be protected.",
                    "type": "string"
                },
                "display_name": {
                    "description": "DisplayName holds the value of the \"display_name\" field.\n[REQUIRED] The display/common name for the service account.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the ServiceAccountQuery when eager-loading is set.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.ServiceAccountEdges"
                        }
                    ]
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                }
            }
        },
        "ent.ServiceAccountEdges": {
            "type": "object",
            "properties": {
                "ServiceAccountToActions": {
                    "description": "ServiceAccountToActions holds the value of the ServiceAccountToActions edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Action"
                    }
                },
                "ServiceAccountToToken": {
                    "description": "ServiceAccountToToken holds the value of the ServiceAccountToToken edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.ServiceToken"
                    }
                }
            }
        },
        "ent.ServiceToken": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the ServiceTokenQuery when eager-loading is set.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.ServiceTokenEdges"
                        }
                    ]
                },
                "expire_at": {
                    "description": "ExpireAt holds the value of the \"expire_at\" field.\n[REQUIRED] The time the token should expire.",
                    "type": "integer"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "refresh_token": {
                    "description": "RefreshToken holds the value of the \"refresh_token\" field.\n[REQUIRED] The refresh token used to renew an expired service account session. These are only valid for 1 hour after the associated token expires.",
                    "type": "string"
                },
                "token": {
                    "description": "Token holds the value of the \"token\" field.\n[REQUIRED] The API token for a service account session.",
                    "type": "string"
                }
            }
        },
        "ent.ServiceTokenEdges": {
            "type": "object",
            "properties": {
                "TokenToServiceAccount": {
                    "description": "TokenToServiceAccount holds the value of the TokenToServiceAccount edge.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.ServiceAccount"
                        }
                    ]
                }
            }
        },
        "ent.Team": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the TeamQuery when eager-loading is set.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.TeamEdges"
                        }
                    ]
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.\n[OPTIONAL] The display name for the team.",
                    "type": "string"
                },
                "team_number": {
                    "description": "TeamNumber holds the value of the \"team_number\" field.\n[REQUIRED] The team number.",
                    "type": "integer"
                }
            }
        },
        "ent.TeamEdges": {
            "type": "object",
            "properties": {
                "TeamToCompetition": {
                    "description": "TeamToCompetition holds the value of the TeamToCompetition edge.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.Competition"
                        }
                    ]
                },
                "TeamToUsers": {
                    "description": "TeamToUsers holds the value of the TeamToUsers edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.User"
                    }
                },
                "TeamToVmObjects": {
                    "description": "TeamToVmObjects holds the value of the TeamToVmObjects edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.VmObject"
                    }
                }
            }
        },
        "ent.Token": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the TokenQuery when eager-loading is set.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.TokenEdges"
                        }
                    ]
                },
                "expire_at": {
                    "description": "ExpireAt holds the value of the \"expire_at\" field.\n[REQUIRED] The time the token should expire.",
                    "type": "integer"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "token": {
                    "description": "Token holds the value of the \"token\" field.\n[REQUIRED] The auth-token cookie value for the user session.",
                    "type": "string"
                }
            }
        },
        "ent.TokenEdges": {
            "type": "object",
            "properties": {
                "TokenToUser": {
                    "description": "TokenToUser holds the value of the TokenToUser edge.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.User"
                        }
                    ]
                }
            }
        },
        "ent.User": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the UserQuery when eager-loading is set.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.UserEdges"
                        }
                    ]
                },
                "first_name": {
                    "description": "FirstName holds the value of the \"first_name\" field.\n[OPTIONAL] The display first name for the user.",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "last_name": {
                    "description": "LastName holds the value of the \"last_name\" field.\n[OPTIONAL] The display last name for the user",
                    "type": "string"
                },
                "provider": {
                    "description": "Provider holds the value of the \"provider\" field.\n[REQUIRED] The type of login the user will be using.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/user.Provider"
                        }
                    ]
                },
                "role": {
                    "description": "Role holds the value of the \"role\" field.\n[REQUIRED] The role of the user. Admins have full access.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/user.Role"
                        }
                    ]
                },
                "username": {
                    "description": "Username holds the value of the \"username\" field.\n[REQUIRED] The username for the user.",
                    "type": "string"
                }
            }
        },
        "ent.UserEdges": {
            "type": "object",
            "properties": {
                "UserToActions": {
                    "description": "UserToActions holds the value of the UserToActions edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Action"
                    }
                },
                "UserToTeam": {
                    "description": "UserToTeam holds the value of the UserToTeam edge.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.Team"
                        }
                    ]
                },
                "UserToToken": {
                    "description": "UserToToken holds the value of the UserToToken edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Token"
                    }
                }
            }
        },
        "ent.VmObject": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the VmObjectQuery when eager-loading is set.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.VmObjectEdges"
                        }
                    ]
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "identifier": {
                    "description": "Identifier holds the value of the \"identifier\" field.\n[REQUIRED] The identifier of the VM. This will be provider-specific.",
                    "type": "string"
                },
                "ip_addresses": {
                    "description": "IPAddresses holds the value of the \"ip_addresses\" field.\n[OPTIONAL] IP addresses of the VM. This will be displayed to the user.",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "locked": {
                    "description": "Locked holds the value of the \"locked\" field.\n[REQUIRED] (default is false) If a vm is locked, standard users will not be able to access this VM.",
                    "type": "boolean"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.\n[REQUIRED] A user-friendly name for the VM. This will be provider-specific.",
                    "type": "string"
                }
            }
        },
        "ent.VmObjectEdges": {
            "type": "object",
            "properties": {
                "VmObjectToTeam": {
                    "description": "VmObjectToTeam holds the value of the VmObjectToTeam edge.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.Team"
                        }
                    ]
                }
            }
        },
        "serviceaccount.Active": {
            "type": "string",
            "enum": [
                "enabled",
                "disabled"
            ],
            "x-enum-varnames": [
                "ActiveEnabled",
                "ActiveDisabled"
            ]
        },
        "user.Provider": {
            "type": "string",
            "enum": [
                "LOCAL",
                "GITLAB"
            ],
            "x-enum-varnames": [
                "ProviderLOCAL",
                "ProviderGITLAB"
            ]
        },
        "user.Role": {
            "type": "string",
            "enum": [
                "USER",
                "ADMIN"
            ],
            "x-enum-varnames": [
                "RoleUSER",
                "RoleADMIN"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
