#!/bin/bash

# Script para executar os testes do projeto Go -> (chmod +x run_integration_tests.sh)

# Compilar e executar os testes
echo ""
go run main.go &

sleep 2

echo "
+==================================+
|  Running Integration Tests... 📌 |
+==================================+
"
echo ""

go test ./tests/application/*
echo ""

if [ $? -eq 0 ]; then
  echo "All tests passed! 🚀"
else
  echo "Test failure! ⛔"
fi
echo ""

SERVER_PID=$(ps aux | grep 'main.go' | grep -v 'grep' | awk '{print $2}')
if [ -n "$SERVER_PID" ]; then
  echo "Encerrando o servidor Fiber (PID: $SERVER_PID)... 🔥"
  sudo lsof -t -i :3000 | xargs kill -9
fi