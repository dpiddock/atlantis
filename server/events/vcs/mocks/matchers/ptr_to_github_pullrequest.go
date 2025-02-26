// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"github.com/petergtz/pegomock"
	"reflect"

	github "github.com/google/go-github/v49/github"
)

func AnyPtrToGithubPullRequest() *github.PullRequest {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*github.PullRequest))(nil)).Elem()))
	var nullValue *github.PullRequest
	return nullValue
}

func EqPtrToGithubPullRequest(value *github.PullRequest) *github.PullRequest {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *github.PullRequest
	return nullValue
}

func NotEqPtrToGithubPullRequest(value *github.PullRequest) *github.PullRequest {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue *github.PullRequest
	return nullValue
}

func PtrToGithubPullRequestThat(matcher pegomock.ArgumentMatcher) *github.PullRequest {
	pegomock.RegisterMatcher(matcher)
	var nullValue *github.PullRequest
	return nullValue
}
