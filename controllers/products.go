package controllers

import (
	"errors"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
	"path/filepath"
	"strconv"

	"github.com/MohamedmuhsinJ/shopify/database"
	"github.com/MohamedmuhsinJ/shopify/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Products struct {
	ProductID   uint
	ProductName string
	ActualPrice string
	Price       string
	Image       string
	SideImage   string
	ZoomImage   string
	Description string
	Color       string
	Brands      string
	Stock       uint
	Category    string
	Size        uint
}

func ListALL(c *gin.Context) {
	var brandFilter []models.Brand
	var categoryFilter []models.Category
	var sizeFilter []models.ShoeSize
	if brand := c.Query("brandSearch"); brand != "" {
		brandFiles := database.Db.Where("brands LIKE ?", "%"+brand+"%").Find(&brandFilter)
		if brandFiles.Error != nil {
			c.JSON(404, gin.H{
				"error": brandFiles.Error.Error(),
			})
			c.Abort()
			return
		}
	}
	if category := c.Query("categorySearch"); category != "" {
		categoryFiles := database.Db.Where("category LIKE ?", "%"+category+"%").Find(&categoryFilter)
		if categoryFiles.Error != nil {
			c.JSON(404, gin.H{
				"error": categoryFiles.Error.Error(),
			})
			c.Abort()
			return
		}
	}

	if size := c.Query("sizeSearch"); size != "" {
		sizes, _ := strconv.Atoi(size)
		sizeFiles := database.Db.Where("size= ?", sizes).Find(&sizeFilter)
		if sizeFiles.Error != nil {
			c.JSON(404, gin.H{
				"error": sizeFiles.Error.Error(),
			})
			c.Abort()
			return
		}
	}
	c.JSON(200, gin.H{
		"available brnds":     brandFilter,
		"avaiable categories": categoryFilter,
		"available sizes":     sizeFilter,
	})
}

func AddProduct(c *gin.Context) {
	prodName := c.PostForm("productName")
	Price := c.PostForm("price")
	price, _ := strconv.Atoi(Price)
	description := c.PostForm("description")
	color := c.PostForm("color")
	brand := c.PostForm("brandID")
	brands, _ := strconv.Atoi(brand)
	Stock := c.PostForm("stock")
	stock, _ := strconv.Atoi(Stock)
	Category := c.PostForm("categoryID")
	category, _ := strconv.Atoi(Category)
	Size := c.PostForm("sizeID")
	size, _ := strconv.Atoi(Size)
	//ading main image
	imagePath, _ := c.FormFile("image")
	err := imageVerify(imagePath)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}
	extension := filepath.Ext(imagePath.Filename)
	image := uuid.New().String() + extension
	c.SaveUploadedFile(imagePath, "./public/"+image)
	//adding sideimage
	SideImage, _ := c.FormFile("sideImage")
	er := imageVerify(SideImage)
	if er != nil {
		c.JSON(400, gin.H{
			"err": err.Error(),
		})
		c.Abort()
		return
	}
	extension = filepath.Ext(SideImage.Filename)
	sideImage := uuid.New().String() + extension
	c.SaveUploadedFile(SideImage, "./public/"+sideImage)

	//adding zoom images

	ZoomImage, _ := c.FormFile("zoomImage")
	e := imageVerify(ZoomImage)
	if e != nil {
		c.JSON(400, gin.H{
			"err": err.Error(),
		})
		c.Abort()
		return
	}
	extension = filepath.Ext(ZoomImage.Filename)
	zoomimage := uuid.New().String() + extension
	c.SaveUploadedFile(ZoomImage, "./public/"+zoomimage)
	Discount := c.PostForm("discount")
	discount, _ := strconv.Atoi(Discount)
	BrandDiscount := c.PostForm("brandDiscount")
	brandDiscount, _ := strconv.Atoi(BrandDiscount)
	var disc int
	res := database.Db.Raw("update brands set discount=? where id=?", brandDiscount, brands).Scan(&models.Brand{})
	if res.Error != nil {
		c.JSON(404, gin.H{
			"err": res.Error.Error(),
		})
		c.Abort()
		return
	}
	//comparing which discount is greater

	if brandDiscount > discount {
		disc = (price * brandDiscount) / 100
	} else {
		disc = (price * discount) / 100
	}

	var count uint
	database.Db.Raw("select count(*) from products where product_name=?", prodName).Scan(&count)
	if count > 0 {
		c.JSON(404, gin.H{
			"count":   count,
			"message": "already a product exists with same name",
		})
		c.Abort()
		return
	}

	product := models.Product{
		ProductName: prodName,
		Price:       uint(price) - uint(disc),
		ActualPrice: uint(price),
		Color:       color,
		Description: description,
		Discount:    uint(disc),
		BrandId:     uint(brands),
		CategoryID:  uint(category),
		ShoeSizeID:  uint(size),
		Image:       image,
		SideImage:   sideImage,
		ZoomImage:   zoomimage,
		Stock:       uint(stock),
	}
	record := database.Db.Create(&product)
	if record.Error != nil {
		c.JSON(404, gin.H{
			"error": "product already exists",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"message": "added successfully",
	})
}

