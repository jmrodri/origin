// This file was automatically generated by informer-gen

package internalversion

import (
	network "github.com/openshift/origin/pkg/sdn/apis/network"
	internalinterfaces "github.com/openshift/origin/pkg/sdn/generated/informers/internalversion/internalinterfaces"
	internalclientset "github.com/openshift/origin/pkg/sdn/generated/internalclientset"
	internalversion "github.com/openshift/origin/pkg/sdn/generated/listers/network/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// EgressNetworkPolicyInformer provides access to a shared informer and lister for
// EgressNetworkPolicies.
type EgressNetworkPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.EgressNetworkPolicyLister
}

type egressNetworkPolicyInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newEgressNetworkPolicyInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.Network().EgressNetworkPolicies(v1.NamespaceAll).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.Network().EgressNetworkPolicies(v1.NamespaceAll).Watch(options)
			},
		},
		&network.EgressNetworkPolicy{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *egressNetworkPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&network.EgressNetworkPolicy{}, newEgressNetworkPolicyInformer)
}

func (f *egressNetworkPolicyInformer) Lister() internalversion.EgressNetworkPolicyLister {
	return internalversion.NewEgressNetworkPolicyLister(f.Informer().GetIndexer())
}
