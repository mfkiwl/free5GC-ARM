/*
 * Npcf_PolicyAuthorization Service API
 *
 * This is the Policy Authorization Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Identifies the events the application subscribes to.
type EventsSubscReqData struct {
	Events   []AfEventSubscription `json:"events" bson:"events"`
	NotifUri string                `json:"notifUri,omitempty" bson:"notifUri"`
	UsgThres *UsageThreshold       `json:"usgThres,omitempty" bson:"usgThres"`
}
