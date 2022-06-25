html:
	go run main.go

local-server:
	cd dist && \
  echo "Visit http://localhost:8080 to see the site" && \
  python -m http.server 8080
