package app

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/fx"

	customerController "github.com/viniciuscluna/tc-fiap-customer/internal/customer/controller"
	customerRepositories "github.com/viniciuscluna/tc-fiap-customer/internal/customer/domain/repositories"
	customerApiController "github.com/viniciuscluna/tc-fiap-customer/internal/customer/infrastructure/api/controller"
	customerPersistence "github.com/viniciuscluna/tc-fiap-customer/internal/customer/infrastructure/persistence"
	customerPresenter "github.com/viniciuscluna/tc-fiap-customer/internal/customer/presenter"
	customerUseCasesAdd "github.com/viniciuscluna/tc-fiap-customer/internal/customer/usecase/addCustomer"
	customerUseCasesGetByCpf "github.com/viniciuscluna/tc-fiap-customer/internal/customer/usecase/getbycpf"

	"github.com/viniciuscluna/tc-fiap-customer/pkg/rest"
	"github.com/viniciuscluna/tc-fiap-customer/pkg/storage/dynamodb"
)

func InitializeApp() *fx.App {
	return fx.New(
		fx.Provide(
			dynamodb.NewDynamoDBClient,
			fx.Annotate(customerPersistence.NewCustomerRepositoryImpl, fx.As(new(customerRepositories.CustomerRepository))),
			fx.Annotate(customerUseCasesAdd.NewAddCustomerUseCaseImpl, fx.As(new(customerUseCasesAdd.AddCustomerUseCase))),
			fx.Annotate(customerUseCasesGetByCpf.NewGetByCpfUseCaseImpl, fx.As(new(customerUseCasesGetByCpf.GetByCpfUseCase))),
			fx.Annotate(customerController.NewCustomerControllerImpl, fx.As(new(customerController.CustomerController))),
			fx.Annotate(customerPresenter.NewCustomerPresenterImpl, fx.As(new(customerPresenter.CustomerPresenter))),
			chi.NewRouter,
			func(customerController customerController.CustomerController) []rest.Controller {
				return []rest.Controller{
					customerApiController.NewCustomerController(customerController),
				}
			},
		),
		fx.Invoke(registerRoutes),
		fx.Invoke(startHTTPServer),
	)
}

func registerRoutes(r *chi.Mux, controllers []rest.Controller) {
	r.Use(middleware.Logger)

	// Swagger UI
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The URL pointing to API definition
	))

	for _, controller := range controllers {
		controller.RegisterRoutes(r)
	}
}

func startHTTPServer(lc fx.Lifecycle, r *chi.Mux) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Println("Starting HTTP server on :8080")
				if err := http.ListenAndServe(":8080", r); err != nil {
					log.Fatalf("Failed to start HTTP server: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Shutting down HTTP server gracefully")
			return nil
		},
	})
}
