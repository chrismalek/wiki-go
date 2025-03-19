package types

import (
	"html/template"
	"time"
	"wiki-go/internal/comments"
	"wiki-go/internal/config"
)

// NavItem represents a navigation item (directory)
type NavItem struct {
	Title    string
	Path     string
	IsDir    bool
	Children []*NavItem
	IsActive bool
}

// BreadcrumbItem represents an item in the breadcrumb trail
type BreadcrumbItem struct {
	Title  string
	Path   string
	IsLast bool
}

// PageData represents the data passed to the template
type PageData struct {
	Navigation         *NavItem
	Content            template.HTML
	DirContent         template.HTML
	Breadcrumbs        []BreadcrumbItem
	Config             *config.Config
	LastModified       time.Time
	CurrentDir         *NavItem           // Current directory as a NavItem
	Title              string             // Page title
	IsLoginPage        bool               // Whether this is the login page
	AvailableLanguages []string           // Available languages for the UI
	Comments           []comments.Comment // Comments for the document
	CommentsAllowed    bool               // Whether comments are allowed for this document
	IsAuthenticated    bool               // Whether the user is authenticated
	IsAdmin            bool               // Whether the user is an admin
	DocPath            string             // Document path for API calls
}
