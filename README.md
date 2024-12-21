# TODO

- Observability
  - Open Telemetry
    - [x] Add tracer
    - [ ] Add meter
    - [ ] Add logger
  - Prometheus/Grafana
    - [ ] Create an APM dashboard
    - [ ] Configure the exporter 
- [ ] Unit tests
- [ ] Integration tests
- CI
  - [ ] lint
  - [ ] arch-lint
  - [ ] vulnerability check
  - [ ] vet, staticcheck
- CD  
  - build
  - deploy to server
- Faktory
  - [ ] Create a client
  - [ ] Setup a faktory server
  - [ ] Create a worker to send email confirmation

### Security
- Implement rate limiting
- Enforce password requirements (8 characters, numbers, special, etc)
- Permit only 5 login failures, after it, timeout
- captcha in the login
- log the login failures with enought data that can be tracked
- implement cors

### OPS
- Adds the Dockerfile with the build instructions
- Add github action pipeline
- Deploy it to some provider
