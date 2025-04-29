# Description
Nice choice! 🌤️ A Weather Fetcher is a great way to learn how to:
- Work with HTTP APIs
- Parse JSON responses
- Handle environment variables (for your API key)
- Make your Go app more dynamic!

# 🔍 What You'll Build (MVP)
- A CLI or Web API that:
- Accepts a city name
- Calls a public weather API (like OpenWeatherMap)
- Returns temperature, condition, etc.

# Requesite .env
1. Create a `.env` file in the root of your project.
```dotenv
cp .env.example .env
```
2. Add your OpenWeatherMap API key to the `.env` file.
```dotenv
OPENWEATHER_API_KEY=your_api_key_here
```
3. Make sure to replace `your_api_key_here` with your actual API key.
4. Save the file.
5. Make sure to add `.env` to your `.gitignore` file to keep your API key safe.

# Scripts
1. Install dependencies.
```bash
go mod tidy 
```
2. Run the app.
```bash
go run . Jakarta
```
