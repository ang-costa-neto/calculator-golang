# Nome do executável
EXEC = calculator-golang

# Diretório de origem
SRC_DIR = ./cmd/calculator-golang

# Arquivos de origem
SRC = $(SRC_DIR)/main.go

# Alvo padrão
all: build

# Compila o executável
build:
	go build -o $(EXEC) $(SRC)

# Executa o programa
run: build
	./${EXEC}
