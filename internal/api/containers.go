
package api

import (
    "context"
    "encoding/json"
    "net/http"

    "github.com/docker/docker/api/types/container"
    "github.com/docker/docker/client"
)

func ListContainersHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        return
    }

    cli, err := client.NewClientWithOpts(
        client.FromEnv,
        client.WithAPIVersionNegotiation(),
    )
    if err != nil {
        http.Error(w, "failed to create Docker client: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer cli.Close()

    // Use container.ListOptions (newer SDK path)
    opts := container.ListOptions{
        All: true,
    }

    containers, err := cli.ContainerList(context.Background(), opts)
    if err != nil {
        http.Error(w, "failed to list containers: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // containers is []types.Container from github.com/docker/docker/api/types
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(containers)
}
