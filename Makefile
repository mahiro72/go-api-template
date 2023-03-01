.PHONY: help
help: ## helpを表示する
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: up
up: ## docker環境を立ち上げる
	docker-compose -f docker-compose.yml up -d --build

.PHONY: down
down: ## docker環境を閉じる
	docker-compose -f docker-compose.yml down

.PHONY: fclean
fclean: down del-volumes

.PHONY: re
re: fclean up ## volumesを削除してdocker環境を立ち上げる

.PHONY: del-volumes
del-volumes:
	rm -rf ./docker/dev/db/mysql/data

.PHONY: log
log: ## docker環境のログを標示する
	docker-compose logs -f

.PHONY: go-fmt
go-fmt: ## goのコードを整形します
	gofmt -l -w .

.PHONY: set-up
set-up: ## 環境の初期化を行います
	@./scripts/setup.sh && echo "set-up command completed successfully 👍"
