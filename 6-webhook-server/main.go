package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	v1 "k8s.io/api/admission/v1"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var allowedNamespaces []string

func getNamespacesFromConfigMap() ([]string, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to create Kubernetes config: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kubernetes client: %v", err)
	}

	cm, err := clientset.CoreV1().ConfigMaps("webhook").Get(context.TODO(), "webhook-namespace-config", metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get configmap: %v", err)
	}

	namespaces := strings.Split(cm.Data["namespaces"], "\n")
	return namespaces, nil
}

func serve(w http.ResponseWriter, r *http.Request) {
	var admissionReviewReq v1.AdmissionReview
	if err := json.NewDecoder(r.Body).Decode(&admissionReviewReq); err != nil {
		http.Error(w, "could not decode request", http.StatusBadRequest)
		return
	}

	podNamespace := admissionReviewReq.Request.Namespace

	isAllowedNamespace := false
	for _, ns := range allowedNamespaces {
		if ns == podNamespace {
			isAllowedNamespace = true
			break
		}
	}

	if !isAllowedNamespace {
		admissionReviewRes := v1.AdmissionReview{
			Response: &v1.AdmissionResponse{
				UID:     admissionReviewReq.Request.UID,
				Allowed: true,
			},
		}
		json.NewEncoder(w).Encode(admissionReviewRes)
		return
	}

	admissionReviewRes := v1.AdmissionReview{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "admission.k8s.io/v1",
			Kind:       "AdmissionReview",
		},
		Response: &v1.AdmissionResponse{
			UID: admissionReviewReq.Request.UID,
		},
	}

	if admissionReviewReq.Request.Kind.Kind != "Deployment" {
		admissionReviewRes.Response.Allowed = true
		json.NewEncoder(w).Encode(admissionReviewRes)
		return
	}

	// Deployment kaynağını çözümle
	deployment := &appsv1.Deployment{}
	if err := json.Unmarshal(admissionReviewReq.Request.Object.Raw, deployment); err != nil {
		http.Error(w, "could not unmarshal deployment", http.StatusBadRequest)
		return
	}

	// Deployment içindeki her bir container'ın CPU limitini kontrol et
	for _, container := range deployment.Spec.Template.Spec.Containers {
		if container.Resources.Limits.Cpu().Cmp(resource.MustParse("1")) > 0 {
			admissionReviewRes.Response.Allowed = false
			admissionReviewRes.Response.Result = &metav1.Status{
				Message: fmt.Sprintf("Deployment %s contains a container %s with CPU limit greater than 1", deployment.Name, container.Name),
			}
			json.NewEncoder(w).Encode(admissionReviewRes)
			return
		}
	}

	admissionReviewRes.Response.Allowed = true
	json.NewEncoder(w).Encode(admissionReviewRes)
}

func main() {
	var err error
	allowedNamespaces, err = getNamespacesFromConfigMap()
	if err != nil {
		log.Fatalf("Error loading namespaces from ConfigMap: %v", err)
	}

	http.HandleFunc("/validate", serve)
	log.Fatal(http.ListenAndServeTLS(":8443", "/etc/webhook/certs/tls.crt", "/etc/webhook/certs/tls.key", nil))
}
