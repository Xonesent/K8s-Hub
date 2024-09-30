package server

import (
	event_dropper "github.com/Xonesent/K8s-Hub/statistics-sender/internal/statistics_dropper/dropper_delivery/event"
	"github.com/Xonesent/K8s-Hub/statistics-sender/internal/statistics_dropper/dropper_usecase"
	timer_reminder "github.com/Xonesent/K8s-Hub/statistics-sender/internal/statistics_reminder/reminder_delivery/timer"
	"github.com/Xonesent/K8s-Hub/statistics-sender/internal/statistics_reminder/reminder_repository"
	"github.com/Xonesent/K8s-Hub/statistics-sender/internal/statistics_reminder/reminder_usecase"
	adminProto "github.com/Xonesent/K8s-Hub/statistics-sender/pkg/api/admin_proto"
	tgProto "github.com/Xonesent/K8s-Hub/statistics-sender/pkg/api/tg_proto"
	"github.com/Xonesent/K8s-Hub/statistics-sender/pkg/constant"
	"github.com/Xonesent/K8s-Hub/statistics-sender/pkg/dependency_connectors/kafka/cons_group"
)

func (s *Server) MapHandlers() (cons_group.Consumer, error) {
	reminderCHRepo := reminder_repository.NewReminderCHRepository(s.cfg, s.clickhouse)

	adminGrpc := adminProto.NewAdminServiceClient(s.cfg.GRPC.GrpcClientConn[constant.AdminGrpc])
	tgGrpc := tgProto.NewTgServiceClient(s.cfg.GRPC.GrpcClientConn[constant.TgGrpc])

	reminderUC := reminder_usecase.NewReminderUseCase(s.cfg, reminderCHRepo, s.kafkaProducer)
	dropperUC := dropper_usecase.NewDropperUseCase(s.cfg, adminGrpc, tgGrpc)

	reminderHDL := timer_reminder.NewReminderHandler(s.cfg, reminderUC)
	dropperHDL := event_dropper.NewDropperHandler(s.cfg, dropperUC)

	consumerHandler := cons_group.NewConsumerHandler(dropperHDL)

	timer_reminder.SetRemindTimer(s.cfg.Timers, reminderHDL)

	return consumerHandler, nil
}
