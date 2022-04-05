package auth

import (
	"api-adminmv/middleware"
	"api-adminmv/util"
	"os"

	"github.com/gin-gonic/gin"
)

type PostCredential struct {
	ClientID     string `form:"client_id" json:"client_id"`
	ClientSecret string `form:"client_secret" json:"client_secret"`
}

func Router(r *gin.Engine) {
	r.POST("oauth/client_credential/accesstoken", func(c *gin.Context) {
		var req PostCredential
		if err := c.Bind(&req); err != nil {
			c.JSON(400, gin.H{
				"responseCode":    "01",
				"responseMessage": err.Error(),
			})
			return
		}

		if req.ClientID != os.Getenv("WSAD_ID") || req.ClientSecret != os.Getenv("WSAD_SECRET") {
			c.JSON(400, gin.H{
				"responseCode":    "01",
				"responseMessage": "invalid credential",
			})
			return
		}

		c.JSON(200, gin.H{
			"refresh_token_expires_in": "0",
			"api_product_list":         "[inquiry-sandbox]",
			"api_product_list_json": []string{
				"inquiry-sandbox",
			},
			"organization_name": "bri",
			"developer.email":   "furkorsan.gantheng@xyz.com",
			"token_type":        "BearerToken",
			"issued_at":         "1557891212144",
			"client_id":         "8E20dpP7KtakFkShw5tQHOFf7FFAU01o",
			"access_token":      "R04XSUbnm1GXNmDiXx9ysWMpFWBr",
			"application_name":  "317d0b2f-6536-4cac-a5f0-3bc9908815b3",
			"scope":             "",
			"expires_in":        "179999",
			"refresh_count":     "0",
			"status":            "approved",
		})
	})

	r.POST("v1.0/wsad/ADAuthentication2", middleware.TokenValidatorDummy(), func(c *gin.Context) {
		var req RequestLoginDummy

		if err := c.BindJSON(&req); err != nil {
			c.JSON(400, gin.H{
				"responseCode":        "01",
				"responseDescription": err.Error(),
			})
			return
		}

		if req.Username != "00001093" {
			c.JSON(404, gin.H{
				"responseCode":        "01",
				"responseDescription": "Username tidak ditemukan",
			})
			return
		}

		c.JSON(200, gin.H{
			"responseCode":        "00",
			"responseDescription": "Success",
		})
	})

	r.POST("apiBristars/1.0/requestDetailPekerja", func(c *gin.Context) {
		var username, password, ok = c.Request.BasicAuth()
		if username != "klasterku" && password != "P@ssw0rdKl45tErku" && !ok {
			c.JSON(401, gin.H{
				"responseCode":    "01",
				"responseMessage": "Invalid credential",
			})
			return
		}

		c.JSON(200, gin.H{
			"responseMessage": "Success",
			"responseData": gin.H{
				"pernr":                  "00001093",
				"nama":                   "Susilo",
				"jenis_kelamin":          "L",
				"area":                   "KW25",
				"desc_area":              "Kantor Wilayah Malang",
				"subarea":                "RT00",
				"desc_subarea":           "Kanwil Malang",
				"cost_center":            "WR85400",
				"desc_cost_center":       "KW Malang",
				"org_unit":               "50345773",
				"desc_org_unit":          "BAGIAN OPERASIONAL, JARINGAN, LAYANAN & PERFORMANCE MANAGEMENT",
				"jabatan":                "50196010",
				"desc_jabatan":           "KEPALA BAGIAN",
				"group_jabatan":          "012",
				"desc_group_jabatan":     "Kepala Bagian/Group Head",
				"jgpg":                   "JG12/PG16",
				"branch":                 "854",
				"group_jabatan_pgs":      "014",
				"desc_group_jabatan_pgs": "Pemimpin Cabang",
				"area_pgs":               "KW25",
				"desc_area_pgs":          "Kantor Wilayah Malang",
				"subarea_pgs":            "RT20",
				"desc_subarea_pgs":       "Mlg Martadinata",
				"cost_center_pgs":        "WR34400",
				"desc_cost_center_pgs":   "KANCA MALANG MARTADINATA",
				"org_unit_pgs":           "50345822",
				"desc_org_unit_pgs":      "KANTOR CABANG MALANG MARTADINATA",
				"branch_pgs":             "344",
				"tipe_pekerja":           "tetap",
			},
		})
	})

	r.Use(util.BasicAuth)

	v1 := r.Group("/v1")
	{
		v1.POST("auth/login", Login)

	}

}
