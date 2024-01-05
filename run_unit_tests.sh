#!/bin/bash

# Script para executar os testes do projeto Go -> (chmod +x run_unit_tests.sh)

# Compilar e executar os testes
echo "
+===========================+
|  Running Unit Tests... 📌 |
+===========================+
"
go test ./tests/infrastructure/*
echo ""
# Verificar o status de saída dos testes
if [ $? -eq 0 ]; then
  echo "All tests passed! 🚀"
else
  echo "Test failure! ⛔"
fi
echo ""