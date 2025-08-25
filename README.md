# 🖥️ Mini Monitoring App
A sample project designed for learning **SRE** and **DevOps** skills.  
This repository contains a small **Go service** that exposes simple metrics (health check & latency), then gets monitored using **Prometheus** and visualized with **Grafana**.  
The project demonstrates how to containerize applications.



## 🎯 Project Goals
- Build a lightweight service in **Go**  
- **Dockerize** the application for portability  
- Collect metrics with **Prometheus**  
- Visualize metrics on **Grafana dashboards**  
- Prepare infrastructure for CI/CD pipelines (GitHub Actions / GitLab CI)



## 🛠️ Tech Stack
- **Golang** → core service  
- **Docker** → build & run container images  
- **Prometheus** → monitoring & metrics collection  
- **Grafana** → dashboards & visualization



## 🚀 Getting Started

### 1. Clone the repo
```
git clone https://github.com/NeginSal/Mini-Monitoring-App
cd mini-monitoring-app‍‍‍
```
‍‍‍‍
2. Build the Docker image

‍‍‍‍```docker build -t mini-monitoring-app .```

3. Run the container

```docker run -p 8080:8080 mini-monitoring-app```

Now the service is available at:
```http://localhost:8080/health```


