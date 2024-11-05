# Zmienna dla folderu źródłowego
SRC_DIR = src
# Zmienna dla pliku main.go
MAIN_FILE = $(SRC_DIR)/main.go
# Zmienna dla nazwy binarki, którą chcemy wygenerować
BINARY_NAME = app

# Komenda dla uruchomienia aplikacji
run:
	@go run $(MAIN_FILE)

# Komenda do kompilowania aplikacji (tworzymy binarkę)
build:
	@go build -o $(BINARY_NAME) $(MAIN_FILE)

# Komenda do uruchomienia testów (jeśli są dostępne)
test:
	@go test ./...

# Komenda do czyszczenia folderu (usuwanie binarki)
clean:
	@rm -f $(BINARY_NAME)

# Domyślna komenda, która uruchamia build
all: build
