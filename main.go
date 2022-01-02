package main

import (
	"cart/common"
	"cart/domain/repository"
	service2 "cart/domain/service"
	"cart/handler"
	pb "cart/proto/cart"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/debug/log"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	// 配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}
	// 注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	// Create service
	service := micro.NewService(
		micro.Name("go.micro.service.category"),
		micro.Version("latest"),
		// 这里设置地址和需要暴露的端口
		micro.Address("127.0.0.1:8082"),
		// 添加consul作为注册中心
		micro.Registry(consulRegistry),
	)

	//只执行一次,数据表初始化
	//rp := repository.NewUserRepository(db)
	//rp.InitTable()

	// 获取mysql配置,路径中不带前缀
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	// 连接数据库
	db, err := gorm.Open("mysql", mysqlInfo.User + ":" + mysqlInfo.Pwd + "@/" + mysqlInfo.Database + "?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Error(err)
	}
	defer db.Close()

	// 禁止复表
	db.SingularTable(true)

	rp := repository.NewCategoryRepository(db)
	rp.InitTable()
	// 初始化服务
	service.Init()

	categoryDataService := service2.NewCategoryDataService(repository.NewCategoryRepository(db))

	// Register handler
	err = pb.RegisterCategoryHandler(service.Server(), &handler.Cart{categoryDataService})
	if err != nil {
		log.Error(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
