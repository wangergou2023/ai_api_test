# APP_NAME = one_chat
# APP_NAME = input_image_chat
# APP_NAME = tts
# APP_NAME = stt

API_KEY = sk-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
BASE_URL = https://llxspace.shop/v1

build:
	GOOS=linux GOARCH=amd64 go build -o $(APP_NAME) $(APP_NAME).go

run:
	OPENAI_API_KEY=$(API_KEY) OPENAI_BASE_URL=$(BASE_URL) go run $(APP_NAME).go

clean:
	rm -f $(APP_NAME)
