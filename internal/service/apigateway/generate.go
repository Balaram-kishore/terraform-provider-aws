//go:generate go run ../../generate/listpages/main.go -ListOps=GetAuthorizers -Paginator=Position
//go:generate go run ../../generate/tags/main.go -ServiceTagsMap -UpdateTags
//go:generate go run ../../generate/servicepackage/main.go
// ONLY generate directives and package declaration! Do not add anything else to this file.

package apigateway
