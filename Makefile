.PHONY: proto-gen
proto-gen:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/hello.proto

.PHONY: docker-tag
docker-tag:
	docker tag grpc-sample-$(target):latest us-central1-docker.pkg.dev/velvety-glazing-420809/grpc-sample-repo/grpc-sample-$(target):latest

.PHONY: docker-push
docker-push:
	docker push us-central1-docker.pkg.dev/velvety-glazing-420809/grpc-sample-repo/grpc-sample-$(target):latest

.PHONY: gcloud-deploy
gcloud-deploy:
	gcloud run services replace ./.infra/$(target).yaml --project=velvety-glazing-420809 --region=us-central1