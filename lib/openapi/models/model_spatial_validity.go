/*
 * Npcf_PolicyAuthorization Service API
 *
 * This is the Policy Authorization Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// describes explicitly the route to an Application location
type SpatialValidity struct {
	PresenceInfoList map[string]PresenceInfo `json:"presenceInfoList" bson:"presenceInfoList"`
}
