/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type ChfInfo struct {
	SupiRangeList *[]SupiRange `json:"supiRangeList,omitempty" bson:"supiRangeList"`

	GpsiRangeList *[]IdentityRange `json:"gpsiRangeList,omitempty" bson:"gpsiRangeList"`

	PlmnRangeList *[]PlmnRange `json:"plmnRangeList,omitempty" bson:"plmnRangeList"`
}
