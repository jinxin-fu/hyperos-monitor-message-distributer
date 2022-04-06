/*
Copyright 2022 Xin Jin.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"fmt"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	infrav1 "distributer/api/v1"
)

// ObjectReconciler reconciles a Object object
type ObjectReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=infra.distributer.pml.com.cn,resources=objects,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=infra.distributer.pml.com.cn,resources=objects/status,verbs=get;update;patch

func (r *ObjectReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
        ctx := context.Background()
	_ = r.Log.WithValues("object", req.NamespacedName)

	// your logic here
        // 1. Print Spec.Detail and Status.Created in log
        obj := &infrav1.Object{}
        if err := r.Get(ctx, req.NamespacedName, obj); err != nil {
            fmt.Errorf("couldn't find object:%s", req.String())
        } else {
        //打印Detail和Created
            r.Log.V(1).Info("Successfully get detail", "Detail", obj.Spec.Detail)
            r.Log.V(1).Info("", "Created", obj.Status.Created)
        }
        // 2. Change Created
        if !obj.Status.Created {
            obj.Status.Created = true
            r.Update(ctx, obj)
        }
	return ctrl.Result{}, nil
}

func (r *ObjectReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&infrav1.Object{}).
		Complete(r)
}
