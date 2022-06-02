package controllers

import (
	"ar5go/app/serializers"
	"ar5go/app/svc"
	"ar5go/infra/errors"
	"ar5go/infra/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

type company struct {
	lc   logger.LogClient
	cSvc svc.ICompany
}

// NewCompanyController will initialize the controllers
func NewCompanyController(grp interface{}, lc logger.LogClient, cSvc svc.ICompany) {
	cc := &company{
		lc:   lc,
		cSvc: cSvc,
	}

	g := grp.(*echo.Group)

	g.POST("/v1/company/signup", cc.CreateWithAdminUser)
}

// swagger:route POST /v1/company/signup Company CreateCompany
// Create a new company
// responses:
//	201: genericSuccessResponse
//	404: errorResponse
//	400: errorResponse
//	500: errorResponse

// CreateWithAdminUser handles POST requests and create a new company with a default admin user
func (ctr *company) CreateWithAdminUser(c echo.Context) error {
	var cp serializers.CompanyReq

	if err := c.Bind(&cp); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		return c.JSON(restErr.Status, restErr)
	}

	result, saveErr := ctr.cSvc.CreateCompanyWithAdminUser(cp)
	if saveErr != nil {
		return c.JSON(saveErr.Status, saveErr)
	}

	return c.JSON(http.StatusCreated, result)
}
