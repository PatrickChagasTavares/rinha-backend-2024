package swagger

import (
	"github.com/patrickchagastavares/rinha-backend-2024/docs"
	"github.com/patrickchagastavares/rinha-backend-2024/pkg/httpRouter"
	httpSwagger "github.com/swaggo/http-swagger"
)

func New(router httpRouter.Router) {

	docs.SwaggerInfo.Title = "Swagger about router of rinha-backend"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router.Get("swagger/*any", router.ParseHandler(
		httpSwagger.Handler(httpSwagger.URL("doc.json")),
	))
}
