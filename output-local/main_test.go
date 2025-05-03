package main

import (
	"testing"
	"time"

	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/v3/pkg/generated/controllers/core"
	"github.com/rancher/wrangler/v3/pkg/generic"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
)

func TestRegister(t *testing.T) {
	testEnv := envtest.Environment{}
	restConfig, err := testEnv.Start()
	mustT(t, err)

	controllerFactory, err := controller.NewSharedControllerFactoryFromConfigWithOptions(restConfig, Scheme, nil)
	mustT(t, err)

	opts := &generic.FactoryOptions{
		SharedControllerFactory: controllerFactory,
	}

	core, err := core.NewFactoryFromConfigWithOptions(restConfig, opts)
	mustT(t, err)

	ctx := t.Context()

	Register(ctx, core)

	err = controllerFactory.SharedCacheFactory().Start(ctx)
	must(err)

	controllerFactory.SharedCacheFactory().WaitForCacheSync(ctx)

	err = controllerFactory.Start(ctx, 10)
	must(err)

	cmCtrl := core.Core().V1().ConfigMap()
	nsCtrl := core.Core().V1().Namespace()
	_, err = nsCtrl.Create(&corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "foo",
		},
	})
	mustT(t, err)
	_, err = nsCtrl.Create(&corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "bar",
		},
	})
	mustT(t, err)

	tests := []struct {
		name string
		cm   *corev1.ConfigMap

		expectedAnnotation bool
	}{
		{
			name: "empty annotation is set",
			cm: &corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "toto",
					Namespace: "foo",
				},
			},
			expectedAnnotation: true,
		},
		{
			name: "annotation is overridden",
			cm: &corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "toto",
					Namespace: "foo",
					Annotations: map[string]string{
						"foo": "toto",
					},
				},
			},
			expectedAnnotation: true,
		},
		{
			name: "annotation is not set in default namespace",
			cm: &corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "toto1",
					Namespace: "default",
				},
			},
		},
		{
			name: "annotation is not set in bar namespace",
			cm: &corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "toto2",
					Namespace: "bar",
				},
			},
		},
		{
			name: "annotation is not removed in other namespaces",
			cm: &corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "toto",
					Namespace: "bar",
					Annotations: map[string]string{
						"foo": "bar",
					},
				},
			},
			expectedAnnotation: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cm := test.cm
			cmCtrl.Create(cm)
			time.Sleep(1 * time.Second)
			gotCm, err := cmCtrl.Get(cm.Namespace, cm.Name, metav1.GetOptions{})
			mustT(t, err)

			if test.expectedAnnotation {
				if gotCm.GetAnnotations() == nil || gotCm.GetAnnotations()["foo"] != "bar" {
					t.Errorf("annotation not set")
				}
			} else {
				if gotCm.GetAnnotations() != nil && gotCm.GetAnnotations()["foo"] == "bar" {
					t.Errorf("annotation set unexpectedly")
				}
			}

		})
	}
}

func mustT(t *testing.T, err error) {
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
}
