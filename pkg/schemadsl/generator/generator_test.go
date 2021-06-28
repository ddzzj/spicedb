package generator

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/authzed/spicedb/pkg/namespace"
	v0 "github.com/authzed/spicedb/pkg/proto/authzed/api/v0"
)

func TestGenerator(t *testing.T) {

	type generatorTest struct {
		name     string
		input    *v0.NamespaceDefinition
		expected string
		okay     bool
	}

	tests := []generatorTest{
		{
			"empty",
			namespace.Namespace("foo/test"),
			"definition foo/test {}",
			true,
		},
		{
			"simple relation",
			namespace.Namespace("foo/test",
				namespace.Relation("somerel", nil, &v0.RelationReference{
					Namespace: "foo/bar",
					Relation:  "hiya",
				}),
			),
			`definition foo/test {
	relation somerel: foo/bar#hiya
}`,
			true,
		},
		{
			"simple permission",
			namespace.Namespace("foo/test",
				namespace.Relation("someperm", namespace.Union(
					namespace.ComputedUserset("anotherrel"),
				)),
			),
			`definition foo/test {
	permission someperm = anotherrel
}`,
			true,
		},
		{
			"complex permission",
			namespace.Namespace("foo/test",
				namespace.Relation("someperm", namespace.Union(
					namespace.Rewrite(
						namespace.Exclusion(
							namespace.ComputedUserset("a"),
							namespace.ComputedUserset("b"),
							namespace.TupleToUserset("y", "z"),
						),
					),
					namespace.ComputedUserset("c"),
				)),
			),
			`definition foo/test {
	permission someperm = (a - b - y->z) + c
}`,
			true,
		},
		{
			"legacy relation",
			namespace.Namespace("foo/test",
				namespace.Relation("somerel", namespace.Union(
					namespace.This(),
					namespace.ComputedUserset("anotherrel"),
				), &v0.RelationReference{
					Namespace: "foo/bar",
					Relation:  "hiya",
				}),
			),
			`definition foo/test {
	relation somerel: foo/bar#hiya = /* _this unsupported here. Please rewrite into a relation and permission */ + anotherrel
}`,
			false,
		},
		{
			"missing type information",
			namespace.Namespace("foo/test",
				namespace.Relation("somerel", nil),
			),
			`definition foo/test {
	relation somerel: /* missing allowed types */
}`,
			false,
		},

		{
			"full example",
			namespace.Namespace("foo/document",
				namespace.Relation("owner", nil,
					&v0.RelationReference{
						Namespace: "foo/user",
						Relation:  "...",
					},
				),
				namespace.Relation("reader", nil,
					&v0.RelationReference{
						Namespace: "foo/user",
						Relation:  "...",
					},
					&v0.RelationReference{
						Namespace: "foo/group",
						Relation:  "member",
					},
				),
				namespace.Relation("read", namespace.Union(
					namespace.ComputedUserset("reader"),
					namespace.ComputedUserset("owner"),
				)),
			),
			`definition foo/document {
	relation owner: foo/user
	relation reader: foo/user | foo/group#member
	permission read = reader + owner
}`,
			true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require := require.New(t)
			source, ok := GenerateSource(test.input)
			require.Equal(test.expected, source)
			require.Equal(test.okay, ok)
		})
	}
}
