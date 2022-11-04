package controller

import (
	"context"
	"fmt"
	"github.com/moadqassem/kubetink/pkg/resources/boots"
	"github.com/moadqassem/kubetink/pkg/resources/hegel"
	"github.com/moadqassem/kubetink/pkg/resources/rufio"
	"github.com/moadqassem/kubetink/pkg/resources/tink"
)

func (r *Reconciler) ensureTinkerbellDeployments(ctx context.Context) error {
	if err := boots.CreateDeployment(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create boots deployment: %v", err)
	}

	if err := hegel.CreateDeployment(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create hegel deployment: %v", err)
	}

	if err := rufio.CreateDeployment(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create rufio deployment: %v", err)
	}

	if err := tink.CreateTinkControllerDeployment(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create tink controller deployment: %v", err)
	}

	if err := tink.CreateTinkServerDeployment(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create tink server deployment: %v", err)
	}

	if err := tink.CreateTinkStackDeployment(ctx, r.Client, r.namespace); err != nil {
		return fmt.Errorf("failed to create tink stack deployment: %v", err)
	}

	return nil
}
