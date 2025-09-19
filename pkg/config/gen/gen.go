package main

import (
	cfg "github.com/conductorone/baton-strongdm/pkg/config"
	"github.com/conductorone/baton-sdk/pkg/config"
)

func main() {
	config.Generate("strongdm", cfg.Config)
}
