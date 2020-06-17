# Create .env file
touch .env
# Set your telegram bot api key token on .env file
apikeybot=your-api-key
# Start server
go run main.go
# Set webhook
ssh -R 80:localhost:3000 username@ssh.localhost.run
curl -F "url=https://username-XXXX.localhost.run"  https://api.telegram.org/bot<token>/setWebhook 
