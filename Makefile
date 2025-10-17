DSN := "postgresql://postgres.taxxwvtptzqojzeywckr:0x9jz5YSiSqEjKHD@aws-1-ap-south-1.pooler.supabase.com:5432/postgres"
MODULE := "app"
TABLES := "users"

GenModelx:
	gentoolx -dsn $(DSN) \
		-outPath "./$(MODULE)/dbmodels" \
		--tables="$(TABLES)"

GenStruct:
	cd $(MODULE) && goctlx model struct --dsn $(DSN) -l postgres

GenSwagger:
	goctlx api plugin -plugin goctl-swagger="swagger -filename $(MODULE)/$(MODULE).json" -dir . -api "$(MODULE)/$(MODULE).api"