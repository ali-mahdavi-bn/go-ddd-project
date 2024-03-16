package service_layer

import (
	"github.com/labstack/echo/v4"
	"strings"
)

var Path = ""

func Add() string {
	baseDocTemplate := `{
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
    "paths": {{Paths}}
}`

	docTemplate := strings.Replace(baseDocTemplate, "{{Paths}}", "{"+Path+"}", -1)
	return docTemplate
}

func isString(status any, defaultValue string) string {
	value, ok := status.(string)
	if ok == false {
		return defaultValue
	}
	return value
}

func Handlers(e *APIRouter, method string, path string, handler HandlerFunc, swaggerConfig map[string]any) {
	fullPath := e.Path + path

	if e.InitializerSwagger != true {
		HandleMethod(e, method, fullPath, handler)
	} else {
		swaggerConfig["path"] = fullPath
		swaggerConfig["method"] = method
		swaggerConfig["tag"] = e.Tag
		GenerateSwaggerConfig(swaggerConfig)
	}

}
func HandleMethod(e *APIRouter, method string, path string, handler HandlerFunc) {
	switch strings.ToLower(method) {
	case "get":
		MethodGetEcho(e, path, handler)
	case "post":
		MethodPostEcho(e, path, handler)
	case "put":
		MethodPutEcho(e, path, handler)
	}
}
func GenerateSwaggerConfig(swaggerConfig map[string]any) {
	description, _ := swaggerConfig["description"]
	summary, _ := swaggerConfig["summary"]
	status, _ := swaggerConfig["status"]
	path, _ := swaggerConfig["path"]
	method, _ := swaggerConfig["method"]
	tag, _ := swaggerConfig["tag"]
	baseConfig := `"{{Path}}": {
            "{{method}}": {
                "description": "{{description}}",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "{{root}}"
                ],
                "summary": "{{summary}}",
                "responses": {
                    "{{status}}": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },`
	config := strings.Replace(baseConfig, "{{Path}}", isString(path, "/"), -1)
	config = strings.Replace(config, "{{method}}", isString(strings.ToLower(method.(string)), "get"), -1)
	config = strings.Replace(config, "{{status}}", isString(status, "200"), -1)
	config = strings.Replace(config, "{{summary}}", isString(summary, ""), -1)
	config = strings.Replace(config, "{{description}}", isString(description, "description"), -1)
	config = strings.Replace(config, "{{root}}", isString(tag, "root"), -1)

	Path += config
}

func MethodGetEcho(e *APIRouter, path string, handler HandlerFunc) {
	e.Router.GET(path, echo.HandlerFunc(handler))
}

func MethodPostEcho(e *APIRouter, path string, handler HandlerFunc) {
	e.Router.POST(path, echo.HandlerFunc(handler))
}

func MethodPutEcho(e *APIRouter, path string, handler HandlerFunc) {
	e.Router.PUT(path, echo.HandlerFunc(handler))
}
