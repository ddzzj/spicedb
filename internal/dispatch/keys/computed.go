package keys

import (
	"fmt"

	v1 "github.com/authzed/spicedb/pkg/proto/dispatch/v1"
	"github.com/authzed/spicedb/pkg/tuple"
)

// cachePrefix defines a unique prefix for a type of cache key.
type cachePrefix string

// Define the various prefixes for the cache entries. These must *all* be unique and must *all*
// also be placed into the cachePrefixes slice below.
const (
	checkViaRelationPrefix   cachePrefix = "cr"
	checkViaCanonicalPrefix  cachePrefix = "cc"
	lookupPrefix             cachePrefix = "l"
	expandPrefix             cachePrefix = "e"
	reachableResourcesPrefix cachePrefix = "rr"
	lookupSubjectsPrefix     cachePrefix = "ls"
)

var cachePrefixes = []cachePrefix{
	checkViaRelationPrefix,
	checkViaCanonicalPrefix,
	lookupPrefix,
	expandPrefix,
	reachableResourcesPrefix,
	lookupSubjectsPrefix,
}

// checkRequestToKey converts a check request into a cache key based on the relation
func checkRequestToKey(req *v1.DispatchCheckRequest, option dispatchCacheKeyHashComputeOption) DispatchCacheKey {
	return dispatchCacheKeyHash(checkViaRelationPrefix, req.Metadata.AtRevision, option,
		hashableRelationReference{req.ResourceRelation},
		hashableIds(req.ResourceIds),
		hashableOnr{req.Subject},
		hashableResultSetting(req.ResultsSetting),
	)
}

// checkRequestToKeyWithCanonical converts a check request into a cache key based
// on the canonical key.
func checkRequestToKeyWithCanonical(req *v1.DispatchCheckRequest, canonicalKey string) DispatchCacheKey {
	if canonicalKey == "" {
		panic(fmt.Sprintf("given empty canonical key for request: %s => %s", req.ResourceRelation, tuple.StringONR(req.Subject)))
	}

	// NOTE: canonical cache keys are only unique *within* a version of a namespace.
	return dispatchCacheKeyHash(checkViaCanonicalPrefix, req.Metadata.AtRevision, computeBothHashes,
		hashableString(req.ResourceRelation.Namespace),
		hashableString(canonicalKey),
		hashableIds(req.ResourceIds),
		hashableOnr{req.Subject},
		hashableResultSetting(req.ResultsSetting),
	)
}

// lookupRequestToKey converts a lookup request into a cache key
func lookupRequestToKey(req *v1.DispatchLookupRequest, option dispatchCacheKeyHashComputeOption) DispatchCacheKey {
	return dispatchCacheKeyHash(lookupPrefix, req.Metadata.AtRevision, option,
		hashableRelationReference{req.ObjectRelation},
		hashableOnr{req.Subject},
		hashableContext{req.Context}, // NOTE: context is included here because lookup does a single dispatch
	)
}

// expandRequestToKey converts an expand request into a cache key
func expandRequestToKey(req *v1.DispatchExpandRequest, option dispatchCacheKeyHashComputeOption) DispatchCacheKey {
	return dispatchCacheKeyHash(expandPrefix, req.Metadata.AtRevision, option,
		hashableOnr{req.ResourceAndRelation},
	)
}

// reachableResourcesRequestToKey converts a reachable resources request into a cache key
func reachableResourcesRequestToKey(req *v1.DispatchReachableResourcesRequest, option dispatchCacheKeyHashComputeOption) DispatchCacheKey {
	return dispatchCacheKeyHash(reachableResourcesPrefix, req.Metadata.AtRevision, option,
		hashableRelationReference{req.ResourceRelation},
		hashableRelationReference{req.SubjectRelation},
		hashableIds(req.SubjectIds),
	)
}

// lookupSubjectsRequestToKey converts a lookup subjects request into a cache key
func lookupSubjectsRequestToKey(req *v1.DispatchLookupSubjectsRequest, option dispatchCacheKeyHashComputeOption) DispatchCacheKey {
	return dispatchCacheKeyHash(lookupSubjectsPrefix, req.Metadata.AtRevision, option,
		hashableRelationReference{req.ResourceRelation},
		hashableRelationReference{req.SubjectRelation},
		hashableIds(req.ResourceIds),
	)
}
