package apijson

import (
	"testing"

	"github.com/tidwall/gjson"
)

type structuralUnion interface {
	structuralUnionVariant()
}

type structuralCat struct {
	Meow string `json:"meow"`
}

func (structuralCat) structuralUnionVariant() {}

type structuralDog struct {
	Bark string `json:"bark"`
}

func (structuralDog) structuralUnionVariant() {}

func TestUntaggedObjectUnionUsesStructuralExactness(t *testing.T) {
	RegisterUnion[structuralUnion](
		"",
		Variant[structuralCat](gjson.JSON),
		Variant[structuralDog](gjson.JSON),
	)

	var decoded structuralUnion
	if err := Unmarshal([]byte(`{"bark":"woof"}`), &decoded); err != nil {
		t.Fatalf("Unmarshal returned an error: %v", err)
	}

	if _, ok := decoded.(structuralDog); !ok {
		t.Fatalf("expected structuralDog, got %T", decoded)
	}
}
