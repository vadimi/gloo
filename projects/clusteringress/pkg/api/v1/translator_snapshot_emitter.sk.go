// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"sync"
	"time"

	gloo_solo_io "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/errutils"
)

var (
	mTranslatorSnapshotIn  = stats.Int64("translator.clusteringress.gloo.solo.io/snap_emitter/snap_in", "The number of snapshots in", "1")
	mTranslatorSnapshotOut = stats.Int64("translator.clusteringress.gloo.solo.io/snap_emitter/snap_out", "The number of snapshots out", "1")

	translatorsnapshotInView = &view.View{
		Name:        "translator.clusteringress.gloo.solo.io_snap_emitter/snap_in",
		Measure:     mTranslatorSnapshotIn,
		Description: "The number of snapshots updates coming in",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
	translatorsnapshotOutView = &view.View{
		Name:        "translator.clusteringress.gloo.solo.io/snap_emitter/snap_out",
		Measure:     mTranslatorSnapshotOut,
		Description: "The number of snapshots updates going out",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
)

func init() {
	view.Register(translatorsnapshotInView, translatorsnapshotOutView)
}

type TranslatorEmitter interface {
	Register() error
	Secret() gloo_solo_io.SecretClient
	Upstream() gloo_solo_io.UpstreamClient
	ClusterIngress() ClusterIngressClient
	Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *TranslatorSnapshot, <-chan error, error)
}

func NewTranslatorEmitter(secretClient gloo_solo_io.SecretClient, upstreamClient gloo_solo_io.UpstreamClient, clusterIngressClient ClusterIngressClient) TranslatorEmitter {
	return NewTranslatorEmitterWithEmit(secretClient, upstreamClient, clusterIngressClient, make(chan struct{}))
}

func NewTranslatorEmitterWithEmit(secretClient gloo_solo_io.SecretClient, upstreamClient gloo_solo_io.UpstreamClient, clusterIngressClient ClusterIngressClient, emit <-chan struct{}) TranslatorEmitter {
	return &translatorEmitter{
		secret:         secretClient,
		upstream:       upstreamClient,
		clusterIngress: clusterIngressClient,
		forceEmit:      emit,
	}
}

type translatorEmitter struct {
	forceEmit      <-chan struct{}
	secret         gloo_solo_io.SecretClient
	upstream       gloo_solo_io.UpstreamClient
	clusterIngress ClusterIngressClient
}

func (c *translatorEmitter) Register() error {
	if err := c.secret.Register(); err != nil {
		return err
	}
	if err := c.upstream.Register(); err != nil {
		return err
	}
	if err := c.clusterIngress.Register(); err != nil {
		return err
	}
	return nil
}

func (c *translatorEmitter) Secret() gloo_solo_io.SecretClient {
	return c.secret
}

func (c *translatorEmitter) Upstream() gloo_solo_io.UpstreamClient {
	return c.upstream
}

func (c *translatorEmitter) ClusterIngress() ClusterIngressClient {
	return c.clusterIngress
}

func (c *translatorEmitter) Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *TranslatorSnapshot, <-chan error, error) {
	errs := make(chan error)
	var done sync.WaitGroup
	ctx := opts.Ctx
	/* Create channel for Secret */
	type secretListWithNamespace struct {
		list      gloo_solo_io.SecretList
		namespace string
	}
	secretChan := make(chan secretListWithNamespace)
	/* Create channel for Upstream */
	type upstreamListWithNamespace struct {
		list      gloo_solo_io.UpstreamList
		namespace string
	}
	upstreamChan := make(chan upstreamListWithNamespace)
	/* Create channel for ClusterIngress */

	for _, namespace := range watchNamespaces {
		/* Setup namespaced watch for Secret */
		secretNamespacesChan, secretErrs, err := c.secret.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Secret watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, secretErrs, namespace+"-secrets")
		}(namespace)
		/* Setup namespaced watch for Upstream */
		upstreamNamespacesChan, upstreamErrs, err := c.upstream.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Upstream watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, upstreamErrs, namespace+"-upstreams")
		}(namespace)

		/* Watch for changes and update snapshot */
		go func(namespace string) {
			for {
				select {
				case <-ctx.Done():
					return
				case secretList := <-secretNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case secretChan <- secretListWithNamespace{list: secretList, namespace: namespace}:
					}
				case upstreamList := <-upstreamNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case upstreamChan <- upstreamListWithNamespace{list: upstreamList, namespace: namespace}:
					}
				}
			}
		}(namespace)
	}
	/* Setup cluster-wide watch for ClusterIngress */

	clusterIngressChan, clusterIngressErrs, err := c.clusterIngress.Watch(opts)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "starting ClusterIngress watch")
	}
	done.Add(1)
	go func() {
		defer done.Done()
		errutils.AggregateErrs(ctx, errs, clusterIngressErrs, "clusteringresses")
	}()

	snapshots := make(chan *TranslatorSnapshot)
	go func() {
		originalSnapshot := TranslatorSnapshot{}
		currentSnapshot := originalSnapshot.Clone()
		timer := time.NewTicker(time.Second * 1)
		sync := func() {
			if originalSnapshot.Hash() == currentSnapshot.Hash() {
				return
			}

			stats.Record(ctx, mTranslatorSnapshotOut.M(1))
			originalSnapshot = currentSnapshot.Clone()
			sentSnapshot := currentSnapshot.Clone()
			snapshots <- &sentSnapshot
		}

		/* TODO (yuval-k): figure out how to make this work to avoid a stale snapshot.
		   		// construct the first snapshot from all the configs that are currently there
		   		// that guarantees that the first snapshot contains all the data.
		   		for range watchNamespaces {
		      secretNamespacedList := <- secretChan
		      currentSnapshot.Secrets.Clear(secretNamespacedList.namespace)
		      secretList := secretNamespacedList.list
		   	currentSnapshot.Secrets.Add(secretList...)
		      upstreamNamespacedList := <- upstreamChan
		      currentSnapshot.Upstreams.Clear(upstreamNamespacedList.namespace)
		      upstreamList := upstreamNamespacedList.list
		   	currentSnapshot.Upstreams.Add(upstreamList...)
		   		}
		*/

		for {
			record := func() { stats.Record(ctx, mTranslatorSnapshotIn.M(1)) }

			select {
			case <-timer.C:
				sync()
			case <-ctx.Done():
				close(snapshots)
				done.Wait()
				close(errs)
				return
			case <-c.forceEmit:
				sentSnapshot := currentSnapshot.Clone()
				snapshots <- &sentSnapshot
			case secretNamespacedList := <-secretChan:
				record()

				namespace := secretNamespacedList.namespace
				secretList := secretNamespacedList.list

				currentSnapshot.Secrets.Clear(namespace)
				currentSnapshot.Secrets.Add(secretList...)
			case upstreamNamespacedList := <-upstreamChan:
				record()

				namespace := upstreamNamespacedList.namespace
				upstreamList := upstreamNamespacedList.list

				currentSnapshot.Upstreams.Clear(namespace)
				currentSnapshot.Upstreams.Add(upstreamList...)
			case clusterIngressList := <-clusterIngressChan:
				record()
				currentSnapshot.Clusteringresses = clusterIngressList
			}
		}
	}()
	return snapshots, errs, nil
}
