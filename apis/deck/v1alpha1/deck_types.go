/*
Copyright 2020 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// DeckCard is used to create the list of cards in a Deck Resource
// We are not relying on the managed resource Card because this dictionary
// is mocking an external service.
//+kubebuilder:skipversion
type DeckCard struct {
	Suit string `json:"suit,omitempty"`
	Rank string `json:"rank,omitempty"`
}

// DeckParameters are the configurable fields of a Deck.
type DeckParameters struct {
	// Decks can be dealt from specific types, like Shuffled, Ordered, or Empty.
	// Empty decks can be useful for creating "discard" piles
	Type string `json:"type,omitempty"`
}

// DeckObservation are the observable fields of a Card.
type DeckObservation struct {
	Cards []DeckCard `json:"cards,omitempty"`
}

// DeckSpec defines the desired state of Deck
type DeckSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DeckParameters `json:"forProvider"`
}

// DeckStatus defines the observed state of Deck
type DeckStatus struct {
	xpv1.ResourceStatus `json:",inline"`

	AtProvider DeckObservation `json:"atProvider,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster

// Deck is the Schema for the decks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,template}
//go:generate go run ../../../vendor/github.com/crossplane/crossplane-tools/cmd/angryjet/main.go generate-methodsets ./...
type Deck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeckSpec   `json:"spec,omitempty"`
	Status DeckStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DeckList contains a list of Deck
type DeckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Deck `json:"items"`
}

// Deck type metadata.
var (
	DeckKind             = reflect.TypeOf(Deck{}).Name()
	DeckGroupKind        = schema.GroupKind{Group: Group, Kind: DeckKind}.String()
	DeckKindAPIVersion   = DeckKind + "." + SchemeGroupVersion.String()
	DeckGroupVersionKind = SchemeGroupVersion.WithKind(DeckKind)
)

func init() {
	SchemeBuilder.Register(&Deck{}, &DeckList{})
}
