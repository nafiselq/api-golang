package cmd

import (
	"fmt"
	"os"

	"github.com/e-ziswaf/eziswaf-api/config"
	"github.com/e-ziswaf/eziswaf-api/internal/app/appcontext"
	"github.com/e-ziswaf/eziswaf-api/internal/app/commons"
	"github.com/e-ziswaf/eziswaf-api/internal/app/repository"
	"github.com/e-ziswaf/eziswaf-api/internal/app/server"
	"github.com/e-ziswaf/eziswaf-api/internal/app/service"
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "eziswaf-api",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
			examples and usage of using your application.`,
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
}

func start() {
	var err error
	cfg := config.ProviderConfig()
	appConfig, err := config.NewAppConfig(cfg)
	if err != nil {
		logrus.Fatalf("ProviderConfig error : %s", err)
	}

	app := appcontext.NewAppContext(cfg)

	var dbMysql *gorm.DB
	if cfg.GetBool("mysql.is_enabled") {
		dbMysql, err = app.GetDBInstance(appcontext.DBDialectMysql)
		if err != nil {
			logrus.Fatalf("Failed to start, error connect to DB MySQL | %v", err)
			return
		}
	}

	var dbPostgre *gorm.DB
	if cfg.GetBool("postgre.is_enabled") {
		dbPostgre, err = app.GetDBInstance(appcontext.DBDialectPostgres)
		if err != nil {
			logrus.Fatalf("Failed to start, error connect to DB Postgre | %v", err)
			return
		}
	}

	var cache *redis.Pool
	if cfg.GetBool("cache.is_enabled") {
		cache = app.GetCachePool()
		cacheConn, err := cache.Dial()
		if err != nil {
			logrus.Fatalf("Failed to start, error connect to DB Cache | %v", err)
			return
		}
		defer cacheConn.Close()
	}

	opt := commons.Options{
		AppCtx:         app,
		ProviderConfig: cfg,
		AppConfig:      appConfig,
		DbMysql:        dbMysql,
		DbPostgre:      dbPostgre,
		CachePool:      cache,
	}

	repo := wiringRepository(repository.Option{
		Options: opt,
	})

	services := wiringService(service.Option{
		Options:    opt,
		Repository: repo,
	})

	newServer := server.NewServer(opt, services)

	// run app
	newServer.StartApp()
}

func wiringRepository(repoOption repository.Option) *repository.Repository {
	// wiring up all your repos here
	cacheRepo := repository.NewCacheRepository(repoOption)
	personRepo := repository.NewPersonRepository(repoOption)
	donorRepo := repository.NewDonorRepository(repoOption)
	campaignRepo := repository.NewCampaignRepository(repoOption)
	donationRepo := repository.NewDonationRepository(repoOption)
	lembagaRepo := repository.NewLembagaRepository(repoOption)

	repo := repository.Repository{
		Cache:    cacheRepo,
		Person:   personRepo,
		Donor:    donorRepo,
		Campaign: campaignRepo,
		Donation: donationRepo,
		Lembaga:  lembagaRepo,
	}

	return &repo
}

func wiringService(serviceOption service.Option) *service.Services {
	// wiring up all services
	hc := service.NewHealthCheck(serviceOption)
	hello := service.NewHelloService(serviceOption)
	donor := service.NewDonorService(serviceOption)
	campaign := service.NewCampaignService(serviceOption)
	donation := service.NewDonationService(serviceOption)

	svc := service.Services{
		HealthCheck: hc,
		Hello:       hello,
		Donor:       donor,
		Campaign:    campaign,
		Donation:    donation,
	}

	return &svc
}
