package middlewares

import (
	"go-blog/databases"
	"go-blog/models/permissions"
	"go-blog/tools"

	"github.com/labstack/echo/v4"
)

// CheckPermissions is a middleware that checks if the user has the required permissions.
func CheckPermissions(db databases.Database, permission permissions.Permission) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			userID, _ := tools.ExtractUserID(ctx)
			hasPermission := HavePermissions(db, userID, permission)

			ctx.Set("permissable", hasPermission)
			return next(ctx)
		}
	}
}

// HavePermissions checks if the user has the required permissions, this can be used in other parts of the application.
func HavePermissions(db databases.Database, userID uint, permission permissions.Permission) bool {
	for _, perm := range db.GetUserPermissions(userID) {
		if perm == permission ||
			perm == permissions.FullAccess ||
			(perm == permissions.FullContents && permission > permissions.FullContents) {
			return true
		}
	}
	return false
}