type Editproduct struct {
	Price    uint   `json:"price,omitempty"`
	Color    string `json:"color,omitempty"`
	Stock    uint   `json:"stock,omitempty"`
	Discount uint   `json:"discount,omitempty"`
}

func EditProduct(c *gin.Context) {

	id := c.Param("id")
	var editProduct Editproduct
	if err := c.ShouldBindJSON(&editProduct); err != nil {
		c.JSON(404, gin.H{
			"error": "failed to read ",
		})
		c.Abort()
		return
	}
	dis := (editProduct.Price * editProduct.Discount) / 100

	var product models.Product

	rec := database.Db.Model(product).Where("product_id=?", id).Updates(models.Product{ActualPrice: editProduct.Price, Price: editProduct.Price - dis, Color: editProduct.Color, Stock: editProduct.Stock, Discount: dis})
	if rec.Error != nil {
		c.JSON(400, gin.H{
			"error": rec.Error.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(202, gin.H{
		"message": "updates successfully",
	})
}

func DeleteProducts(c *gin.Context) {
	id := c.Param("id")
	var products models.Product
	database.Db.First(&products, id)
	if products.ProductId == 0 {
		c.JSON(400, gin.H{
			"error": "product doesnot exists",
		})
		c.Abort()
		return
	}
	rec := database.Db.Delete(&products)
	if rec.Error != nil {
		c.JSON(400, gin.H{
			"error": "Cannot delete product",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"message": "product deleted successfully",
	})
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	var products []Products
	record := database.Db.Raw("SELECT product_id,product_name,actual_price,price,image,side_image,zoom_image,color,description,stock,brands.brands,categories.category,shoe_sizes.size FROM products join brands on products.brand_id=brands.id join categories on products.category_id=categories.id join shoe_sizes on products.shoe_size_id=shoe_sizes.id where product_id=?", id).Scan(&products)
	if record.Error != nil {
		c.JSON(400, gin.H{
			"error": record.Error.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"product": products,
	})
}

func imageVerify(img *multipart.FileHeader) (err error) {
	src, _ := img.Open()
	defer src.Close()
	imge, _, err := image.Decode(src)

	if err != nil {
		err = errors.New("invalid  image format")
		return
	}
	if imge == nil {

		err = errors.New("invalid")
		return
	}
	return

}

func ProductView(c *gin.Context) {
	var products []Products
	database.Db.Raw("SELECT product_id,product_name,price,actual_price,image,side_image,zoom_image,color,stock,description,brands.brands,categories.category,shoe_sizes.size FROM products join brands on products.brand_id=brands.id join categories on products.category_id=categories.id join shoe_sizes on products.shoe_size_id=shoe_sizes.id").Scan(&products)

	for i, _ := range products {
		c.JSON(200, gin.H{
			"products": products[i],
		})

	}

}
func ProductSearch(c *gin.Context) {
	var products []Products
	if search := c.Query("search"); search != "" {
		database.Db.Raw("SELECT product_id,product_name,price,actual_price,image,side_image,zoom_image,color,stock,description,brands.brands,categories.category,shoe_sizes.size FROM products join brands on products.brand_id=brands.id join categories on products.category_id=categories.id join shoe_sizes on products.shoe_size_id=shoe_sizes.id where product_name like ?", "%"+search+"%").Scan(&products)
		for i, _ := range products {
			c.JSON(200, gin.H{
				"products": products[i],
			})
		}
	}
}
func ProductSort(c *gin.Context) {
	var products []Products
	if sort := c.Query("sort"); sort == "" {
		database.Db.Raw("SELECT product_id,product_name,price,actual_price,image,side_image,zoom_image,color,stock,description,brands.brands,categories.category,shoe_sizes.size FROM products join brands on products.brand_id=brands.id join categories on products.category_id=categories.id join shoe_sizes on products.shoe_size_id=shoe_sizes.id order by price").Scan(&products)
		for i, _ := range products {
			c.JSON(200, gin.H{
				"products": products[i],
			})
		}
	}
}
