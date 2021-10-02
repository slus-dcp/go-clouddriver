package kubernetes_test

import (
	"errors"
	"net/http"

	"github.com/homedepot/go-clouddriver/internal/kubernetes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Patch", func() {
	BeforeEach(func() {
		setup()
	})

	JustBeforeEach(func() {
		kubernetesController.Patch(c, patchManifestRequest)
	})

	When("getting the provider returns an error", func() {
		BeforeEach(func() {
			fakeSQLClient.GetKubernetesProviderReturns(kubernetes.Provider{}, errors.New("error getting provider"))
		})

		It("returns an error", func() {
			Expect(c.Writer.Status()).To(Equal(http.StatusBadRequest))
			Expect(c.Errors.Last().Error()).To(Equal("internal: error getting kubernetes provider spin-cluster-account: error getting provider"))
		})
	})

	When("patching the manifest returns an error", func() {
		BeforeEach(func() {
			fakeKubeClient.PatchUsingStrategyReturns(kubernetes.Metadata{}, nil, errors.New("error patching manifest"))
		})

		It("returns an error", func() {
			Expect(c.Writer.Status()).To(Equal(http.StatusInternalServerError))
			Expect(c.Errors.Last().Error()).To(Equal("error patching manifest"))
		})
	})

	When("creating the resource returns an error", func() {
		BeforeEach(func() {
			fakeSQLClient.CreateKubernetesResourceReturns(errors.New("error creating resource"))
		})

		It("returns an error", func() {
			Expect(c.Writer.Status()).To(Equal(http.StatusInternalServerError))
			Expect(c.Errors.Last().Error()).To(Equal("error creating resource"))
		})
	})

	Context("merge strategies", func() {
		Context("strategic patch type", func() {
			BeforeEach(func() {
				patchManifestRequest.Options.MergeStrategy = "strategic"
			})

			It("succeeds", func() {
				Expect(c.Writer.Status()).To(Equal(http.StatusOK))
			})
		})

		Context("json patch type", func() {
			BeforeEach(func() {
				patchManifestRequest.Options.MergeStrategy = "json"
			})

			It("succeeds", func() {
				Expect(c.Writer.Status()).To(Equal(http.StatusOK))
			})
		})

		Context("merge patch type", func() {
			BeforeEach(func() {
				patchManifestRequest.Options.MergeStrategy = "merge"
			})

			It("succeeds", func() {
				Expect(c.Writer.Status()).To(Equal(http.StatusOK))
			})
		})

		Context("unknown patch type", func() {
			BeforeEach(func() {
				patchManifestRequest.Options.MergeStrategy = "unknown"
			})

			It("returns an error", func() {
				Expect(c.Writer.Status()).To(Equal(http.StatusBadRequest))
				Expect(c.Errors.Last().Error()).To(Equal("invalid merge strategy unknown"))
			})
		})

		It("succeeds", func() {
			Expect(c.Writer.Status()).To(Equal(http.StatusOK))
			kind, name, namespace, _, strategy := fakeKubeClient.PatchUsingStrategyArgsForCall(0)
			Expect(string(kind)).To(Equal("deployment"))
			Expect(string(name)).To(Equal("test-deployment"))
			Expect(string(namespace)).To(Equal(""))
			Expect(string(strategy)).To(Equal("application/strategic-merge-patch+json"))
		})
	})

	When("Using a namespace-scoped provider", func() {
		BeforeEach(func() {
			fakeSQLClient.GetKubernetesProviderReturns(namespaceScopedProvider, nil)
		})

		When("the kind is not supported", func() {
			BeforeEach(func() {
				patchManifestRequest.ManifestName = "namespace someNamespace"
			})

			It("returns an error", func() {
				Expect(c.Writer.Status()).To(Equal(http.StatusBadRequest))
				Expect(c.Errors.Last().Error()).To(Equal("namespace-scoped account not allowed to access cluster-scoped kind: 'namespace'"))
			})
		})

		When("the kind is supported", func() {
			It("succeeds", func() {
				Expect(c.Writer.Status()).To(Equal(http.StatusOK))
				kind, name, namespace, _, strategy := fakeKubeClient.PatchUsingStrategyArgsForCall(0)
				Expect(string(kind)).To(Equal("deployment"))
				Expect(string(name)).To(Equal("test-deployment"))
				Expect(string(namespace)).To(Equal("provider-namespace"))
				Expect(string(strategy)).To(Equal("application/strategic-merge-patch+json"))
			})
		})
	})
})