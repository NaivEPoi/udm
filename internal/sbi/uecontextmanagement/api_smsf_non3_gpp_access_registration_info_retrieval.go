/*
 * Nudm_UECM
 *
 * Nudm Context Management Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package uecontextmanagement

import (
	"net/http"

	"github.com/free5gc/openapi"
	udm_context "github.com/free5gc/udm/internal/context"
	"github.com/gin-gonic/gin"
)

// GetSmsfNon3gppAccess - retrieve the SMSF registration for non-3GPP access information
func HTTPGetSmsfNon3gppAccess(c *gin.Context) {
	scopes := []string{"nudm-uecm"}
	_, oauth_err := openapi.CheckOAuth(c.Request.Header.Get("Authorization"), scopes)
	if oauth_err != nil && udm_context.UDM_Self().OAuth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": oauth_err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
