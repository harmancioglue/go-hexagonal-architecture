This repository aims to show hexagonal architecture in Go. For simplicity, there is a "Todo Task App" with one endpoint to create task.

![image](https://github.com/harmancioglue/go-hexagonal-architecture/assets/27441734/e066175a-7370-45c6-ae4d-3cedb4b616b4)

## Tech

- Go

## Structure

**/cmd**	

Entrypoint of the project.

**/internal/controller**	

Controller for HTTP requests. It is also primary adapter.

**/internal/infra**	

Intrastructure related code. It is also secondary adapter.

**/internal/core**	

Hexagonal architecture is implemented in directory.

_port_

This directory includes primary and secondary ports. "Service" is primary port and "Repository" is secondary port.
