// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/gocrane/api/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Analytics returns a AnalyticsInformer.
	Analytics() AnalyticsInformer
	// ConfigSets returns a ConfigSetInformer.
	ConfigSets() ConfigSetInformer
	// Recommendations returns a RecommendationInformer.
	Recommendations() RecommendationInformer
	// RecommendationRules returns a RecommendationRuleInformer.
	RecommendationRules() RecommendationRuleInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Analytics returns a AnalyticsInformer.
func (v *version) Analytics() AnalyticsInformer {
	return &analyticsInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ConfigSets returns a ConfigSetInformer.
func (v *version) ConfigSets() ConfigSetInformer {
	return &configSetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Recommendations returns a RecommendationInformer.
func (v *version) Recommendations() RecommendationInformer {
	return &recommendationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// RecommendationRules returns a RecommendationRuleInformer.
func (v *version) RecommendationRules() RecommendationRuleInformer {
	return &recommendationRuleInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
