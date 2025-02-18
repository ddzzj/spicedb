package development

import (
	"fmt"

	devinterface "github.com/authzed/spicedb/pkg/proto/developer/v1"
	v1 "github.com/authzed/spicedb/pkg/proto/dispatch/v1"
	"github.com/authzed/spicedb/pkg/tuple"
	"github.com/authzed/spicedb/pkg/validationfile/blocks"
)

const maxDispatchDepth = 25

// RunAllAssertions runs all assertions found in the given assertions block against the
// developer context, returning whether any errors occurred.
func RunAllAssertions(devContext *DevContext, assertions *blocks.Assertions) ([]*devinterface.DeveloperError, error) {
	trueFailures, err := runAssertions(devContext, assertions.AssertTrue, true, "Expected relation or permission %s to exist")
	if err != nil {
		return nil, err
	}

	falseFailures, err := runAssertions(devContext, assertions.AssertFalse, false, "Expected relation or permission %s to not exist")
	if err != nil {
		return nil, err
	}

	failures := append(trueFailures, falseFailures...)
	return failures, nil
}

func runAssertions(devContext *DevContext, assertions []blocks.Assertion, expected bool, fmtString string) ([]*devinterface.DeveloperError, error) {
	var failures []*devinterface.DeveloperError

	// TODO(jschorr): Support caveats via some sort of `assertMaybe`?
	for _, assertion := range assertions {
		tpl := tuple.MustFromRelationship(assertion.Relationship)

		tplString, err := tuple.String(tpl)
		if err != nil {
			return nil, err
		}

		if tpl.Caveat != nil {
			failures = append(failures, &devinterface.DeveloperError{
				Message: fmt.Sprintf("cannot specify a caveat on an assertion: `%s`", tplString),
				Source:  devinterface.DeveloperError_ASSERTION,
				Kind:    devinterface.DeveloperError_UNKNOWN_RELATION,
				Context: tplString,
				Line:    uint32(assertion.SourcePosition.LineNumber),
				Column:  uint32(assertion.SourcePosition.ColumnPosition),
			})
			continue
		}

		cr, _, err := RunCheck(devContext, tpl.ResourceAndRelation, tpl.Subject)
		if err != nil {
			devErr, wireErr := DistinguishGraphError(
				devContext,
				err,
				devinterface.DeveloperError_ASSERTION,
				uint32(assertion.SourcePosition.LineNumber),
				uint32(assertion.SourcePosition.ColumnPosition),
				tplString,
			)
			if wireErr != nil {
				return nil, wireErr
			}
			if devErr != nil {
				failures = append(failures, devErr)
			}
		} else if (cr == v1.ResourceCheckResult_MEMBER) != expected {
			failures = append(failures, &devinterface.DeveloperError{
				Message: fmt.Sprintf(fmtString, tplString),
				Source:  devinterface.DeveloperError_ASSERTION,
				Kind:    devinterface.DeveloperError_ASSERTION_FAILED,
				Context: tplString,
				Line:    uint32(assertion.SourcePosition.LineNumber),
				Column:  uint32(assertion.SourcePosition.ColumnPosition),
			})
		}
	}

	return failures, nil
}
