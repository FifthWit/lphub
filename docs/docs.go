// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "GNU General Public License, Version 2",
            "url": "https://www.gnu.org/licenses/old-licenses/gpl-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/chapters/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "games \u0026 chapters"
                ],
                "summary": "Get maps from the specified chapter id.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Chapter ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.ChapterMapsResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/demo": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rankings"
                ],
                "summary": "Get rankings of every player.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.RankingsResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/demos": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/octet-stream"
                ],
                "tags": [
                    "demo"
                ],
                "summary": "Get demo with specified demo uuid.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Demo UUID",
                        "name": "uuid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Demo File",
                        "schema": {
                            "type": "file"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/games": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "games \u0026 chapters"
                ],
                "summary": "Get games from the leaderboards.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.Game"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/games/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "games \u0026 chapters"
                ],
                "summary": "Get chapters from the specified game id.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Game ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.ChaptersResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/login": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login"
                ],
                "summary": "Get (redirect) login page for Steam auth.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.LoginResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/maps/{id}/leaderboards": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "maps"
                ],
                "summary": "Get map leaderboards with specified id.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Map ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/models.Map"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "data": {
                                                            "$ref": "#/definitions/models.MapRecords"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/maps/{id}/record": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "maps"
                ],
                "summary": "Post record with demo of a specific map.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "JWT Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "file"
                        },
                        "description": "Demos",
                        "name": "demos",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Score Count",
                        "name": "score_count",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Score Time",
                        "name": "score_time",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "Is Partner Orange",
                        "name": "is_partner_orange",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Partner ID",
                        "name": "partner_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.RecordRequest"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/maps/{id}/summary": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "maps"
                ],
                "summary": "Get map summary with specified id.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Map ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/models.Map"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "data": {
                                                            "$ref": "#/definitions/models.MapSummary"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/profile": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get profile page of session user.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "JWT Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.ProfileResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update country code of session user.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "JWT Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Country Code [XX]",
                        "name": "country_code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.ProfileResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update profile page of session user.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "JWT Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.ProfileResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/search": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "search"
                ],
                "summary": "Get all user and map data.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.SearchResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get profile page of another user.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.ProfileResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Chapter": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.ChapterMapsResponse": {
            "type": "object",
            "properties": {
                "chapter": {
                    "$ref": "#/definitions/models.Chapter"
                },
                "maps": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.MapShort"
                    }
                }
            }
        },
        "models.ChaptersResponse": {
            "type": "object",
            "properties": {
                "chapters": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Chapter"
                    }
                },
                "game": {
                    "$ref": "#/definitions/models.Game"
                }
            }
        },
        "models.Game": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "models.Map": {
            "type": "object",
            "properties": {
                "chapter_name": {
                    "type": "string"
                },
                "data": {},
                "game_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "map_name": {
                    "type": "string"
                }
            }
        },
        "models.MapCategoryScores": {
            "type": "object",
            "properties": {
                "any": {
                    "type": "integer"
                },
                "cm": {
                    "type": "integer"
                },
                "inbounds_sla": {
                    "type": "integer"
                },
                "no_sla": {
                    "type": "integer"
                }
            }
        },
        "models.MapHistory": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "runner_name": {
                    "type": "string"
                },
                "score_count": {
                    "type": "integer"
                }
            }
        },
        "models.MapRecords": {
            "type": "object",
            "properties": {
                "records": {}
            }
        },
        "models.MapShort": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.MapSummary": {
            "type": "object",
            "properties": {
                "category_scores": {
                    "$ref": "#/definitions/models.MapCategoryScores"
                },
                "description": {
                    "type": "string"
                },
                "history": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.MapHistory"
                    }
                },
                "rating": {
                    "type": "number"
                },
                "routers": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "showcase": {
                    "type": "string"
                }
            }
        },
        "models.ProfileResponse": {
            "type": "object",
            "properties": {
                "avatar_link": {
                    "type": "string"
                },
                "country_code": {
                    "type": "string"
                },
                "profile": {
                    "type": "boolean"
                },
                "scores_mp": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ScoreResponse"
                    }
                },
                "scores_sp": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ScoreResponse"
                    }
                },
                "steam_id": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "models.RankingsResponse": {
            "type": "object",
            "properties": {
                "rankings_mp": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.UserRanking"
                    }
                },
                "rankings_sp": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.UserRanking"
                    }
                }
            }
        },
        "models.RecordRequest": {
            "type": "object",
            "required": [
                "is_partner_orange",
                "partner_id",
                "score_count",
                "score_time"
            ],
            "properties": {
                "is_partner_orange": {
                    "type": "boolean"
                },
                "partner_id": {
                    "type": "string"
                },
                "score_count": {
                    "type": "integer"
                },
                "score_time": {
                    "type": "integer"
                }
            }
        },
        "models.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "models.ScoreResponse": {
            "type": "object",
            "properties": {
                "map_id": {
                    "type": "integer"
                },
                "records": {}
            }
        },
        "models.SearchResponse": {
            "type": "object",
            "properties": {
                "maps": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "id": {
                                "type": "integer"
                            },
                            "name": {
                                "type": "string"
                            }
                        }
                    }
                },
                "players": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "steam_id": {
                                "type": "string"
                            },
                            "user_name": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "models.UserRanking": {
            "type": "object",
            "properties": {
                "total_score": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "lp.ardapektezol.com/api",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Least Portals Database API",
	Description:      "Backend API endpoints for the Least Portals Database.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
