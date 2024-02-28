version=1.0.1

generate:
	$(MAKE) -C ./modules/go-client-generator/ generate-go-client service=order version=${version}
