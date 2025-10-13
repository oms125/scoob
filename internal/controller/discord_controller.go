/*
Copyright 2025.

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

package controller

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/bwmarrin/discordgo"
	"github.com/go-logr/logr"
	configv1 "scoob.ritsec.cloud/kubebuilder/api/v1"
	"scoob.ritsec.cloud/kubebuilder/internal/bot"
)

// DiscordReconciler reconciles a Discord object
type DiscordReconciler struct {
	client.Client
	Scheme            *runtime.Scheme
	DiscordBotManager *bot.DiscordBotManager
}

// +kubebuilder:rbac:groups=config.ritsec.cloud,resources=discords,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=config.ritsec.cloud,resources=discords/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=config.ritsec.cloud,resources=discords/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Discord object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.22.1/pkg/reconcile
func (r *DiscordReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := logf.FromContext(ctx)

	var discList configv1.DiscordList

	err := r.List(ctx, &discList)
	log(err, logger, "Failed to fetch API list")

	discToken := discList.Items[0].Spec.Token

	fmt.Println("TOKEN: " + discToken)

	discSession, err := discordgo.New("Bot " + discToken)
	log(err, logger, "Failed to start Discord bot session")

	discSession.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages

	err = discSession.Open()
	log(err, logger, "Failed to open Discord bot session")

	err = r.DiscordBotManager.SetSession(discSession)
	log(err, logger, "Failed to set Discord bot Session")

	err = r.DiscordBotManager.SendMessage(discList.Items[0].Spec.Channels.LogChannel, "SCOOB has arrived!")
	log(err, logger, "Failed to send Discord bot join message")

	return ctrl.Result{}, nil
}

func log(err error, logger logr.Logger, s ...string) {
	if err != nil {
		if len(s) == 0 {
			logger.Error(err, "")
		} else {
			logger.Error(err, s[0])
		}
	}
}

// SetupWithManager sets up the controller with the Manager.
func (r *DiscordReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&configv1.Discord{}).
		Named("discord").
		Complete(r)
}
