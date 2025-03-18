package initializers

//import "ic-project/initializers"
import "ic-project/models"


func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}