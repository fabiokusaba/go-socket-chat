# 💬 Go Group Chat — PoC

PoC de um **chat em grupo em tempo real**, desenvolvido em **Go**, utilizando comunicação por **sockets** e uma interface de terminal (TUI) construída com **Bubble Tea**.

O objetivo do projeto é explorar conceitos de comunicação bidirecional, conexões concorrentes e aplicações interativas executadas diretamente no terminal.

---

## 🚀 Tecnologias

* **Go** — linguagem principal
* **TCP Sockets** — comunicação em tempo real entre cliente e servidor
* **Bubble Tea** — TUI (Terminal User Interface)
* **Goroutines** — gerenciamento concorrente das conexões

### Principais conceitos

* TCP sockets
* Comunicação cliente ↔ servidor
* Chat em grupo
* Múltiplas conexões simultâneas
* Concorrência com goroutines
* Broadcast de mensagens
* Comunicação bidirecional
* Terminal User Interface (TUI)

---

## 🏗️ Arquitetura

A aplicação utiliza uma arquitetura simples de **cliente-servidor**.

```text
                    ┌──────────────────┐
                    │      Server      │
                    │                  │
                    │    TCP Socket    │
                    │ Connection Mgmt  │
                    │ Message Broadcast│
                    └────────┬─────────┘
                             │
              ┌──────────────┼──────────────┐
              │              │              │
              ▼              ▼              ▼
        ┌──────────┐   ┌──────────┐   ┌──────────┐
        │ Client 1 │   │ Client 2 │   │ Client 3 │
        │ Bubble   │   │ Bubble   │   │ Bubble   │
        │   Tea    │   │   Tea    │   │   Tea    │
        └──────────┘   └──────────┘   └──────────┘
```

Cada cliente estabelece uma conexão TCP com o servidor.

Quando um cliente envia uma mensagem:

```text
Client
   │
   │ "Olá, pessoal!"
   ▼
Server
   │
   ├──────────────► Client 1
   ├──────────────► Client 2
   └──────────────► Client 3
```

O servidor recebe a mensagem e realiza o **broadcast** para os clientes conectados.

---

## 📁 Estrutura do projeto

Uma possível organização:

```text
go-chat/
│
├── server/
│   │
│   └── main.go
│
├── client/
│   ├── ui/
│   │   ├── message.go
│   │   └── ui.go
│   │
│   └── main.go
└── README.md
```

### `/server/main.go`

Responsável por iniciar o servidor TCP.

### `/client/main.go`

Responsável por iniciar o cliente e a interface TUI.

---

## ⚙️ Como funciona

### 1. Inicialização do servidor

O servidor abre uma porta TCP e aguarda novas conexões:

```text
                 TCP :8080
                    │
                    ▼
              ┌───────────┐
              │  Server   │
              └─────┬─────┘
                    │
              Accept Connection
                    │
                    ▼
                New Client
```

Cada conexão aceita pelo servidor é tratada de forma concorrente utilizando uma goroutine.

---

### 2. Conexão do cliente

O cliente estabelece uma conexão TCP com o servidor:

```text
Client ─────── TCP ───────► Server
```

Depois da conexão, o cliente trabalha com os fluxos de envio e recebimento de mensagens.

```text
             Client
                │
       ┌────────┴────────┐
       │                 │
       ▼                 ▼
   Input/User        Messages
       │                 │
       ▼                 ▼
    Server           Bubble Tea
```

O cliente pode enviar mensagens ao servidor enquanto recebe mensagens enviadas pelos demais usuários.

---

### 3. Broadcast

Quando o servidor recebe uma mensagem:

```text
                 Message
                    │
                    ▼
               ┌─────────┐
               │ Server  │
               └────┬────┘
                    │
        ┌───────────┼───────────┐
        ▼           ▼           ▼
     Client A    Client B    Client C
```

A mensagem é distribuída para os clientes conectados ao chat.

---

## 🖥️ Interface TUI

A interface do cliente utiliza **Bubble Tea**, permitindo criar uma experiência semelhante a um aplicativo de chat diretamente no terminal.

