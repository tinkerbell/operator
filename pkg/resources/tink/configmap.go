package tink

import (
	"context"
	"fmt"
	"strings"
	"text/template"

	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func CreateNginxConfigMap(ctx context.Context, client ctrlruntimeclient.Client, clusterDNS, ns string) error {
	tmpl, err := template.New("nginx-conf").Parse(nginxConfigData)
	if err != nil {
		return fmt.Errorf("failed to parse nginx-conf template: %w", err)
	}

	data := struct {
		ClusterDNS string
	}{
		ClusterDNS: clusterDNS,
	}

	var buf strings.Builder
	if err = tmpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("failed to execute nginx-conf template: %w", err)
	}

	nginxConf := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "nginx-conf",
			Namespace: ns,
		},
		Data: map[string]string{
			"nginx.conf": buf.String(),
		},
	}

	if err := client.Create(ctx, nginxConf); err != nil {
		if kerrors.IsAlreadyExists(err) {
			return nil
		}

		return err
	}

	return nil
}

// TODO: parse nginx configs from args/operator configs
var nginxConfigData = `
worker_processes 1;
events {
    worker_connections  1024;
}
user root;

http {
  server {
    listen 80;
    location / {
      proxy_set_header X-Real-IP $remote_addr;
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
      resolver {{ .ClusterDNS }};
      set $smee_dns smee.tinkerbell.svc.cluster.local; # needed in Kubernetes for dynamic DNS resolution

      proxy_pass http://$smee_dns;
    }
  }

  server {
    listen 50061;
    location / {
      proxy_set_header X-Real-IP $remote_addr;
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
      resolver {{ .ClusterDNS }};
      set $hegel_dns hegel.tinkerbell.svc.cluster.local; # needed in Kubernetes for dynamic DNS resolution

      proxy_pass http://$hegel_dns:50061;
    }
  }

  server {
    listen 42113 http2;
    location / {
      proxy_set_header X-Real-IP $remote_addr;
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
      resolver {{ .ClusterDNS }};
      set $tink_dns tink-server.tinkerbell.svc.cluster.local; # needed in Kubernetes for dynamic DNS resolution

      grpc_pass grpc://$tink_dns:42113;
    }
  }

   server {
    listen 8080;
    location / {
      root /usr/share/nginx/html;
    }
  }
}

stream {
  log_format logger-json escape=json '{"source": "nginx", "time": $msec, "address": "$remote_addr", "status": $status, "upstream_addr": "$upstream_addr"}';

  server {
      listen 67 udp;
      resolver {{ .ClusterDNS }}; # needed in Kubernetes for dynamic DNS resolution
      set $smee_dns smee.tinkerbell.svc.cluster.local; # needed in Kubernetes for dynamic DNS resolution
      proxy_pass $smee_dns:67;
      proxy_bind $remote_addr:$remote_port transparent;
      proxy_responses 0;
      access_log /dev/stdout logger-json;
  }
  server {
      listen 69 udp;
      resolver {{ .ClusterDNS }};
      set $smee_dns smee.tinkerbell.svc.cluster.local; # needed in Kubernetes for dynamic DNS resolution
      proxy_pass $smee_dns:69;
      proxy_timeout 1s;
      access_log /dev/stdout logger-json;
  }
  server {
      listen 514 udp;
      resolver {{ .ClusterDNS }};
      set $smee_dns smee.tinkerbell.svc.cluster.local; # needed in Kubernetes for dynamic DNS resolution
      proxy_pass $smee_dns:514;
      proxy_bind $remote_addr:$remote_port transparent;
      proxy_responses 0;
      access_log /dev/stdout logger-json;
  }
}`
