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

## 🌐 Services

Go App → ```http://localhost:8080```

Prometheus →  ```http://localhost:9090```

Grafana →  ```http://localhost:3000```
 (default user/pass: admin / admin)


## 🐳 DockerHub Image
You can pull and run the app directly from DockerHub:
```
docker pull negin007/mini-monitoring-app:latest
docker run -p 8080:8080 negin007/mini-monitoring-app:latest
```
## 📊 Monitoring Setup
- ```Prometheus``` scrapes metrics from the Go application.
- ```Grafana``` provides dashboards and visualization.

## 📝 License

This project is for educational and portfolio purposes. Feel free to fork and adapt.

## 📖 Learn More
You can learn more about this project in this article [mini-monitoring-app](https://dev.to/negin/mini-monitoring-app-in-go-with-prometheus-grafana-cicd-f50)



