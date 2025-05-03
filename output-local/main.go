package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/v3/pkg/generated/controllers/core"
	"github.com/rancher/wrangler/v3/pkg/generic"
	"github.com/rancher/wrangler/v3/pkg/kubeconfig"
	"github.com/rancher/wrangler/v3/pkg/schemes"
	"github.com/tomleb/make-make-quick-quick/output-local/foo/pkg/generated"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var Scheme = runtime.NewScheme()

func init() {
	metav1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(schemes.AddToScheme(Scheme))
	generated.Init()
}

func main() {
	restConfig, err := kubeconfig.GetNonInteractiveClientConfig(os.Getenv("KUBECONFIG")).ClientConfig()
	must(err)

	controllerFactory, err := controller.NewSharedControllerFactoryFromConfigWithOptions(restConfig, Scheme, nil)
	must(err)

	opts := &generic.FactoryOptions{
		SharedControllerFactory: controllerFactory,
	}

	core, err := core.NewFactoryFromConfigWithOptions(restConfig, opts)
	must(err)

	ctx := context.Background()

	Register(ctx, core)

	err = controllerFactory.SharedCacheFactory().Start(ctx)
	must(err)

	controllerFactory.SharedCacheFactory().WaitForCacheSync(ctx)

	log.Println("Starting controllers")
	err = controllerFactory.Start(ctx, 10)
	must(err)

	log.Println("Running")
	<-ctx.Done()
}

func Register(ctx context.Context, core *core.Factory) {
	log.Println("Configuring OnChange handler")
	core.Core().V1().ConfigMap().OnChange(ctx, "on-change", func(key string, cm *corev1.ConfigMap) (*corev1.ConfigMap, error) {
		if cm != nil && cm.Namespace != "foo" {
			return cm, nil
		}

		newCm := cm.DeepCopy()
		if !setAnnotations(&newCm.ObjectMeta, "foo", "bar") {
			return cm, nil
		}

		log.Println("Updated ConfigMap", key)
		updatedCm, err := core.Core().V1().ConfigMap().Update(newCm)
		if err != nil {
			return nil, fmt.Errorf("update: %w", err)
		}

		return updatedCm, nil
	})
}

func setAnnotations(objMeta *metav1.ObjectMeta, key string, value string) bool {
	ann := objMeta.GetAnnotations()
	if ann == nil {
		ann = make(map[string]string)
	}
	if ann[key] == value {
		return false
	}
	ann[key] = value
	objMeta.SetAnnotations(ann)
	return true
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
