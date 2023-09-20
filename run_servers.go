package main

import (
	"log"
	"os/exec"
	"path"
)

func main() {
	userRpcPath := path.Join(".", "user_rpc")
	notiRpcPath := path.Join(".", "noti_rpc")
	badgeRpcPath := path.Join(".", "badges_rpc")
	categoryRpcPath := path.Join(".", "category_rpc")
	featureRpcPath := path.Join(".", "features_rpc")
	mailRpcPath := path.Join(".", "mail_rpc")

	userRpcCmd := exec.Command("go", "run", ".")
	userRpcCmd.Dir = userRpcPath

	notiRpcCmd := exec.Command("go", "run", ".")
	notiRpcCmd.Dir = notiRpcPath

	badgeRpcCmd := exec.Command("go", "run", ".")
	badgeRpcCmd.Dir = badgeRpcPath

	categoryRpcCmd := exec.Command("go", "run", ".")
	categoryRpcCmd.Dir = categoryRpcPath

	featureRpcCmd := exec.Command("go", "run", ".")
	featureRpcCmd.Dir = featureRpcPath

	mailRpcCmd := exec.Command("go", "run", ".")
	mailRpcCmd.Dir = mailRpcPath

	// Start user_rpc server
	if err := userRpcCmd.Start(); err != nil {
		log.Fatal("Failed to start user_rpc server:", err)
	}
	log.Println("user_rpc server is running on port :8010")

	// Start badge_rpc server
	if err := badgeRpcCmd.Start(); err != nil {
		log.Fatal("Failed to start badge_rpc server:", err)
	}
	log.Println("badge_rpc server is running on port :8020")

	// Start category_rpc server
	if err := categoryRpcCmd.Start(); err != nil {
		log.Fatal("Failed to start category server:", err)
	}
	log.Println("category_rpc server is running on port :8030")

	// Start feature_rpc server
	if err := featureRpcCmd.Start(); err != nil {
		log.Fatal("Failed to start feature server:", err)
	}
	log.Println("feature server is running on port :8040")

	// Start mail_rpc server
	if err := mailRpcCmd.Start(); err != nil {
		log.Fatal("Failed to start mail_rpc server:", err)
	}
	log.Println("mail_rpc server is running on port :8050")

	// Start noti_rpc server
	if err := notiRpcCmd.Start(); err != nil {
		log.Fatal("Failed to start noti_rpc server:", err)
	}
	log.Println("noti_rpc server is running on port :8060")

	userRpcCmd.Wait()
	mailRpcCmd.Wait()
	badgeRpcCmd.Wait()
	categoryRpcCmd.Wait()
	featureRpcCmd.Wait()
	mailRpcCmd.Wait()
}