Exemplo conceitual:

```text
┌──────────────────────────────────────────────────────┐
│                  💬 Go Chat                          │
├──────────────────────────────────────────────────────┤
│                                                      │
│  Alice: Olá pessoal!                                │
│                                                      │
│  Bob: Fala Alice!                                    │
│                                                      │
│  Charlie: Tudo certo por aqui 🚀                    │
│                                                      │
│  Bob: Vamos testar o chat?                           │
│                                                      │
├──────────────────────────────────────────────────────┤
│ > Digite sua mensagem...                             │
└──────────────────────────────────────────────────────┘
```

O Bubble Tea segue o padrão **The Elm Architecture**, trabalhando principalmente com:

```text
Model
  │
  ├── View
  │
  └── Update
```

Isso permite separar o estado da aplicação, os eventos recebidos e a renderização da interface.

---

## 📦 Instalação

Clone o projeto:

```bash
git clone https://github.com/fabiokusaba/go-socket-chat.git

cd go-socket-chat
```

Instale as dependências:

```bash
go mod tidy
```

---

## ▶️ Executando

### 1. Inicie o servidor

Em um terminal:

```bash
go run ./server/main.go
```

Por padrão:

```text
Server listening on :8080
```

---

### 2. Inicie o primeiro cliente

Em outro terminal:

```bash
go run ./client/main.go
```

---

### 3. Inicie outros clientes

Abra novos terminais e execute:

```bash
go run ./client/main.go
```

Você poderá abrir múltiplos clientes simultaneamente:

```text
Terminal 1
└── Server

Terminal 2
└── Client - Alice

Terminal 3
└── Client - Bob

Terminal 4
└── Client - Charlie
```

As mensagens enviadas por um cliente serão distribuídas pelo servidor aos demais clientes conectados.

---

## 🧪 Exemplo

Com três clientes conectados:

```text
Alice: Olá!

Bob: Olá Alice!

Charlie: Boa noite pessoal!
```

Todos os clientes recebem as mensagens:

```text
Alice
────────────────────────
Alice: Olá!
Bob: Olá Alice!
Charlie: Boa noite pessoal!


Bob
────────────────────────
Alice: Olá!
Bob: Olá Alice!
Charlie: Boa noite pessoal!


Charlie
────────────────────────
Alice: Olá!
Bob: Olá Alice!
Charlie: Boa noite pessoal!
```

---

## 🧵 Concorrência

Um dos principais objetivos do PoC é explorar o modelo de concorrência do Go.

Cada cliente conectado pode ser tratado de forma concorrente através de goroutines:

```go
go handleConnection(connection)
```

Conceitualmente:

```text
Server
  │
  ├── Goroutine ── Client A
  │
  ├── Goroutine ── Client B
  │
  ├── Goroutine ── Client C
  │
  └── Goroutine ── Client D
```

Dessa forma, o servidor consegue lidar com múltiplas conexões simultaneamente sem precisar processá-las de maneira estritamente sequencial.

---

## 📡 Comunicação

O PoC utiliza comunicação baseada em **TCP sockets**.

Fluxo básico:

```text
┌────────┐                      ┌────────┐
│ Client │                      │ Server │
└───┬────┘                      └───┬────┘
    │                               │
    │──── Connect ─────────────────►│
    │                               │
    │──── Message ─────────────────►│
    │                               │
    │◄──── Broadcast ───────────────│
    │                               │
    │──── Message ─────────────────►│
    │                               │
    │◄──── Broadcast ───────────────│
```

---

## 🎯 Objetivos do PoC

Este projeto foi desenvolvido principalmente para estudar e experimentar:

* [x] TCP sockets
* [x] Comunicação cliente-servidor
* [x] Comunicação em tempo real
* [x] Goroutines
* [x] Gerenciamento de múltiplas conexões
* [x] Broadcast de mensagens
* [x] Bubble Tea
* [x] Desenvolvimento de TUI
* [x] Comunicação bidirecional

---
