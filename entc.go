package main

import (
	"log"

	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
)

func main() {
	cfg := &gen.Config{
		Target:   "ent",
		Package:  "entsegv/ent",
		IDType:   &field.TypeInfo{Type: field.TypeString},
		Features: []gen.Feature{gen.FeatureUpsert},
	}

	eoas, err := entoas.NewExtension()
	if err != nil {
		log.Fatalf("creating entoas extension: %v", err)
	}

	exts := entc.Extensions(&encodeExtension{}, eoas)
	if err := entc.Generate("./schema", cfg, exts); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
