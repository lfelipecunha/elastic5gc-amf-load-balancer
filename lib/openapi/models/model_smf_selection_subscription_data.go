/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type SmfSelectionSubscriptionData struct {
	SupportedFeatures     string                `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures" mapstructure:"SupportedFeatures"`
	SubscribedSnssaiInfos map[string]SnssaiInfo `json:"subscribedSnssaiInfos,omitempty" yaml:"subscribedSnssaiInfos" bson:"subscribedSnssaiInfos" mapstructure:"SubscribedSnssaiInfos"`
	SharedSnssaiInfosId   string                `json:"sharedSnssaiInfosId,omitempty" yaml:"sharedSnssaiInfosId" bson:"sharedSnssaiInfosId" mapstructure:"SharedSnssaiInfosId"`
}
