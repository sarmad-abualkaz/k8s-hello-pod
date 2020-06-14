package status

import (
	"net/http"
	"fmt"
	"bytes"
	"k8s-hello-pod/config"
	"encoding/json"
)

// Return the pod metadata as json:
func GiveHealth(w http.ResponseWriter, r *http.Request) {
	
	pod := config.ConfigEnv.Pod
	b:= new(bytes.Buffer)

	json.NewEncoder(b).Encode(pod)

	fmt.Fprint(w, b)
}