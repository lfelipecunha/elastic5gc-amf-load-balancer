/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type RmInfo struct {
	RmState    RmState    `json:"rmState" bson:"rmState"`
	AccessType AccessType `json:"accessType" bson:"accessType"`
}
