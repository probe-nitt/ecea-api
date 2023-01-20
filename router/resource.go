package router

import (
	"github.com/ecea-nitt/ecea-server/controllers"
	"github.com/ecea-nitt/ecea-server/middlewares"
	"github.com/labstack/echo/v4"
)

func StudyMaterialRoutes(e *echo.Group, c controllers.StudyMaterialController) {
	studyMaterial := e.Group("/studymaterial")

	studyMaterial.POST("/add", middlewares.Authorizer(c.AddStudyMaterial))
	studyMaterial.PUT("/edit/document", middlewares.Authorizer(c.UpdateStudyMaterialURL))
	//	studyMaterial.PUT("/edit/name", middlewares.Authorizer(c.EditRessourceName))
	//	studyMaterial.PUT("/edit/bookid", middlewares.Authorizer(c.EditMemberRole))
	studyMaterial.PUT("/edit/subject", middlewares.Authorizer(c.UpdateStudyMaterialSubject))
	studyMaterial.GET("/getcat/:category", c.GetCategoryStudyMaterials)
	studyMaterial.GET("/get/:name", c.GetStudyMaterial)
	studyMaterial.DELETE("/delete/:name", middlewares.Authorizer(c.DeleteStudyMaterial))

}
